package vshare

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type AddLanguageReq struct {
	g.Meta `tags:"多语言" method:"post" path:"/language" dc:"添加"`
	entity.Language
}
type GetLanguageReq struct {
	g.Meta `tags:"多语言" method:"get" path:"/language" dc:"查询一条数据"`
	Id     uint64 `v:"required"  json:"id"`
}
type GetLanguageRes struct {
	entity.Language
}

type ListLanguageReq struct {
	g.Meta `tags:"多语言" method:"get" path:"/language/list" dc:"查询列表数据"`
	model.CommonPageReq
}
type ListLanguageRes struct {
	List  []*entity.Language `json:"list"`
	Total int                `json:"total"`
}

type DelLanguageReq struct {
	g.Meta `tags:"多语言" method:"delete" path:"/language" dc:"删除"`
	Id     uint64 `v:"required" json:"id"`
}

type UpdateLanguageReq struct {
	g.Meta `tags:"多语言" method:"put" path:"/language" dc:"修改"`
	entity.Language
}
