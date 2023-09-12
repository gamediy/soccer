package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/service/withdrawsvc"
	"star_net/model"
)

type CreateWithdrawReq struct {
	g.Meta            `tags:"钱包" sm:"创建提现订单" method:"post" path:"/withdraw/create"`
	AmountItemId      int     `json:"amountItemId" v:"required" dc:"平台支持的银行ID 从 /wallet/deposit/platform 获取选择"`
	Amount            float64 `json:"amount" v:"required" dc:"提现金额"`
	WithdrawAccountId int     `json:"withdrawAccountId" v:"required" dc:"用户银行的ID （从 /withdraw/account/list 接口获取选择）"`
}
type BindWithdrawAccountReq struct {
	g.Meta  `tags:"钱包" sm:"绑定提现账户信息" method:"post" path:"/withdraw/account/bind"`
	BankId  string `v:"required#银行ID不能为空" dc:"银行ID" json:"bankId"`
	Address string `v:"required#收款地址不能为空" json:"address" dc:"卡号"`
	Title   string `json:"title" dc:"持卡人" v:"required#持卡人不能为空"`
}
type WithdrawAccountListReq struct {
	g.Meta `tags:"钱包" sm:"查询银行卡列表" method:"get" path:"/withdraw/account/list"`
	model.CommonPageReq
}
type WithdrawAccountListRes struct {
	Total int                            `json:"total"`
	List  []*withdrawsvc.WithdrawAccount `json:"list"`
}

type SetDefaultWithdrawAccountReq struct {
	g.Meta `tags:"钱包"  sm:"设置默认提款银行卡" method:"put" path:"/withdraw/account/setDefault"`
	Id     int64 `json:"id" v:"required"`
}
type DelWithdrawAccountReq struct {
	g.Meta  `tags:"钱包" sm:"删除提现银行" method:"delete" path:"/withdraw/account/del"`
	Id      int64  `json:"id" v:"required"`
	PayPass string `json:"payPass" dc:"交易密码" v:"required"`
}
