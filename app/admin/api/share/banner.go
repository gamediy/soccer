package vshare

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

// AddBannerReq 添加
type AddBannerReq struct {
	g.Meta `tags:"banner" method:"post" path:"/banner" dc:"添加banner"`
	entity.Banner
}

// GetBannerReq 获取
type GetBannerReq struct {
	g.Meta `tags:"banner" method:"get" path:"/banner" dc:"根据id获取banner"`
	Id     uint64 `v:"required" json:"id"`
}
type GetBannerRes struct {
	g.Meta `tags:"后台"`
	entity.Banner
}

// ListBannerReq 集合
type ListBannerReq struct {
	g.Meta `tags:"banner" method:"get" path:"/banner/list" dc:"获取banner列表"`
	model.CommonPageReq
}
type ListBannerRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.Banner `json:"list"`
	model.CommonPageRes
	Total int `json:"total"`
}

// DelBannerReq 删除
type DelBannerReq struct {
	g.Meta `tags:"banner" method:"delete" path:"/banner" dc:"删除banner"`
	Id     uint64 `v:"required" json:"id"`
}

// UpdateBannerReq 修改
type UpdateBannerReq struct {
	g.Meta `tags:"banner" method:"put" path:"/banner" dc:"更新banner"`
	entity.Banner
}
