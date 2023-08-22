package user

import "github.com/gogf/gf/v2/frame/g"

type PutLoginPassReq struct {
	g.Meta  `tags:"用户" sm:"修改登录密码" method:"put" path:"/pass"`
	OldPass string `v:"required"`
	NewPass string `v:"required|password#|密码格式不正确"`
}
