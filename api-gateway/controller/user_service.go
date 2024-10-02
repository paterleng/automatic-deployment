package controller

import (
	"api-gateway/rpcservice/user-service"
	"api-gateway/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserServiceController struct {
	PB *utils.Pb
	LG *zap.Logger
}

func (u *UserServiceController) Login(c *gin.Context) {
	//转到另一个服务器进行处理
	resp, err := u.PB.UserService.UserLogin(c, &rpc.LoginRequest{})
	if err != nil {
		utils.ResponseError(c, 1004)
		return
	}
	utils.ResponseSuccess(c, resp)
}
