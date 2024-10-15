package utils

import (
	coderpc "api-gateway/rpcservice/code-service"
	kubernetesrpc "api-gateway/rpcservice/kubernetes-service"
	userrpc "api-gateway/rpcservice/user-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type Pb struct {
	UserService       userrpc.UserService
	CodeService       coderpc.CodePackageService
	KubernetesService kubernetesrpc.KubernetesDeployService
}

func DiscoveryService() {
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(micro.Registry(reg))
	Tools.PB.UserService = userrpc.NewUserService("user-service", service.Client())
	Tools.PB.KubernetesService = kubernetesrpc.NewKubernetesDeployService("kubernetes-deploy", service.Client())
	//Tools.PB.CodeService = coderpc.NewCodePackageService("code-package", service.Client())
	Tools.LG.Info("客户端创建成功")
}
