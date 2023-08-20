package order

import (
	"context"
	"star_net/app/api-soccer/internal/model"
)

var (
	Service = &order{}
)

type order struct {
}

type Order interface {
	Bet(ctx context.Context, input model.BetInput) error
}
