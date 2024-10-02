package api

import (
	"context"
	"github.com/micro/go-micro/v2"
	"user-service/rpc"
)

type UserService struct {
}

func Register(service micro.Service) error {
	err := rpc.RegisterUserServiceHandler(service.Server(), &UserService{})
	return err
}
func (h *UserService) CheckStatus(ctx context.Context, req *rpc.UserRequest, resp *rpc.UserResponse) error {
	return nil
}

// 用户的登录
func (h *UserService) UserLogin(ctx context.Context, req *rpc.LoginRequest, resp *rpc.LoginResponse) error {

	resp.Con = "hello"
	return nil
}
