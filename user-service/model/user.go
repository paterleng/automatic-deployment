package model

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"user-service/utils"
)

type User struct {
	UserID     string `json:"column:userid" binding:"required"`
	UserName   string `json:"column:username" binding:"required"`
	MailBox    string `json:"column:mailbox" binding:"required, email"`
	MailPassWD string `json:"column:mailpasswd" binding:"required, min=6"`
	Atoken     string `json:"column:atoken"`
	Rtoken     string `json:"column:rtoken"`
	Role       string `json:"column:role" binding:"required"`
}

const (
	PassWordCost = 12 // 密码加密难度
)

// 用户邮箱密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		utils.Tools.LG.Error("用户密码加密失败", zap.Error(err))
		return err
	}
	user.MailPassWD = string(bytes)
	return nil
}

// 用户邮箱密码检验
func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.MailPassWD), []byte(password))
	return err
}
