package user

import (
	"github.com/gogf/gf/v2/frame/g"
	adminModel "star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"star_net/model"
)

type ReadUserReq struct {
	g.Meta `tags:"/user" method:"get" path:"/" dc:"用户查询"`
	Id     uint64 `json:"id"`
}
type ReadUserRes entity.User
type ReadListUserReq struct {
	g.Meta `tags:"/user" method:"get" path:"/list" dc:"用户查询列表"`
	model.CommonPageReq
}
type ReadListUserRes struct {
	Total int                `json:"total"`
	List  []*adminModel.User `json:"list"`
}
type CreateUserReq struct {
	g.Meta `tags:"/user" method:"post" path:"/" dc:"用户添加"`
	*entity.User
}
type CreateUserRes model.CommonRes
type UpdateUserReq struct {
	g.Meta `tags:"/user" method:"put" path:"/" dc:"用户修改"`
	*entity.User
}
type UpdateUserRes model.CommonRes
type DelUserReq struct {
	g.Meta `tags:"/user" method:"delete" path:"/" dc:"用户删除"`
	Id     uint64 `json:"id"`
}
type DelUserRes model.CommonRes
type PutPassReq struct {
	g.Meta `tags:"/user" method:"put" path:"/pass" dc:"用户修改用户登录密码"`
	Pass   string `json:"pass" v:"required"`
	Id     uint64 `json:"id" v:"required"`
}
