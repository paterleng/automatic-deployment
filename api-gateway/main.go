package main

import (
	"api-gateway/api"
	"api-gateway/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	api.CreateApiManager()
	dao.CreateSecretManager()
	api.ApiRoutes(engine)
	engine.Run(":8080")
}
