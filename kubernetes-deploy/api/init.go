package api

import (
	"go.uber.org/zap"
	"kubernetes-deploy/controller/handle"
	"kubernetes-deploy/utils"
)

func init() {
	// 注册所有的manager
	err := handle.CreateDeployManager()
	if err != nil {
		utils.Tools.LG.Error("初始化deployment manager失败", zap.Error(err))
		return
	}
	err = handle.CreateServiceManager()
	if err != nil {
		utils.Tools.LG.Error("初始化service manager失败", zap.Error(err))
		return
	}
	handle.CreateJobManager()
}
