package pkg

import (
	"crypto/rand"
	"gopkg.in/gomail.v2"
	"math/big"
)

type Email struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	BodyType string `json:"bodytype"`
	Body     string `json:"body"`
}

// 生成随机验证码功能 字母数字大小写
func GenerateCaptcha(length int) (string, error) {
	//const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const charset = "0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

// 邮箱内容发送
func (email *Email) SendEmail() error {
	// 设置 Mail SMTP 服务器的地址和端口
	smtpHost := "smtp.qq.com"
	smtpPort := 465
	smtpUsername := "207555435@qq.com" // 你的 Gmail 地址
	smtpPassword := "jcxyhjfykcjlbgbd" // 你的应用专用密码

	// 设置邮箱内容
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody(email.BodyType, email.Body)

	// 发送邮件
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
