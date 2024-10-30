package api

import "github.com/gin-gonic/gin"

func KubernetesApi(r *gin.Engine) {
	ks := r.Group("/api/kubernetes")
	ks.GET("/config", GetManager().GetConfig)
	ks.POST("/create/resource", GetManager().CreateResource)
	ks.GET("/get/secret", GetManager().GetSecret)
	ks.POST("/create/secret", GetManager().CreateSecret)
	ks.PUT("/put/secret", GetManager().UpdateSecret)
	ks.DELETE("/delete/secret", GetManager().DeleteSecret)
	ks.POST("/docking", GetManager().ClusterDocking)
	ks.GET("/node_info", GetManager().GetNodeInfo)
}
