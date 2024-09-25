package api

import "github.com/gin-gonic/gin"

type UserInterface interface {
	Login(c *gin.Context)
}

type KubernetesApiInterface interface {
	GetConfig(c *gin.Context)
	Create(c *gin.Context)
}
