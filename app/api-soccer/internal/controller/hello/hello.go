package hello

import (
	"context"
	"fmt"
	v1 "star_net/app/api-soccer/api/hello/v1"
	"star_net/app/api-soccer/internal/service/order"
)

var (
	Ctrl = &cHello{}
)

type cHello struct {
}

func (cHello) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {

	order.Bet.OddsId = 1
	order.Bet.Exec(ctx)
	return nil, fmt.Errorf("用户名已存在")
}
