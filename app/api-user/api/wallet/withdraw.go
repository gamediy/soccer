package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateReq struct {
	g.Meta            `tags:"钱包" sm:"创建提现订单" method:"post" path:"/withdraw/create"`
	WithdrawId        int     `json:"withdrawId" v:"required:"`
	Amount            float64 `json:"amount"`
	WithdrawAccountId int     `json:"WithdrawAccount"`
}
type BindWithdrawAccountReq struct {
	g.Meta  `tags:"钱包" sm:"绑定提现账户信息" method:"post" path:"/withdraw/account/bind"`
	BankId  uint64 `v:"required#银行ID不能为空" dc:"银行ID" json:"bankId"`
	Address string `v:"required#收款地址不能为空" json:"address"`
	Title   string `json:"title" dc:"持卡人" v:"required#持卡人不能为空"`
}
type WithdrawAccountListReq struct {
	g.Meta `tags:"钱包" sm:"查询银行卡列表" method:"get" path:"/withdraw/account/list"`
	model.CommonPageReq
}
type WithdrawAccountListRes struct {
	Total int                       `json:"total"`
	List  []*entity.WithdrawAccount `json:"list"`
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
