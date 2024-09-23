package handle

import (
	"code-package/rpc"
	"context"
	"github.com/micro/go-micro/v2"
)

type CodePackage struct{}

func Register(service micro.Service) error {
	err := rpc.RegisterCodePackageHandler(service.Server(), &CodePackage{})
	return err
}

func (h *CodePackage) CheckStatus(ctx context.Context, req *rpc.Request, rsp *rpc.Response) error {
	return nil
}
