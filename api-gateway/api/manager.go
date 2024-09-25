package api

import "api-gateway/controller"

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
	return &router

}

func GetManager() ApiManager {
	return apiManager
}
