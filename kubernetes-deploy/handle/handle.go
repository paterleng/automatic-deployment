package handle

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	//创建kubernetes资源
	//定义 Deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "my-app"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "my-app"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "my-container",
						Image: "docker.rainbond.cc/nginx:latest",
					}},
				},
			},
		},
	}
	clientset, err := utils.NewKubeConfig()

	// 创建 Deployment
	deploymentsClient := clientset.AppsV1().Deployments("default")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return nil
}
