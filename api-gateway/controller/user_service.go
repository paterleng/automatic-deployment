package controller

import (
	"api-gateway/dao"
	"api-gateway/pkg"
	"net/http"

	//"api-gateway/dao"
	"api-gateway/model"

	"api-gateway/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
	"time"
)

// controller和logic层写一块（方法不多)

type UserServiceController struct {
	PB *utils.Pb
	LG *zap.Logger
}

const (
	lenMailCaptcha = 6
	expireTime     = 60 * time.Second
)

var CodeStore sync.Map

// 邮箱账号发送验证码接口
func (u *UserServiceController) SendMail(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		utils.Tools.LG.Error("参数获取失败")
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	if pkg.IsValidEmail(user.MailBox) {
		// 判断用户是否存在
		//b := dao.UserCheckMail(user.MailBox)
		//if b {
		// 用户存在 发送验证码

		// 修改 是邮箱就可以发送邮件
		// 生成验证码
		captcha, err := pkg.GenerateCaptcha(lenMailCaptcha)
		if err != nil {
			utils.ResponseError(c, utils.COdeCaptcha)
			return
		}
		captchaExpire, err := model.GenerateCaptchaExpire(captcha, expireTime)
		if err != nil {
			utils.ResponseError(c, utils.CodeCaptchaExpire)
			return
		}
		// 在内存中存储验证码
		CodeStore.Store(user.MailBox, captchaExpire)
		mail := &pkg.Email{
			To:       user.MailBox,
			Subject:  "登录验证码",
			BodyType: "text/plain",
			Body:     captcha,
		}
		err = mail.SendEmail()
		if err != nil {
			utils.ResponseError(c, utils.CodeServerBusy)
			return
		}
	}
	utils.ResponseSuccess(c, "发送成功")
}

// 用户注册
func (u *UserServiceController) RegisterMail(c *gin.Context) {
	// 验证码发送之后 填写用户数据 对用户数据进行处理并存入数据库
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 检查用户验证码是否正确
	if !VerifyCode(user.MailBox, user.Captcha) {
		utils.ResponseError(c, utils.COdeCaptcha)
		utils.Tools.LG.Error("验证码错误")
		return
	}
	// 检查用户是否注册
	if dao.UserCheckMail(user.MailBox) {
		utils.ResponseError(c, utils.CodeUserExist)
		utils.Tools.LG.Error("用户已存在")
		return
	}
	// 用户uuid生成  密码加密
	user.UserID, _ = pkg.GenerateUUID()
	aToken, rToken, err := pkg.GenToken(user.UserID, user.MailPassWD)
	if err != nil {
		utils.ResponseError(c, utils.CodeCreateError)
		utils.Tools.LG.Error("用户注册生成token失败")
		return
	}
	user.Atoken = aToken
	user.Rtoken = rToken
	err = user.SetPassword(user.MailPassWD)
	if err != nil {
		utils.ResponseError(c, utils.CodeCreateError)
		utils.Tools.LG.Error("用户注册生成密码失败")
		return
	}
	// 注册进数据库
	err = dao.UserRegister(user)
	if err != nil {
		utils.ResponseError(c, utils.CodeUserRes)
		utils.Tools.LG.Error("用户注册数据库失败")
		return
	}
	utils.ResponseSuccess(c, user)
}

