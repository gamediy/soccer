package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
)

type GetReq struct {
	g.Meta `tags:"/order" method:"get" path:"/" dc:"查询钱包信息"`
	Uid    uint64 `v:"required"`
}
type GetRes entity.Wallet
type UpdateBalanceReq struct {
	g.Meta      `tags:"/order" method:"put" path:"/update" dc:"修改用户钱包"`
	Uid         uint64 `json:"uid" v:"required"`
	Amount      int    `json:"amount" v:"required" dc:"数量"`
	Note        string `json:"note" dc:"说明"`
	BalanceCode int    `json:"balance_code" v:"required"`
}
type FreezeReq struct {
	g.Meta      `tags:"/order" method:"put" path:"/freeze" dc:"冻结用户余额"`
	Uid         uint64 `json:"uid" v:"required"`
	BalanceCode int    `json:"balanceCode" `
	Amount      int    `json:"amount" v:"required" dc:"数量"`
	Note        string `json:"note" v:"required" dc:"说明"`
}
