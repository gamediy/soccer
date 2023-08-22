package user

import "github.com/gogf/gf/v2/frame/g"

type PutLoginPassReq struct {
	g.Meta  `tags:"用户" sm:"修改登录密码" method:"put" path:"/pass"`
	OldPass string `v:"required" json:"oldPass"`
	NewPass string `v:"required|password#|密码格式不正确" json:"newPass"`
}
type SetPayPassReq struct {
	g.Meta `tags:"用户" sm:"设置交易密码" method:"put"  path:"/setPayPass"`
	Pass   string `v:"required|integer|size:6#交易密码必填|纯数字|6位" json:"pass"`
}
