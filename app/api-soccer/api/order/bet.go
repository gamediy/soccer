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
