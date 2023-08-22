package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta   `tags:"用户" method:"post" path:"/register" sm:"注册" dc:"注册完成后，请访问一下登录接口，自动登录一下哈.. 框架受限不能返回token"`
	Account  string `dc:"账号" json:"account" v:"required|password#账户必填|账户格式不正确"`
	Password string `dc:"密码" json:"password" v:"required|password#密码必填|密码格式不正确"`
	Xid      string `dc:"邀请码" json:"xid"`
	Country  string `json:"country" dc:"国家" v:"required"`
	City     string `json:"city" dc:"城市" v:"required"`
	Phone    string `json:"phone" dc:"电话号码" v:"required"`
	Email    string `json:"email" dc:"有效" v:"email#邮箱格式不正确"`
	RealName string `json:"realName" dc:"真实姓名" v:"required"`
}
