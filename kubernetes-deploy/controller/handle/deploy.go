package handle

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/rpc"
	"kubernetes-deploy/utils"
)

var deployManager DeployManager

type DeployManager interface {
	CommentResource
	ListPods(namespace string) (*corev1.PodList, error)
}

type DeployInterface struct {
	DeployHandle
}

func GetDeployManager() DeployManager {
	//根据资源类型获取到相对应的处理函数
	return deployManager
}

func CreateDeployManager() error {
	var dpmanager DeployInterface
	client, err := utils.NewKubeConfig()
	if err != nil {
		return err
	}
	dpmanager.client = client
	deployManager = &dpmanager
	return nil
}

type DeployHandle struct {
	client *kubernetes.Clientset
	rpc.Deployment
}

func (d *DeployHandle) Before() error {
	//都存到对应的表中，然后在创建deployment的时候在对应的命名空间下创建出来

	return nil
}

func (d *DeployHandle) CreateResources(r interface{}) error {
	req := r.(rpc.Deployment)
	//在创建资源之前先创建相对应的secret
	d.Deployment = req
	err := d.Before()

	err = CheckNameSpace(d.client, req.NameSpace)
	if err != nil {
		return err
	}

	if req.Labels == nil {
		req.Labels = make(map[string]string)
	}
	if req.MatchLabels == nil {
		req.MatchLabels = make(map[string]string)
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &req.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: req.MatchLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: req.Labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  req.Name,
						Image: "docker.rainbond.cc/" + req.ImageName,
					}},
					RestartPolicy: corev1.RestartPolicyOnFailure,
				},
			},
		},
	}

	deploymentsClient := d.client.AppsV1().Deployments(req.NameSpace)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	//在创建之后，会有一个事件，去通知到对应日志处理，用于进行日志输出

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta())
	return nil
}

func (d *DeployHandle) After() error {
	return nil
}

func (d *DeployHandle) ListPods(namespace string) (*corev1.PodList, error) {
	podList, err := d.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	return podList, err
}
