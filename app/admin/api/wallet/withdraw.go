package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateWithdrawRes model.CommonRes
type UpdateWithdrawReq struct {
	g.Meta `tags:"/wallet/withdraw" method:"put" path:"/withdraw" dc:"修改提现订单"`
	*entity.Withdraw
}
type UpdateWithdrawRes model.CommonRes
type ReadWithdrawReq struct {
	g.Meta  `tags:"/wallet/withdraw" method:"get" path:"/withdraw" dc:"查询提现订单"`
	OrderNo string
}
type ReadWithdrawRes entity.Withdraw
type ReadListWithdrawReq struct {
	g.Meta `tags:"/wallet/withdraw" method:"get" path:"/withdraw/list" dc:"查询提现订单列表"`
	model.CommonPageReq
	OrderNo         string
	Account         string
	Address         string
	Net             string
	TransferOrderNo string
	Detail          string
	Protocol        string
	Status          string
	AdminOperate    string
	Note            string
}
type ReadListWithdrawRes struct {
	Total int                `json:"total"`
	List  []*entity.Withdraw `json:"list"`
}
type CreateWithdrawReq struct {
	g.Meta `tags:"/wallet/withdraw" method:"post" path:"/withdraw" dc:"添加提现订单"`
	*entity.Withdraw
}
type DelWithdrawReq struct {
	g.Meta  `tags:"/wallet/withdraw" method:"delete" path:"/withdraw" dc:"删除提现订单"`
	OrderNo string
}
type DelWithdrawRes model.CommonRes

type CheckOrderByAdminReq struct {
	g.Meta       `tags:"/wallet/order" method:"put" path:"/withdraw/checkOrder" dc:"审核提现订单"`
	OrderNo      int64 `v:"required"`
	Status       int   `v:"required"`
	StatusRemark string
	SysRemark    string
}
