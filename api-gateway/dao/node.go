package dao

import (
	"api-gateway/model"
	"api-gateway/utils"
	"gorm.io/gorm"
)

type NodeDao struct {
	DB *gorm.DB
}

var nodeManager NodeInterface

type NodeInterface interface {
	KubernetesDao
	Get() ([]model.Cluster, error)
}

type NodeManager struct {
}

func GetNodeManager() NodeInterface {
	return nodeManager
}

func NewNodeManager() {
	var dao *NodeDao
	dao.DB = utils.Tools.DB
	nodeManager = dao
}

func (d *NodeDao) Get() (cluster []model.Cluster, err error) {
	err = d.DB.Find(&cluster).Error
	return
}

func (d *NodeDao) Create(p interface{}) error {
	cluster := p.(model.Cluster)
	err := d.DB.Create(&cluster).Error
	return err
}

func (d *NodeDao) Update(p interface{}) error {
	return nil
}

func (d *NodeDao) Delete(p []int) error {
	return nil
}
