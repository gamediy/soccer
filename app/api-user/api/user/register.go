package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta   `tags:"/user/register" method:"post" path:"/register" dc:"注册"`
	Account  string `dc:"账号" json:"account"`
	Phone    string `dc:"手机号" json:"phone"`
	Password string `dc:"密码" json:"password"`
	Xid      string `dc:"邀请码" json:"xid" v:"required#please input Invitation code"`
	Country  string `json:"country" dc:"国家"`
	City     string `json:"city" dc:"城市"`
}
type RegisterRes struct {
	Token string `json:"token"`
}
