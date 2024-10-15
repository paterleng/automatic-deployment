package api

import (
	"context"
	"github.com/micro/go-micro/v2"
	"user-service/rpc"
	"user-service/utils"
)

type UserService struct {
}

// 检查用户是否注册过
func (user *UserService) UserCheckMail(ctx context.Context, request *rpc.UserCheckMailRequest, response *rpc.UserCheckMailResponse) error {
	var sum int64
	err := utils.Tools.DB.Table("user").Select("username = ?", request.Mailbox).Count(&sum).Error
	if err != nil {
		return err
	}
	if int(sum) == 1 {
		response.Msg = true
	} else {
		response.Msg = true
	}
	return nil
}

// 用户的登录
func (h *UserService) UserLogin(ctx context.Context, req *rpc.LoginRequest, resp *rpc.LoginResponse) error {
	// 接收传来的req
	if req.MailPasswd == " " {

	}
	return nil
}

func Register(service micro.Service) error {
	err := rpc.RegisterUserServiceHandler(service.Server(), &UserService{})
	return err
}
func (h *UserService) CheckStatus(ctx context.Context, req *rpc.UserRequest, resp *rpc.UserResponse) error {
	return nil
}
