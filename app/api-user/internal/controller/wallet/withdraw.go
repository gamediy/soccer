package wallet

import (
	"context"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/withdrawsvc"
	"star_net/model"
)

func (c cWallet) BindWithdrawAccount(ctx context.Context, req *wallet.BindWithdrawAccountReq) (res *model.CommonRes, err error) {
	x := withdrawsvc.BindWithdrawAccount{
		BankId:  req.BankId,
		Address: req.Address,
		Title:   req.Title,
	}
	err = x.Exec(ctx)
	return
}
func (c cWallet) ListWithdrawAccount(ctx context.Context, req *wallet.WithdrawAccountListReq) (res *wallet.WithdrawAccountListRes, err error) {
	x := withdrawsvc.WithdrawAccountList{
		Page: req.Page,
		Size: req.Size,
	}
	total, accounts, err := x.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &wallet.WithdrawAccountListRes{
		Total: total,
		List:  accounts,
	}, nil
}
