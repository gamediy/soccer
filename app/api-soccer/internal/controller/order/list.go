package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/api-soccer/api/order"
	"star_net/app/api-soccer/internal/service/ordersvc"
)

func (cOrder) List(ctx context.Context, req *order.OrderListReq) (res *order.OrderListRes, err error) {
	list := ordersvc.OrderList{}
	gconv.Struct(req, &list)
	exec := list.Exec(ctx)
	res.List = exec
	return res, err
}
