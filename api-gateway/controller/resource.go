package controller

import "api-gateway/controller/kubernetes_deploy"

type Routes struct {
	UserServiceController
	controller.KubernetesController
	controller.SecretController
	controller.NodeController
}
