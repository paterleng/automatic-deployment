package utils

import (
	coderpc "api-gateway/rpc/code-service"
	kubernetesrpc "api-gateway/rpc/kubernetes-service"
	userrpc "api-gateway/rpc/user-service"
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
	//服务发现
	var pb *Pb
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(micro.Registry(reg))
	pb.UserService = userrpc.NewUserService("code-package", service.Client())
	pb.KubernetesService = kubernetesrpc.NewKubernetesDeployService("kubernetes-deploy", service.Client())
	pb.CodeService = coderpc.NewCodePackageService("code-package", service.Client())
	Tools.LG.Info("客户端创建成功")
	Tools.PB = pb
}
