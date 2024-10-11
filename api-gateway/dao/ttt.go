package dao

import "gorm.io/gorm"

type Aaa struct {
	DB *gorm.DB
}

var aManager SecretManager

type AManager interface {
	KubernetesDao
}

func CreateAManager() {
	//aManager = NewAManager()
}

func NewAManager() *Aaa {
	var dao Aaa
	return &dao
}

func GetAManager() SecretManager {
	return aManager
}

func (p Aaa) Create(a interface{}) error {

	return nil
}
