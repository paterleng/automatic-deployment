package api

import "github.com/gin-gonic/gin"

func UserApi(r *gin.Engine) {
	user := r.Group("/user")
	user.GET("")
}
