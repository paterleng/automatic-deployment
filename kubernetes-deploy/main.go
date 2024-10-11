package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go.uber.org/zap"
	"kubernetes-deploy/api"
	"kubernetes-deploy/utils"
)

func main() {
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(
		micro.Name("kubernetes-deploy"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()
	if err := api.Register(service); err != nil {
		utils.Tools.LG.Error("服务注册失败：", zap.Error(err))
		return
	}
	utils.Tools.LG.Info("kubernetes-deploy服务注册成功")
	if err := service.Run(); err != nil {
		utils.Tools.LG.Error("服务运行失败：", zap.Error(err))
		return
	}
	utils.Tools.LG.Info("kubernetes-deploy服务启动成功")
}
