package api

import "kubernetes-deploy/controller/handle"

func init() {
	// 注册所有的manager
	handle.CreateDeployManager()
	handle.CreateJobManager()
}
