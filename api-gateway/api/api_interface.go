package api

import "github.com/gin-gonic/gin"

type UserInterface interface {
	Login(c *gin.Context)
}

type KubernetesApiInterface interface {
	GetConfig(c *gin.Context)
	CreateResource(c *gin.Context)
	GetSecret(c *gin.Context)
	CreateSecret(c *gin.Context)
	UpdateSecret(c *gin.Context)
	DeleteSecret(c *gin.Context)
}
