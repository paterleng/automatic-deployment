package controller

import (
	"api-gateway/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NodeController struct {
	PB *utils.Pb
	LG *zap.Logger
}

func (m *NodeController) GetNodeInfo(c *gin.Context) {

}
