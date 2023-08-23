package wallet

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/walletsvc"
)

func (c cWallet) BalanceCodes(ctx context.Context, _ *wallet.BalanceCodesReq) (res wallet.BalanceCodesRes, err error) {
	handler := walletsvc.ListBalanceCodes{
		Lang: ghttp.RequestFromCtx(ctx).GetHeader("lang"),
	}
	return handler.Exec(ctx)
}
