package main

import (
	"api-gateway/api"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	api.CreateApiManager()
	api.ApiRoutes(engine)
	engine.Run(":8080")

}
