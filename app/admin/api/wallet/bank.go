package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateBankRes model.CommonRes
type UpdateBankReq struct {
	g.Meta `tags:"/wallet/bank" method:"put" path:"/bank" dc:"修改银行"`
	*entity.Bank
}
type UpdateBankRes model.CommonRes
type ReadBankReq struct {
	g.Meta `tags:"/wallet/bank" method:"get" path:"/bank" dc:"查询银行"`
	Id     uint64 `json:"id"`
}
type ReadBankRes entity.Bank
type ReadListBankReq struct {
	g.Meta `tags:"/wallet/bank" method:"get" path:"/bank/list" dc:"查询银行列表"`
	model.CommonPageReq
}
type ReadListBankRes struct {
	Total int            `json:"total"`
	List  []*entity.Bank `json:"list"`
}
type CreateBankReq struct {
	g.Meta `tags:"/wallet/bank" method:"post" path:"/bank" dc:"添加银行"`
	*entity.Bank
}
type DelBankReq struct {
	g.Meta `tags:"/wallet/bank" method:"delete" path:"/bank" dc:"删除银行"`
	Id     uint64 `json:"id"`
}
type DelBankRes model.CommonRes
