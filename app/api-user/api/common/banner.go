package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
)

type BannersReq struct {
	g.Meta `tags:"通用" sm:"banners" path:"/banners"`
}
type BannersRes []*entity.Banner
