package pkg

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("password")

type Claims struct {
	Password string `json:"password"`
	UserID   string `json:"userid"`
	jwt.StandardClaims
}

// jwt加密 签发token
func GenToken(userid, password string) (atoken, rtoken string, err error) {
	nowtime := time.Now()
	expire_time := nowtime.Add(24 * time.Hour)
	claims := Claims{
		Password: password,
		UserID:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire_time.Unix(),
			Issuer:    "cong",
		},
	}
	// 加密方式
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	atoken, err = tokenClaims.SignedString(jwtSecret)
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Hour * 24).Unix(),
		Issuer:    "cong0",
	}).SignedString(jwtSecret)
	return
}

// 解密token  服务端识别用户身份
func ParseToken(tokenString string) (*Claims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象的claim进行类型断言
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is not valid")
}

// 使用rtoken刷新atoken
func RefreshToken(atoken, rtoken string) (newtoken, newrtoken string, err error) {
	var cliams Claims
	_, err = jwt.ParseWithClaims(atoken, &cliams, func(atoken *jwt.Token) (interface{}, error) {
		return cliams, nil
	})
	return GenToken(cliams.UserID, cliams.Password)
}
