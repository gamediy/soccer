package walletsvc

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/core/wallet"
)

type UpdateWalletByAdmin struct {
	Ctx         context.Context
	Uid         uint64
	Amount      int
	Note        string
	BalanceCode int
}

func (s *UpdateWalletByAdmin) Exec() error {
	x := wallet.BalanceUpdate{
		Uid:         int64(s.Uid),
		Amount:      gconv.Float64(s.Amount),
		BalanceCode: s.BalanceCode,
		Note:        s.Note,
	}
	if err := x.Update(s.Ctx, nil); err != nil {
		return err
	}
	return nil
}
