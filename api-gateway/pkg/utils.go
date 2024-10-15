package pkg

import (
	"github.com/google/uuid"
	"regexp"
)

// 生成唯一id
func GenerateUUID() (string, error) {
	uuids, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuids.String(), nil
}

// 定义一个验证邮箱的函数
func IsValidEmail(email string) bool {
	// 邮箱的正则表达式
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
