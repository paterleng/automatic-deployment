package api

import "github.com/gin-gonic/gin"

func KubernetesApi(r *gin.Engine) {
	ks := r.Group("/kubernetes")
	ks.GET("/config", GetManager().GetConfig)
	ks.POST("/create")
}
