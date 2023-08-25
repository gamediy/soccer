package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type ListAdminLoginLogReq struct {
	g.Meta `tags:"/sys/adminLoginLog" method:"get" path:"/adminLoginLog/list" dc:"管理员查询登录日志"`
	model.CommonPageReq
	Account string `json:"account"`
	Ip      string `json:"ip"`
}
type ListAdminLoginLogRes struct {
	Total int                     `json:"total"`
	List  []*entity.AdminLoginLog `json:"list"`
}
