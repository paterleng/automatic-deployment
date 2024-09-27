package handle

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-deploy/utils"
)

var deployManager DeployManager

type DeployManager interface {
	CommentResource
	ListPods() ([]corev1.Pod, error)
}

type DeployInterface struct {
	DeployHandle
}

func GetDeployManager() DeployManager {
	//根据资源类型获取到相对应的处理函数
	return deployManager
}

func CreateDeployManager() {
	var dpmanager DeployInterface
	deployManager = &dpmanager
}

type DeployHandle struct{}

func (d *DeployHandle) Before() error {

	return nil
}

func (d *DeployHandle) CreateResources() error {
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

func (d *DeployHandle) After() error {

	return nil
}

func (d *DeployHandle) ListPods() ([]corev1.Pod, error) {
	return nil, nil
}