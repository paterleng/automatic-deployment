package controller

import "api-gateway/controller/kubernetes_deploy"

type Routes struct {
	UserServiceController
	CodePackageController
	controller.KubernetesController
	controller.SecretController
	controller.NodeController
}
