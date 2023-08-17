package setting

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type ReadAmountCategoryReq struct {
	g.Meta `tags:"/setting/amountCategory" method:"get" path:"/amountCategory" dc:"查询AmountCategory"`
	Id     uint64 `json:"id"`
}
type ReadAmountCategoryRes entity.AmountCategory
type ReadListAmountCategoryReq struct {
	g.Meta `tags:"/setting/amountCategory" method:"get" path:"/amountCategory/list" dc:"查询AmountCategory列表"`
	model.CommonPageReq
}
type ReadListAmountCategoryRes struct {
	Total int                      `json:"total"`
	List  []*entity.AmountCategory `json:"list"`
}
type CreateAmountCategoryReq struct {
	g.Meta `tags:"/setting/amountCategory" method:"post" path:"/amountCategory" dc:"添加AmountCategory"`
	*entity.AmountCategory
}
type CreateAmountCategoryRes model.CommonRes
type UpdateAmountCategoryReq struct {
	g.Meta `tags:"/setting/amountCategory" method:"put" path:"/amountCategory" dc:"修改AmountCategory"`
	*entity.AmountCategory
}
type UpdateAmountCategoryRes model.CommonRes
type DelAmountCategoryReq struct {
	g.Meta `tags:"/setting/amountCategory" method:"delete" path:"/amountCategory" dc:"删除AmountCategory"`
	Id     uint64 `json:"id"`
}
type DelAmountCategoryRes model.CommonRes
