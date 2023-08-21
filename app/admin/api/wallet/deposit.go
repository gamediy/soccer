package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	adminModel "star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateDepositRes model.CommonRes
type UpdateDepositReq struct {
	g.Meta `tags:"/wallet/deposit" method:"put" path:"/deposit" dc:"修改提现订单"`
	*adminModel.Deposit
}
type UpdateDepositRes model.CommonRes
type ReadDepositReq struct {
	g.Meta  `tags:"/wallet/deposit" method:"get" path:"/deposit" dc:"查询提现订单"`
	OrderNo string `json:"orderNo"`
}
type ReadDepositRes adminModel.Deposit
type ReadListDepositReq struct {
	g.Meta `tags:"/wallet/deposit" method:"get" path:"/deposit/list" dc:"查询提现订单列表"`
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
}
type ReadListDepositRes struct {
	Total int                   `json:"total"`
	List  []*adminModel.Deposit `json:"list"`
}
type CreateDepositReq struct {
	g.Meta `tags:"/wallet/deposit" method:"post" path:"/deposit" dc:"添加提现订单"`
	*entity.Deposit
}
type DelDepositReq struct {
	g.Meta  `tags:"/wallet/deposit" method:"delete" path:"/deposit" dc:"删除提现订单"`
	OrderNo string `json:"orderNo"`
}
type CheckDepositByAdminReq struct {
	g.Meta       `tags:"/wallet/deposit" method:"put" path:"/deposit/checkOrder" dc:"审核充值订单"`
	OrderNo      int64 `v:"required"`
	Status       int   `v:"required"`
	StatusRemark string
	SysRemark    string
}
type DelDepositRes model.CommonRes
