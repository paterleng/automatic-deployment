package controller

import (
	"api-gateway/model"
	rpc "api-gateway/rpcservice/code-service"
	"api-gateway/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CodePackageController struct {
	PB *utils.Pb
	LG *zap.Logger
}

// 拉取代码
func (p *CodePackageController) PullCode(c *gin.Context) {
	//处理数据
	var code model.CodeRepository
	if err := c.ShouldBindJSON(&code); err != nil {
		p.LG.Error("参数绑定失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeInvalidParam, err)
		return
	}
	codeRequest := rpc.PullCodeRequest{
		Url:      code.URL,
		Branch:   code.Branch,
		Account:  code.Account,
		Password: code.Password,
	}
	codeResponse, err := p.PB.CodeService.PullCode(c, &codeRequest)
	if err != nil {
		p.LG.Error("拉取代码失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	//没拉成功
	if codeResponse.Msg != "" {
		p.LG.Error("拉取代码失败", zap.Error(fmt.Errorf(codeResponse.Msg)))
		utils.ResponseErrorWithMsg(c, utils.CodeServerBusy, err)
		return
	}
	//成功就返回位置及是什么语言
	//走代码检测逻辑

}
