package middleware

import (
	"api-gateway/dao"
	"api-gateway/model"
	"api-gateway/pkg"
	"api-gateway/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

// 基于jwt的中间件验证
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//解析客户端发送的请求 请求头的Header的authorization中 使用bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.Tools.LG.Error("请求中auth为空", zap.Error(errors.New("token不存在")))
			c.Abort()
			return
		}
		// 接受到的数据按照空格分割出token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.Tools.LG.Error("请求中auth格式有误", zap.Error(errors.New("token不正确")))
			c.Abort()
			return
		}
		// 在这里进行token解析 获取用户信息
		mc, err := pkg.ParseToken(parts[1])
		if err != nil {
			utils.Tools.LG.Error("token 无效", zap.Error(errors.New("token不正确")))
			c.Abort()
			return
		}

		//id是唯一标识 这里获取id就可以
		c.Set("userid", mc.UserID)
		// 使用userid查出用户信息
		var user model.User
		user, err = dao.SelectIdUser(mc.UserID)
		if err != nil {
			utils.Tools.LG.Error("中间件使用id查找数据库出错")
			utils.ResponseError(c, utils.CodeServerBusy)
			c.Abort()
			return
		}

		// 检查该用户是否可以访问此接口
		b := false
		//fmt.Println(c.Request.URL.Path + "test path")
		//fmt.Println(user.Role.ApiRouter)
		for _, route := range user.Role.ApiRouter {
			//if strings.HasPrefix(c.Request.URL.Path, route) {
			if c.Request.URL.Path == route {
				b = true
				break
			}
		}
		if !b {
			utils.Tools.LG.Error("用户无权限")
			utils.ResponseError(c, utils.CodeNoPer)
			c.Abort()
			return
		}
		// 检查使用token请求信息的设备是否改变（登录的设备）
		deviceIdentifier := c.ClientIP() + c.Request.UserAgent()
		if user.LastLoginDevice != deviceIdentifier {
			utils.Tools.LG.Error("登录设备不同出错")
			c.Abort()
			return
		}
		c.Next()
	}
}
