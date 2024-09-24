package api

var apiManager ApiManager

type ApiManager interface {
	UserInterface
	KubernetesApiInterface
}

func GetManager() ApiManager {
	return apiManager
}
