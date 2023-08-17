package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	modelAdmin "star_net/app/admin/internal/model"
	"star_net/model"
)

type ListAdminLoginLogReq struct {
	g.Meta `tags:"/sys/adminLoginLog" method:"get" path:"/adminLoginLog/list" dc:"管理员查询登录日志"`
	model.CommonPageReq
	Uname string `json:"uname"`
	Ip    string `json:"ip"`
}
type ListAdminLoginLogRes struct {
	Total int                         `json:"total"`
	List  []*modelAdmin.AdminLoginLog `json:"list"`
}
