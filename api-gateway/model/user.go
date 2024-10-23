package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID          string `json:"userid" gorm:"column:userid"`
	UserName        string `json:"username" gorm:"column:username"`
	MailBox         string `json:"mailbox" gorm:"column:mailbox"`
	MailPassWD      string `json:"mailpasswd" gorm:"column:mailpasswd"`
	Atoken          string `json:"atoken" gorm:"column:atoken;type:varchar(512)"`
	Rtoken          string `json:"rtoken" gorm:"column:rtoken;type:varchar(512)"`
	RoleName        string `json:"rolename" gorm:"column:role_name"`
	Role            Role   `json:"role" gorm:"foreignKey:RoleName;references:Name"` // 关联角色表
	Captcha         string `json:"captcha" gorm:"column:captcha"`
	LastLoginDevice string `json:"lastlogindevice" gorm:"column:lastlogindevice; type:varchar(255)"`
}

type Route struct {
	Menu          []Menu   `json:"menu"`
	Permissions   []string `json:"permissions"`
	DashboardGrid []string `json:"dashboardGrid"`
}
type Menu struct {
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Meta     Meta       `json:"meta"`
	Children []Children `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Type  string `json:"type"`
	Affix bool   `json:"affix"`
	Tag   string `json:"tag"`
}

type Children struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Meta      Meta   `json:"meta"`
	Component string `json:"component"`
}

// 角色表
type Role struct {
	gorm.Model
	Name      string      `json:"name" gorm:"column:name; unique"`
	ApiRouter StringSlice `json:"api_router" gorm:"type:json"`
}
type StringSlice []string

// 实现 Scanner 接口，反序列化 JSON 数据
func (s *StringSlice) Scan(value interface{}) error {
	// 检查类型是否为 []uint8，即二进制数据
	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("type assertion to []uint8 failed")
	}

	// 将二进制数据解码为 JSON
	return json.Unmarshal(bytes, s)
}

// 实现 Valuer 接口，将数据序列化为 JSON
func (s StringSlice) Value() (driver.Value, error) {
	// 将 StringSlice 序列化为 JSON 字符串
	return json.Marshal(s)
}

// 设计一个结构体 存储验证码和过期时间
type CaptchaExpire struct {
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// 生成相对应的结构体实例
func GenerateCaptchaExpire(code string, expireTime time.Duration) (CaptchaExpire, error) {
	expirationTime := time.Now().Add(expireTime)
	captchaexpire := CaptchaExpire{
		Code:      code,
		ExpiresAt: expirationTime,
	}
	return captchaexpire, nil
}

const (
	PassWordCost = 12 // 密码加密难度
)

// 用户邮箱密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.MailPassWD = string(bytes)
	return nil
}

func SetPasswords(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return ""
	}
	return string(bytes)

}

// 用户邮箱密码检验
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.MailPassWD), []byte(password))
	if err != nil {
		return false
	}
	return true
}
