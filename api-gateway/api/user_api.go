package api

import (
	"api-gateway/utils"
	"github.com/gin-gonic/gin"
)

func UserApi(r *gin.Engine) {
	r.Use(utils.Cors())
	user := r.Group("/api/user")
	// 实现验证码发送的接口
	user.POST("/loginmail", GetManager().SendMail)
	// 实现用户注册接口
	user.POST("/register", GetManager().RegisterMail)
	// 登录接口
	user.POST("/login", GetManager().Login)
	user.GET("/router", GetManager().GetRouters)
	// 登录之后的操作  请求都需要使用atoken
	userLogin := r.Group("/user/login")
	//userLogin.Use(middleware.JWTAuthMiddleware())
	{
		// 赋权和创建角色接口 （超级管理员）
		userLogin.POST("/empt", GetManager().Empowerment)
		userLogin.POST("/creatempt", GetManager().CreatEmpt)
	}
}
