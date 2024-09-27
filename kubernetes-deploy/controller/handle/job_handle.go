package handle

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-deploy/utils"
)

type JobHandle struct{}

var jobManager JobManager

type JobInterface struct {
	JobHandle
}

type JobManager interface {
	CommentResource
}

func GetJobManager() JobManager {
	return jobManager
}

func CreateJobManager() {
	var manager JobInterface
	jobManager = &manager
}

func (d *JobHandle) Before() error {
	return nil
}

func (d *JobHandle) CreateResources() error {
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

func (d *JobHandle) After() error {
	return nil
}
