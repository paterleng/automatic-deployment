package handle

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/model"
	"kubernetes-deploy/rpc"
	"kubernetes-deploy/utils"
)

var serviceManager ServiceManager

type ServiceManager interface {
	CommentResource
}

type ServiceInterface struct {
	DeployHandle
}

func GetServiceManager() ServiceManager {
	//根据资源类型获取到相对应的处理函数
	return serviceManager
}

func CreateServiceManager() error {
	var dpmanager DeployInterface
	client, err := utils.NewKubeConfig()
	if err != nil {
		return err
	}
	dpmanager.client = client
	deployManager = &dpmanager
	return nil
}

type ServiceHandle struct {
	client  *kubernetes.Clientset
	service model.ServiceReq
}

func (s *ServiceHandle) Before() error {
	//检查端口是否可用
	//listService, err := s.client.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	utils.Tools.LG.Error("获取service列表失败", zap.Error(err))
	//	return err
	//}

	return nil
}

func (s *ServiceHandle) CreateResources(r interface{}) error {
	req := r.(rpc.Service)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: corev1.ServiceSpec{
			Selector: req.Selector,
			Ports: []corev1.ServicePort{
				{
					Port:       req.Port,                         //通过service服务暴露的端口
					TargetPort: intstr.FromInt32(req.TargetPort), // Pod 中服务监听的端口
					Protocol:   corev1.ProtocolTCP,               //端口传输协议
					NodePort:   req.NodePort,                     //外部要访问的端口
				},
			},
			Type: corev1.ServiceTypeNodePort, // 可以根据需要选择其他类型
		},
	}
	_, err := s.client.CoreV1().Services(req.NameSpace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create service: %v", err)
	}
	return nil
}

func (s *ServiceHandle) After() error {

	return nil
}
