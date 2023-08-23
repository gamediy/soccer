package wallet

import (
	"context"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/depositsvc"
	"star_net/model"
)

func (c cWallet) CreateDeposit(ctx context.Context, req *wallet.CreateDepositReq) (_ model.CommonRes, _ error) {
	return
}
func (c cWallet) ListPlatformDeposit(ctx context.Context, _ *wallet.ListPlatformDepositReq) (_ wallet.ListPlatformDepositRes, err error) {
	data, err := depositsvc.GetDepositItem.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return data.Item, nil
}
