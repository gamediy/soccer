package wallet

import (
	"context"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/walletsvc"
)

func (c cWallet) Banks(ctx context.Context, _ *wallet.BanksReq) (_ wallet.BanksRes, _ error) {
	x := walletsvc.Bank{}
	banks := x.GetBanks(ctx)
	return banks, nil
}
