package model

import "gorm.io/gorm"

type ServiceReq struct {
	Name      string
	NameSpace string
	Port      int
}

type Secret struct {
	gorm.Model
	Name     string `json:"name"`
	Registry string `json:"registry"` //镜像仓库地址
	Platform string `json:"platform"`
	Account  string `json:"account"`
	PassWord string `json:"pass_word"`
	UserId   string `json:"user_id"`
}
