package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta   `tags:"/user/register" method:"post" path:"/register" dc:"注册"`
	Account  string `dc:"账号" json:"account"`
	Password string `dc:"密码" json:"password"`
	Xid      string `dc:"邀请码" json:"xid" `
	Country  string `json:"country" dc:"国家"`
	City     string `json:"city" dc:"城市"`
	Phone    string `json:"phone" dc:"电话号码" v:"required"`
	Email    string `json:"email" dc:"有效"`
}
