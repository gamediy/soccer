package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/model"
)

type CreateDepositReq struct {
	g.Meta          `tags:"钱包" sm:"创建充值订单" method:"post" path:"/deposit/create"`
	PayId           int     `json:"payId" v:"required" sm:"充值类型ID" dc:"从 /api/user/wallet/deposit/platform 获取选择"` //充值Id
	Amount          float64 `json:"amount" v:"required"`                                                          //充值金额
	TransferOrderNo string  `json:"transferOrderNo" v:"required" sm:"成功订单号" dc:"用户转账成功的订单号，一般在用户的银行app里面获取"`      //成功订单号
	TranserImg      string  `json:"transerImg" v:"required" sm:"凭证" dc:"上传用户转账的截图"`                               //成功图片
}
type ListPlatformDepositReq struct {
	g.Meta `tags:"钱包" sm:"查询平台充值方式" method:"get" path:"/deposit/platform"`
}
type ListPlatformDepositRes []model.AmountItem
