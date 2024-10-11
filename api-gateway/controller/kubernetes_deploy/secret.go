package controller

import (
	"api-gateway/dao"
	"api-gateway/model"
	"api-gateway/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SecretController struct {
	PB *utils.Pb
	LG *zap.Logger
}

func (s *SecretController) GetSecret(c *gin.Context) {
	var p model.Secret
	secrets, err := dao.GetSecretManager().Get(p)
	if err != nil {
		s.LG.Error("创建失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, secrets)
}

func (s *SecretController) CreateSecret(c *gin.Context) {
	var p model.Secret
	if err := c.ShouldBindJSON(&p); err != nil {
		s.LG.Error("参数绑定失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	err := dao.GetSecretManager().Create(p)
	if err != nil {
		s.LG.Error("创建失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}

func (s *SecretController) UpdateSecret(c *gin.Context) {
	var p model.Secret
	if err := c.ShouldBindJSON(&p); err != nil {
		s.LG.Error("参数绑定失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	err := dao.GetSecretManager().Update(p)
	if err != nil {
		s.LG.Error("创建失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)

}

func (s *SecretController) DeleteSecret(c *gin.Context) {
	var p model.SecretReq
	if err := c.ShouldBindJSON(&p); err != nil {
		s.LG.Error("参数绑定失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeInvalidParam)
		return
	}
	err := dao.GetSecretManager().Delete(p.Ids)
	if err != nil {
		s.LG.Error("创建失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}
