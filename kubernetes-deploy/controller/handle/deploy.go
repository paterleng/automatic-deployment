package handle

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/model"
	"kubernetes-deploy/rpc"
	"kubernetes-deploy/utils"
	"sync"
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
	dpmanager.DB = utils.Tools.DB
	deployManager = &dpmanager

	return nil
}

type DeployHandle struct {
	client     *kubernetes.Clientset
	DB         *gorm.DB
	Deployment rpc.Deployment
}

func (d *DeployHandle) Before() error {
	//都存到对应的表中，然后在创建deployment的时候在对应的命名空间下创建出来
	//从表中查出相应的secret，然后创建出来
	var secrets []model.Secret
	secretMap := make(map[string]v1.Secret)
	var wg sync.WaitGroup
	err := d.DB.Model(&model.Secret{}).Find(&secrets).Error
	if err != nil {
		return err
	}
	getSecret, err := GetSecret(d.client, d.Deployment.NameSpace)
	if err != nil {
		return err
	}
	for _, secret := range getSecret.Items {
		secretMap[secret.Name] = secret
	}
	//循环遍历，找到该命名空间下没有的secret，创建出来
	for _, item := range secrets {
		if _, ok := secretMap[item.Name]; ok {
			continue
		}
		//创建
		wg.Add(1)
		go func(model.Secret) error {
			err := CreateSecret(d.client, d.Deployment.NameSpace, item)
			if err != nil {
				return err
			}
			wg.Done()
			return nil
		}(item)
	}
	wg.Wait()
	return nil
}

func (d *DeployHandle) CreateResources(r interface{}) error {
	req := r.(*rpc.Deployment)
	d.Deployment = *req
	err := CheckNameSpace(d.client, req.NameSpace)
	if err != nil {
		return err
	}
	err = d.Before()
	if err != nil {
		return err
	}

	if req.Labels == nil {
		req.Labels = make(map[string]string)
		req.Labels["deployment"] = req.Name
	}
	if req.MatchLabels == nil {
		req.MatchLabels = make(map[string]string)
		req.MatchLabels["deployment"] = req.Name
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
					ImagePullSecrets: imagePullSecret(d.client, d.Deployment.NameSpace),
					Containers: []corev1.Container{{
						Name: req.Name,
						//"docker.rainbond.cc/" +
						Image: req.ImageName,
					}},
					RestartPolicy: corev1.RestartPolicyAlways,
				},
			},
		},
	}

	deploymentsClient := d.client.AppsV1().Deployments(req.NameSpace)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	//在创建完deployment之后，创建service，用于对外进行访问使用
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
