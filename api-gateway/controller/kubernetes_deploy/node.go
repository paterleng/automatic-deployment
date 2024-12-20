package controller

import (
	"api-gateway/dao"
	internal "api-gateway/internal/kubenetes"
	"api-gateway/model"
	"api-gateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
)

type NodeController struct {
	PB *utils.Pb
	LG *zap.Logger
}

// 对接集群
func (m *NodeController) ClusterDocking(c *gin.Context) {
	var req model.ClusterResquest
	var cluster model.Cluster
	if err := c.ShouldBindJSON(&req); err != nil {
		m.LG.Error("参数错误")
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	//	创建config并入库
	//	首先先判断在.kube下有没有config，如果有就直接使用这个，如果没有就自己创建一个
	if utils.FileExists(utils.KubeConfigFile) {
		//文件存在，获取到文件中的内容并赋值
		data, err := utils.GetFileData(utils.KubeConfigFile)
		if err != nil {
			m.LG.Error("读取文件失败")
			utils.ResponseError(c, utils.CodeServerBusy)
			return
		}
		cluster.Config = data
	} else {
		//创建config
		config, err := internal.GetKubeConfig().CreateConfig(req)
		if err != nil {
			m.LG.Error("创建config字符串失败")
			utils.ResponseError(c, utils.CodeServerBusy)
			return
		}
		cluster.Config = config
	}
	cluster.ClusterAdr = req.ClusterAdr
	cluster.Name = req.Name
	//在对接前先看一下是否已经对接过

	clientSet, err := utils.CreateClientSet(cluster.Config)
	if err != nil {
		m.LG.Error("创建客户端失败")
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	nodeList, err := clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		m.LG.Error("获取节点信息失败")
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	for _, item := range nodeList.Items {
		cluster.Status = "Activing"
		cluster.Arm = item.Status.NodeInfo.Architecture
		cluster.OsImage = item.Status.NodeInfo.OSImage
		cluster.Version = item.Status.NodeInfo.KubeletVersion
	}

	versionInfo, err := discovery.NewDiscoveryClient(clientSet.RESTClient()).ServerVersion()
	if err != nil {
		m.LG.Error("获取集群版本信息失败")
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}

	cluster.Version = versionInfo.String()
	err = dao.GetNodeManager().Create(cluster)
	if err != nil {
		m.LG.Error("对接集群失败")
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}

func (m *NodeController) GetNodeInfo(c *gin.Context) {
	clusters, err := dao.GetNodeManager().Get()
	if err != nil {
		m.LG.Error("获取集群节点失败")
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, clusters)
}
