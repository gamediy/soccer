package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `tags:"用户" method:"post" path:"/login/demo"  sm:"登录" dc:"注意：登录地址为 /api/user/user/login" `
	Account  string `v:"required" dc:"用户名" json:"account" json:"account"`
	Password string `v:"required" dc:"密码" json:"password" json:"password"`
}
type LoginRes struct {
	Token string `json:"token"`
}
