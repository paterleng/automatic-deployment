package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go.uber.org/zap"
	"kubernetes-deploy/handle"
	"kubernetes-deploy/utils"
)

func main() {
	//注册rpc服务
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	//向etcd注册一个新的服务
	service := micro.NewService(
		micro.Name("kubernetes-deploy"),
		micro.Version("latest"),
		// 使用服务注册插件
		micro.Registry(reg),
	)
	//初始化服务
	service.Init()
	if err := handle.Register(service); err != nil {
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
