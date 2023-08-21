package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta      `tags:"/user/login" method:"post" path:"/login" dc:"login"`
	Account     string `v:"required" dc:"用户名" json:"account"`
	Password    string `v:"required" dc:"密码" json:"password"`
	CaptchaId   string `v:"required" dc:"验证码ID" json:"captchaId"`
	CaptchaCode string `v:"required" dc:"验证码" json:"captchaCode"`
}
type LoginRes struct {
	Token string `json:"token"`
}