// 登录接口
func (u *UserServiceController) Login(c *gin.Context) {
	//接收参数
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.Tools.LG.Error("数据绑定失败")
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	// 判断用户是否存在
	b := dao.UserCheckMail(user.MailBox)
	if !b {
		utils.ResponseError(c, utils.CodeUserNotExist)
		return
	}
	// 判断用户登录方式
	// 密码登录
	if len(user.MailPassWD) >= 6 {
		u, err := dao.UserCheckPW(user.MailBox, user.MailPassWD)
		if err != nil {
			utils.ResponseError(c, utils.CodeInvalidPassword)
			return
		}
		user = *u
	} else {
		// 验证码登录
		if VerifyCode(user.MailBox, user.Captcha) {
			// 验证码正确 返回用户信息
			u, err := dao.SelectUser(user.MailBox)
			if err != nil {
				utils.ResponseError(c, utils.CodeServerBusy)
				return
			}
			user = *u
		} else {
			utils.ResponseError(c, utils.COdeCaptcha)
			utils.Tools.LG.Error("验证码错误")
			return
		}
	}
	// 获取登录设备信息 并更新用户表的登录信息
	deviceIdentifier := c.ClientIP() + c.Request.UserAgent()
	err = dao.UpdateDevice(user, deviceIdentifier)
	if err != nil {
		utils.ResponseError(c, utils.CodeResetLogin)
		utils.Tools.LG.Error("登录设备信息更新失败")
		return
	}
	utils.ResponseSuccess(c, user)
}

// 超级管理员赋权接口
func (u *UserServiceController) Empowerment(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.Tools.LG.Error("绑定用户结构体失败")
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	// 对用户进行角色赋权改变
	err = dao.EmpowerUser(user)
	if err != nil {
		utils.ResponseError(c, utils.CodeEmpowerUser)
		utils.Tools.LG.Error("用户角色赋权失败")
		return
	}
	utils.ResponseSuccess(c, "赋权成功")
}

// 超级管理员创建新的角色接口
func (u *UserServiceController) CreatEmpt(c *gin.Context) {
	var role model.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		utils.Tools.LG.Error("创建角色结构体数据有误")
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	// 存入数据库
	err = dao.CreateRoleEmpt(role)
	if err != nil {
		utils.Tools.LG.Error("角色创建失败")
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, role)
}

// 验证验证码是否正确 实现验证map里是否存在并且不过期
func VerifyCode(mailbox, inputCode string) bool {
	if value, ok := CodeStore.Load(mailbox); ok {
		captchCode := value.(model.CaptchaExpire)
		if time.Now().After(captchCode.ExpiresAt) {
			return false
			// 过期之后执行删除功能
			CodeStore.Delete(mailbox)
		}
		if captchCode.Code == inputCode {
			return true
		}
	}
	return false
}

func (u *UserServiceController) GetRouters(c *gin.Context) {
	var route model.Route
	child := make([]model.Children, 3)
	child[0].Name = "dashboard"
	child[0].Path = "/dashboard"
	child[0].Component = "home"
	child[0].Meta.Title = "控制台"
	child[0].Meta.Icon = "el-icon-menu"
	child[0].Meta.Affix = true

	child[1].Name = "newPage"
	child[1].Path = "/newPage"
	child[1].Component = "newPage"
	child[1].Meta.Title = "新的页面"

	child[2].Name = "userCenter"
	child[2].Path = "/userCenter"
	child[2].Component = "userCenter"
	child[2].Meta.Title = "账号信息"
	child[2].Meta.Icon = "el-icon-user"
	child[2].Meta.Tag = "NEW"
	menus := make([]model.Menu, 1)
	menus[0].Name = "home"
	menus[0].Path = "/home"
	menus[0].Meta.Title = "首页"
	menus[0].Meta.Icon = "el-icon-eleme-filled"
	menus[0].Meta.Type = "menu"
	menus[0].Children = child
	route.Menu = menus
	permissions := []string{
		"list.add",
		"list.edit",
		"list.delete",
		"user.add",
		"user.edit",
		"user.delete",
	}
	dashboardGrid := []string{
		"welcome",
		"time",
		"progress",
		"echarts",
	}
	route.DashboardGrid = dashboardGrid
	route.Permissions = permissions

	data := struct {
		Code    int         `json:"code"`
		Data    model.Route `json:"data"`
		Message string      `json:"message"`
	}{
		Code:    200,
		Data:    route,
		Message: "",
	}
	c.JSON(http.StatusOK, data)
}
