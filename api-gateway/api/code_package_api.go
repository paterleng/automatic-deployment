package api

import "github.com/gin-gonic/gin"

func CodePackageApi(r *gin.Engine) {
	cp := r.Group("/api/code")

	cp.GET("/api/test")
}
