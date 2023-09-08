package order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/api-soccer/api/order"
	"star_net/app/api-soccer/internal/service/ordersvc"
)

func (cOrder) Bet(ctx context.Context, req *order.OrderBetReq) (res *order.OrderBetRes, err error) {

	bet := ordersvc.Bet{}
	gconv.Struct(req, &bet.BetInput)
	fmt.Println(&bet.BetInput.Amount)
	err = bet.Exec(ctx)
	return nil, err
}

func (s *cOrder) BetBatch(ctx context.Context, req *order.OrderBetBatchReq) (res *order.OrderBetBatchRes, err error) {

	bet := ordersvc.BetBatch{}
	gconv.Struct(req, &bet.BetBatchInput)
	res = &order.OrderBetBatchRes{}
	exec, err := bet.Exec(ctx)
	res.Errors = exec.Errors
	return res, err
}
