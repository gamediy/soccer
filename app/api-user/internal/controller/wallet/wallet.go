package wallet

import (
	"context"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/walletsvc"
)

var (
	Wallet = cWallet{}
)

type cWallet struct{}

func (c cWallet) Logs(ctx context.Context, req *wallet.LogsReq) (res *wallet.LogsRes, err error) {
	handler := walletsvc.Logs{
		Page:        req.Page,
		Size:        req.Size,
		BalanceCode: req.BalanceCode,
	}
	total, walletLogs, err := handler.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &wallet.LogsRes{
		Total: total,
		Logs:  walletLogs,
	}, nil
}
