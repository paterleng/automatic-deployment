package api

import (
	"api-gateway/controller"
	"api-gateway/utils"
)

var apiManager ApiManager

type ApiManager interface {
	UserInterface
	KubernetesApiInterface
}

func CreateApiManager() {
	apiManager = NewManager()
}

func NewManager() *controller.Routes {
	var router controller.Routes
	router.KubernetesController.LG = utils.Tools.LG
	router.KubernetesController.PB = utils.Tools.PB
	return &router

}

func GetManager() ApiManager {
	return apiManager
}
