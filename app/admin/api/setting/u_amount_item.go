package setting

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type ReadAmountItemReq struct {
	g.Meta `tags:"/setting/amountItem" method:"get" path:"/amountItem" dc:"查询AmountItem"`
	Id     uint64 `json:"id"`
}
type ReadAmountItemRes entity.AmountItem
type ReadListAmountItemReq struct {
	g.Meta `tags:"/setting/amountItem" method:"get" path:"/amountItem/list" dc:"查询AmountItem列表"`
	model.CommonPageReq
}
type ReadListAmountItemRes struct {
	Total int                  `json:"total"`
	List  []*entity.AmountItem `json:"list"`
}
type CreateAmountItemReq struct {
	g.Meta `tags:"/setting/amountItem" method:"post" path:"/amountItem" dc:"添加AmountItem"`
	*entity.AmountItem
}
type CreateAmountItemRes model.CommonRes
type UpdateAmountItemReq struct {
	g.Meta `tags:"/setting/amountItem" method:"put" path:"/amountItem" dc:"修改AmountItem"`
	*entity.AmountItem
}
type UpdateAmountItemRes model.CommonRes
type DelAmountItemReq struct {
	g.Meta `tags:"/setting/amountItem" method:"delete" path:"/amountItem" dc:"删除AmountItem"`
	Id     uint64 `json:"id"`
}
type DelAmountItemRes model.CommonRes
