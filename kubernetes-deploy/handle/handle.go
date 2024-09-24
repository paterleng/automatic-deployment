package handle

import (
	"context"
	"github.com/micro/go-micro/v2"
	"kubernetes-deploy/rpc"
)

type KubernetesDeploy struct{}

func Register(service micro.Service) error {
	err := rpc.RegisterKubernetesDeployHandler(service.Server(), &KubernetesDeploy{})
	return err
}

func (h *KubernetesDeploy) CheckStatus(ctx context.Context, req *rpc.Request, rsp *rpc.Response) error {
	return nil
}
