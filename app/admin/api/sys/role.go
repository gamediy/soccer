package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	model2 "star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"star_net/model"
)

type GetRoleByIdReq struct {
	g.Meta `tags:"/sys/role" path:"/role" method:"get" dc:"角色查询根据ID"`
	Id     uint64
}
type GetRoleByIdRes struct {
	*entity.Role
	Menu []*entity.RoleMenu `json:"menu"`
}

type ListRoleReq struct {
	g.Meta `tags:"/sys/role" path:"/role/list"  method:"get" dc:"角色列表"`
}
type ListRoleRes struct {
	Total int            `json:"total"`
	List  []*entity.Role `json:"list"`
}
type AddRoleReq struct {
	g.Meta `tags:"/sys/role" path:"/role"  method:"post" dc:"角色添加"`
	Role   model2.Role `json:"role"`
}
type AddRoleRes model.CommonRes

type DelRoleReq struct {
	g.Meta `tags:"/sys/role" path:"/role"  method:"delete" dc:"角色删除"`
	Id     int `json:"id"`
}
type DelRoleRes model.CommonRes

type EditRoleReq struct {
	g.Meta `tags:"/sys/role" path:"/role"  method:"put" dc:"角色修该"`
	Role   model2.Role `json:"role"`
}
type EditRoleRes model.CommonRes
