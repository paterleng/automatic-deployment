package model

import "gorm.io/gorm"

type ResourceReq struct {
	ResourceType string            `json:"resource_type" validate:"required"`
	Name         string            `json:"name"`
	NameSpace    string            `json:"name_space"`
	ImageName    string            `json:"image_name" validate:"required"`
	Replicas     int32             `json:"replicas"`
	Labels       map[string]string `json:"labels"`
	MatchLabels  map[string]string `json:"match_labels"`
}

type Secret struct {
	gorm.Model
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Account  string `json:"account"`
	PassWord string `json:"pass_word"`
	UserId   string `json:"user_id"`
}

type SecretReq struct {
	Ids []int `json:"ids"`
}

type ClusterResquest struct {
	Name       string `json:"name"`
	ClusterAdr string `json:"cluster_adr"`
}

type Cluster struct {
	gorm.Model
	Name       string `json:"name"`
	ClusterAdr string `json:"cluster_adr"` //集群地址
	Version    string `json:"version"`
	Status     string `json:"status"`   //状态
	OsImage    string `json:"os_image"` //系统镜像版本
	Arm        string `json:"arm"`      //架构
	Config     string `json:"config"`
}
