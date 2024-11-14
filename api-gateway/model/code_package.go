package model

type CodeRepository struct {
	URL      string `json:"url"`      //路径
	Branch   string `json:"branch"`   //分支
	Account  string `json:"account"`  //认证账号
	Password string `json:"password"` //密码
}
