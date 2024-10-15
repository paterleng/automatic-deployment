package dao

import (
	"api-gateway/model"
	"gorm.io/gorm"
)

type SecretDao struct {
	DB *gorm.DB
}

func (d *SecretDao) Get(p interface{}) (secrets []model.Secret, err error) {
	secret := p.(model.Secret)
	err = d.DB.Find(&secrets).Where("id = ", secret.ID).Error
	return
}

func (d *SecretDao) Create(p interface{}) error {
	secret := p.(model.Secret)
	err := d.DB.Create(&secret).Error
	return err
}

func (d *SecretDao) Update(p interface{}) error {
	secret := p.(model.Secret)
	err := d.DB.Omit("created_at").Save(&secret).Error
	return err
}

func (d *SecretDao) Delete(p []int) error {
	err := d.DB.Where("id IN ?", p).Delete(&model.Secret{}).Error
	return err
}
