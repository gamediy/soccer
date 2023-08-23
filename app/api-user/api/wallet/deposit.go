package wallet

import "github.com/gogf/gf/v2/frame/g"

type CreateDepositReq struct {
	g.Meta          `tags:"钱包" sm:"创建充值订单" method:"post" path:"/deposit/create"`
	PayId           int     `json:"payId" v:"required"`           //充值Id
	Amount          float64 `json:"amount" v:"required"`          //充值金额
	TransferOrderNo string  `json:"transferOrderNo" v:"required"` //成功订单号
	TranserImg      string  `json:"transerImg" v:"required"`      //成功图片
}
