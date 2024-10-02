package controller

import (
	"api-gateway/model"
	"api-gateway/rpcservice/kubernetes-service"
	"api-gateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesController struct {
	PB *utils.Pb
	LG *zap.Logger
}

// get kubernetes config
func (p *KubernetesController) GetConfig(c *gin.Context) {
	//获取证书，并存入数据库中
	//判断系统中是否有.kube/config文件，如果有就使用这个
	_, err := p.PB.KubernetesService.GetKubernetesConfig(context.Background(), &rpc.ConfigRequest{})
	if err != nil {
		p.LG.Error("调用k8s服务获取config失败", zap.Error(err))
		return
	}
	path := "D:\\GoCode\\automatic-deployment\\api-gateway\\controller\\config"
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		p.LG.Error("kubeConfig转换失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		p.LG.Error("集群连接失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		p.LG.Error("获取pod列表失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	c.Set("kubeconfig", kubeConfig)
	utils.ResponseSuccess(c, podList)
	return
}

// create kubernetes resource
func (p *KubernetesController) CreateResource(c *gin.Context) {
	//	参数处理
	var m model.ResourceReq
	if err := c.ShouldBindJSON(&m); err != nil {
		p.LG.Error("参数绑定失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeInvalidParam, err)
		return
	}

	req := &rpc.CreateResourceRequest{
		ResourceType: m.ResourceType,
		//Name:         m.Name,
		//NameSpace:    m.NameSpace,
		//ImageName:    m.ImageName,
		//Replicas:     m.Replicas,
		//Labels:       m.Labels,
		//MatchLabels:  m.MatchLabels,
	}

	_, err := p.PB.KubernetesService.CreateResource(c, req)
	if err != nil {
		p.LG.Error("创建资源失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}
