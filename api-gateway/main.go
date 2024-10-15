package main

import (
	"api-gateway/api"
	"api-gateway/controller"
	"api-gateway/model"
	"api-gateway/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	engine := gin.Default()
	api.CreateApiManager()
	api.ApiRoutes(engine)
	//utils.Tools.DB.SingularTable(true)
	err := utils.Tools.DB.AutoMigrate(&model.Role{})
	err = utils.Tools.DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	// 生成bcrypt密码加密  然后可以进行验证
	a := model.SetPasswords("1234567")
	fmt.Println(a)
	engine.Run(":8080")
	go cleanUpExpiredCodes()
}

// 定期清理过期验证码  防止内存消耗过大
func cleanUpExpiredCodes() {
	ticker := time.NewTicker(10 * time.Minute) // 每10分钟清理一次
	for range ticker.C {
		controller.CodeStore.Range(func(key, value interface{}) bool {
			verificationCode := value.(model.CaptchaExpire)
			if time.Now().After(verificationCode.ExpiresAt) {
				controller.CodeStore.Delete(key)
				//fmt.Println("清理过期验证码:", key)
			}
			return true
		})
	}
}
