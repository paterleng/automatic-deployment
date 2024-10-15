package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"user-service/pkg"
	"user-service/utils"
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
		c.Next()
	}
}
