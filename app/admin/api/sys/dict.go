package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

// GetDictReq 查询单个
type GetDictReq struct {
	g.Meta `tags:"/sys/dict" method:"get" path:"/dict" dc:"字典查询单个"`
	Id     uint64 `json:"id"`
}
type GetDictRes entity.Dict

// ListDictReq 查询列表
type ListDictReq struct {
	g.Meta `tags:"/sys/dict" method:"get" path:"/dict/list" dc:"字典查询列表"`
	model.CommonPageReq
}
type ListDictRes struct {
	Total int            `json:"total"`
	List  []*entity.Dict `json:"list"`
}

// AddDictReq 添加
type AddDictReq struct {
	g.Meta `tags:"/sys/dict" method:"post" path:"/dict" dc:"字典添加"`
	Title  string `json:"title"    v:"required" `
	K      string `json:"k"        v:"required" `
	V      string `json:"v"        v:"required" `
	Desc   string `json:"desc"      `
	Group  string `json:"group"    v:"required" `
	Status int    `json:"status"    `
	Type   int    `json:"type"      description:"1 文本，2 img"`
}
type AddDictRes model.CommonRes

// UpdateDictReq 修改
type UpdateDictReq struct {
	g.Meta `tags:"/sys/dict" method:"put" path:"/dict" summary:"字典修改单个" dc:"字典 key 不会被修改"`
	*entity.Dict
}
type UpdateDictRes model.CommonRes

// DelDictReq 删除
type DelDictReq struct {
	g.Meta `tags:"/sys/dict" method:"delete" path:"/dict" dc:"字典删除单个"`
	Id     uint64 `json:"id"`
}

type DelDictRes model.CommonRes
