package api

import "github.com/gin-gonic/gin"

type UserInterface interface {
	Login()
}

type KubernetesApiInterface interface {
	Create(c *gin.Context)
}
