package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-soccer/internal/service/ordersvc"
)

type OrderListReq struct {
	g.Meta `tags:"订单" method:"get" path:"/list" dc:"订单列表"`
	ordersvc.OrderList
}
type OrderListRes struct {
	List []ordersvc.OrderListOutput
}
