package controller

import "kubernetes-deploy/controller/handle"

var sourceManager SourceManager

type InitInterface struct {
	handle.DeployHandle
}

type SourceManager interface {
	CommentResource
}

func GetManager() SourceManager {
	return sourceManager
}

func CreateSourceManager() *InitInterface {
	var manager InitInterface
	sourceManager = &manager
	return &manager
}
