package api

import "github.com/gin-gonic/gin"

func KubernetesApi(r *gin.Engine) {
	ks := r.Group("/user")
	ks.POST("/create", GetManager().Create)

}
