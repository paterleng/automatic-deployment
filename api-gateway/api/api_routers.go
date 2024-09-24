package api

import "github.com/gin-gonic/gin"

func ApiRoutes(r *gin.Engine) {
	UserApi(r)
	KubernetesApi(r)
	CodePackageApi(r)
}
