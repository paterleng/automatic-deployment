package dao

import (
	"api-gateway/model"
	"api-gateway/utils"
)

type KubernetesDao interface {
	Get(interface{}) ([]model.Secret, error)
	Create(interface{}) error
	Update(interface{}) error
	Delete([]int) error
}

var secretManager SecretManager

type SecretManager interface {
	KubernetesDao
}

func CreateSecretManager() {
	secretManager = NewSecretManager()
}

func NewSecretManager() *SecretDao {
	var dao SecretDao
	dao.DB = utils.Tools.DB
	return &dao
}

func GetSecretManager() SecretManager {
	return secretManager
}
