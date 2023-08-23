package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type LogsReq struct {
	g.Meta `tags:"钱包" sm:"查询账变记录" method:"get" path:"/logs"`
	model.CommonPageReq
	BalanceCode int `json:"balanceCode" dc:"交易类型 可选 不传查询所有"`
}
type LogsRes struct {
	Total int                 `json:"total"`
	Logs  []*entity.WalletLog `json:"logs"`
}
