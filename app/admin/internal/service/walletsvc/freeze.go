package walletsvc

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/core/wallet"
)

type Freeze struct {
	Uid         uint64
	Amount      int
	BalanceCode int
	Title       string
	Note        string
}

func (x *Freeze) Exec(ctx context.Context) error {
	xx := wallet.BalanceUpdate{
		Uid:         int64(x.Uid),
		Amount:      gconv.Float64(x.Amount),
		BalanceCode: x.BalanceCode,
		Title:       x.Title,
		Note:        x.Note,
	}
	if err := xx.Update(ctx, nil); err != nil {
		return err
	}
	return nil
}
