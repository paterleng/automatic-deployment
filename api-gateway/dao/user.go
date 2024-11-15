package dao

import (
	"api-gateway/model"
	"api-gateway/pkg"
	"api-gateway/utils"
)

// 检查用户邮箱是否注册 注册返回true
func UserCheckMail(mailbox string) bool {
	var sum int64
	err := utils.Tools.DB.Table("users").Where("mailbox = ?", mailbox).Count(&sum).Error
	if err != nil {
		return false
	}
	if int(sum) == 1 {
		return true
	} else {
		return false
	}
}

// 用户注册
func UserRegister(user model.User) (err error) {
	err = utils.Tools.DB.Table("users").Create(user).Error
	return err
}

// 用户登录密码验证  并生成atoken rtoken
func UserCheckPW(mailbox, password string) (user *model.User, err error) {
	u := new(model.User)
	err = utils.Tools.DB.Table("users").Where("mailbox = ?", mailbox).Find(&u).Error
	user = u
	// 密码验证
	ok := user.CheckPassword(password)
	if !ok {
		utils.Tools.LG.Error("密码验证失败")
		return nil, err
	}
	atoken, rtoken, err := pkg.GenToken(user.UserID, password)
	if err != nil {
		utils.Tools.LG.Error("生成token失败")
		return nil, err
	}
	user.Atoken = atoken
	user.Rtoken = rtoken
	return
}

// 用户验证码登录验证和返回数据
func SelectUser(mailbox string) (user *model.User, err error) {
	u := new(model.User)
	err = utils.Tools.DB.Table("users").Where("mailbox = ?", mailbox).Find(&u).Error
	user = u
	atoken, rtoken, err := pkg.GenToken(user.UserID, user.MailPassWD)
	if err != nil {
		utils.Tools.LG.Error("生成token失败")
		return nil, err
	}
	user.Atoken = atoken
	user.Rtoken = rtoken
	return
}

// 用户id 查用户信息 加载关联的角色表信息
func SelectIdUser(id string) (user model.User, err error) {
	err = utils.Tools.DB.Preload("Role").Where("userid = ?", id).First(&user).Error
	return
}

// 更新用户登录设备
func UpdateDevice(user model.User, device string) (err error) {
	err = utils.Tools.DB.Table("users").Where("mailbox = ?", user.MailBox).Update("lastlogindevice", device).Error
	return
}

// 创建角色
func CreateRoleEmpt(role model.Role) (err error) {
	err = utils.Tools.DB.Table("roles").Create(&role).Error
	return err
}

// 用户赋权
func EmpowerUser(user model.User) (err error) {
	err = utils.Tools.DB.Table("users").Where("mailbox = ?", user.MailBox).Update("role_name", user.RoleName).Error
	return
}
