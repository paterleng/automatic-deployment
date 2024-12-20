package api

import (
	"api-gateway/controller"
	"api-gateway/utils"
)

var apiManager ApiManager

type ApiManager interface {
	UserInterface
	CodePackage
	KubernetesApiInterface
}

func CreateApiManager() {
	apiManager = NewManager()
}

func NewManager() *controller.Routes {
	var router controller.Routes
	router.KubernetesController.LG = utils.Tools.LG
	router.KubernetesController.PB = utils.Tools.PB
	router.SecretController.LG = utils.Tools.LG
	router.SecretController.PB = utils.Tools.PB
	router.NodeController.LG = utils.Tools.LG
	router.NodeController.PB = utils.Tools.PB
	router.CodePackageController.LG = utils.Tools.LG
	router.CodePackageController.PB = utils.Tools.PB
	return &router

}

func GetManager() ApiManager {
	return apiManager
}
