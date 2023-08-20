package vpassport

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta    `tags:"/user/register" method:"post" path:"/register" dc:"register"`
	AreaCode  string `dc:"手机区号" json:"areaCode"`
	Email     string `dc:"邮箱" json:"email"`
	Account   string `dc:"账号" json:"account"`
	Phone     string `dc:"手机号" json:"phone"`
	Password  string `dc:"密码" json:"password"`
	Xid       string `dc:"邀请码" json:"xid" v:"required#please input Invitation code"`
	Ip        string `json:"-"`
	UserAgent string `json:"-"`
}
type RegisterRes struct {
	Token string `json:"token"`
}
