package middleware

import "github.com/gin-gonic/gin"

// 基于jwt的中间件验证
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//解析客户端发送的请求 请求头的Header的authorization中 使用bearer开头
		authHeader := c.Request.Header.Get("")
	}
}
