package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

// GetApiReq 查询单个
type GetApiReq struct {
	g.Meta `tags:"/sys/api" method:"get" path:"/api" dc:"API查询单个"`
	Id     uint64 `json:"id"`
}

type GetApiRes entity.Api

// ListApiReq 查询列表
type ListApiReq struct {
	g.Meta `tags:"/sys/api" method:"get" path:"/api/list" dc:"API查询列表"`
	model.CommonPageReq
	Group  string
	Method string
	Url    string
	Desc   string
}
type ListApiRes struct {
	Total int           `json:"total"`
	List  []*entity.Api `json:"list"`
}

// AddApiReq 添加
type AddApiReq struct {
	g.Meta `tags:"/sys/api" method:"post" path:"/api" dc:"API添加"`
	*entity.Api
}
type AddApiRes model.CommonRes

// UpdateApiReq 修改
type UpdateApiReq struct {
	g.Meta `tags:"/sys/api" method:"put" path:"/api" dc:"API修改单个数据"`
	*entity.Api
}
type UpdateApiRes model.CommonRes

// DelApiReq 删除
type DelApiReq struct {
	g.Meta `tags:"/sys/api" method:"delete" path:"/api" dc:"API删除单个数据"`
	Id     uint64 `json:"id"`
}

type DelApiRes model.CommonRes
