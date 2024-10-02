package api

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"go.uber.org/zap"
	"kubernetes-deploy/controller/handle"
	"kubernetes-deploy/rpc"
	"kubernetes-deploy/utils"
)

type KubernetesDeploy struct{}

func Register(service micro.Service) error {
	err := rpc.RegisterKubernetesDeployHandler(service.Server(), &KubernetesDeploy{})
	return err
}

func (h *KubernetesDeploy) GetKubernetesConfig(ctx context.Context, req *rpc.ConfigRequest, resp *rpc.ConfigResponse) (err error) {
	resp.Config = "123"
	return nil
}

func (h *KubernetesDeploy) CheckStatus(ctx context.Context, req *rpc.KsRequest, resp *rpc.KsResponse) error {
	return nil
}

func (h *KubernetesDeploy) CreateResource(ctx context.Context, req *rpc.CreateResourceRequest, resp *rpc.CreateResourceResponse) error {
	//校验资源类型，从而对不同资源类型做不同的处理
	var err error
	switch req.ResourceType {
	case utils.DeploymentResource:
		err = handle.GetDeployManager().CreateResources(req.DeploymentResource)
	case utils.JobResource:
		err = handle.GetJobManager().CreateResources(req.JobResource)
	case utils.CornJobResource:
		handle.GetCornJobManager().CreateResources(req.CornJobResource)
	case utils.ServiceResource:
		handle.GetServiceManager().CreateResources(req.ServiceResource)
	case utils.PodResource:

	default:
		utils.Tools.LG.Error("不能创建此资源", zap.Error(err))
		return err
	}
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("创建%s资源失败", req.ResourceType), zap.Error(err))
		return err
	}
	return nil
}
