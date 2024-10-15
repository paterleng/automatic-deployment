package api

import "github.com/gin-gonic/gin"

type UserInterface interface {
	Login(c *gin.Context)
	LoginMail(c *gin.Context)
	RegisterMail(c *gin.Context)
	Test(c *gin.Context)
	Empowerment(c *gin.Context)
	CreatEmpt(c *gin.Context)
}

type KubernetesApiInterface interface {
	GetConfig(c *gin.Context)
	CreateResource(c *gin.Context)
	GetSecret(c *gin.Context)
	CreateSecret(c *gin.Context)
	UpdateSecret(c *gin.Context)
	DeleteSecret(c *gin.Context)
	ClusterDocking(c *gin.Context)
	GetNodeInfo(c *gin.Context)
}
