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
	err := handle.GetDeployManager().CreateResources()
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("创建%s资源失败", req.ResourceType), zap.Error(err))
	}

	return nil
}
