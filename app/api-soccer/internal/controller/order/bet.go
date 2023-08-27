package order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/api-soccer/api/order"
	"star_net/app/api-soccer/internal/service/ordersvc"
)

func (cOrder) Bet(ctx context.Context, req *order.OrderBetReq) (res *order.OrderBetRes, err error) {

	bet := ordersvc.Bet
	gconv.Struct(req, &bet.BetInput)
	fmt.Println(&bet.BetInput.Amount)
	err = bet.Exec(ctx)
	return nil, err
}
