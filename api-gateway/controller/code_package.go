package controller

import (
	"api-gateway/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"go.uber.org/zap"
	"os"
)

type CodePackageController struct {
	PB *utils.Pb
	LG *zap.Logger
}

func (cp *CodePackageController) CloneCode(c *gin.Context) {
	var a transport.AuthMethod
	git.PlainClone("/code", false, &git.CloneOptions{
		URL:           "https://github.com/paterleng/bug-notify.git",
		Auth:          a,
		ReferenceName: "main",
		Progress:      os.Stdout,
	})
}
