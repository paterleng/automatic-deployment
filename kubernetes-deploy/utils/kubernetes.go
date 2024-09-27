package utils

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubeConfig() (clientset *kubernetes.Clientset, err error) {
	path := "D:\\GoCode\\automatic-deployment\\api-gateway\\controller\\config"
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		Tools.LG.Error("kubeConfig转换失败", zap.Error(err))
		return
	}

	clientset, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		Tools.LG.Error("集群连接失败", zap.Error(err))
		return
	}
	return
}
