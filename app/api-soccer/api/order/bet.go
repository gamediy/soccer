package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-soccer/internal/service/ordersvc"
)

type OrderBetReq struct {
	g.Meta `tags:"订单" method:"get" path:"/bet" dc:"投注"`
	ordersvc.BetInput
}
type OrderBetRes struct {
}

type OrderBetBatchReq struct {
	g.Meta `tags:"订单" method:"get" path:"/bet_batch" dc:"投注批量"`
	ordersvc.BetBatchInput
}
type OrderBetBatchRes struct {
	Errors []string `json:"errors" dc:"下注失败返回的信息"`
}
