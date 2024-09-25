package controller

import (
	"api-gateway/rpcservice/kubernetes-service"
	"api-gateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesController struct{}

// get kubernetes config
func (p *KubernetesController) GetConfig(c *gin.Context) {
	//获取证书，并存入数据库中
	//判断系统中是否有.kube/config文件，如果有就使用这个
	_, err := utils.Tools.PB.KubernetesService.GetKubernetesConfig(context.Background(), &rpc.ConfigRequest{})
	if err != nil {
		utils.Tools.LG.Error("调用k8s服务获取config失败", zap.Error(err))
		return
	}
	path := "D:\\GoCode\\automatic-deployment\\api-gateway\\controller\\config"
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		utils.Tools.LG.Error("kubeConfig转换失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		utils.Tools.LG.Error("集群连接失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		utils.Tools.LG.Error("获取pod列表失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	utils.ResponseSuccess(c, podList)
	return
}

// create kubernetes yaml
func (p *KubernetesController) Create(c *gin.Context) {

}
