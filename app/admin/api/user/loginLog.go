package user

import (
	"github.com/gogf/gf/v2/frame/g"
	userModel "star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"star_net/model"
)

type ReadUserLoginLogReq struct {
	g.Meta `tags:"/user/loginLog" method:"get" path:"/loginLog" dc:"登录日志查询"`
	Id     uint64 `json:"id"`
}
type ReadUserLoginLogRes entity.UserLoginLog
type ReadListUserLoginLogReq struct {
	g.Meta `tags:"/user/loginLog" method:"get" path:"/loginLog/list" dc:"登录日志查询列表"`
	model.CommonPageReq
}
type ReadListUserLoginLogRes struct {
	Total int                       `json:"total"`
	List  []*userModel.UserLoginLog `json:"list"`
}
type CreateUserLoginLogReq struct {
	g.Meta `tags:"/user/loginLog" method:"post" path:"/loginLog" dc:"登录日志添加"`
	*entity.UserLoginLog
}
type CreateUserLoginLogRes model.CommonRes
type UpdateUserLoginLogReq struct {
	g.Meta `tags:"/user/loginLog" method:"put" path:"/loginLog" dc:"登录日志修改"`
	*entity.UserLoginLog
}
type UpdateUserLoginLogRes model.CommonRes
type DelUserLoginLogReq struct {
	g.Meta `tags:"/user/loginLog" method:"delete" path:"/loginLog" dc:"登录日志删除"`
	Id     uint64 `json:"id"`
}
type DelUserLoginLogRes model.CommonRes
