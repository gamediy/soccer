package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta   `tags:"用户" method:"post" path:"/register" dc:"注册"`
	Account  string `dc:"账号" json:"account" v:"required|password#-200|-200"`
	Password string `dc:"密码" json:"password" v:"required|password#-201|-201"`
	Xid      string `dc:"邀请码" json:"xid"`
	Country  string `json:"country" dc:"国家" v:"required"`
	City     string `json:"city" dc:"城市" v:"required"`
	Phone    string `json:"phone" dc:"电话号码" v:"required"`
	Email    string `json:"email" dc:"有效"`
	RealName string `json:"realName" dc:"真实姓名" v:"required"`
}
