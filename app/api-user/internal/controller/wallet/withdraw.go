package wallet

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/app/api-user/api/wallet"
	"star_net/app/api-user/internal/service/usersvc"
	"star_net/app/api-user/internal/service/withdrawsvc"
	"star_net/model"
)

func (c cWallet) BindWithdrawAccount(ctx context.Context, req *wallet.BindWithdrawAccountReq) (res *model.CommonRes, err error) {
	x := withdrawsvc.BindWithdrawAccount{
		BankId:  req.BankId,
		Address: req.Address,
		Title:   req.Title,
		Pass:    req.Pass,
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
func (c cWallet) SetDefaultWithdrawAccount(ctx context.Context, req *wallet.SetDefaultWithdrawAccountReq) (res *model.CommonRes, err error) {
	x := withdrawsvc.SetDefaultWithdrawAccount{Id: req.Id}
	err = x.Exec(ctx)
	return
}
func (c cWallet) DelWithdrawAccount(ctx context.Context, req *wallet.DelWithdrawAccountReq) (res *model.CommonRes, err error) {
	x := withdrawsvc.DelWithdrawAccount{Id: req.Id, PayPass: req.PayPass}
	if err = x.Exec(ctx); err != nil {
		return nil, err
	}
	return
}

func (c cWallet) CreateWithdraw(ctx context.Context, req *wallet.CreateWithdrawReq) (res *model.CommonRes, err error) {
	x := withdrawsvc.CreateWithdraw{
		Amount:            req.Amount,
		AmountItemId:      req.AmountItemId,
		WithdrawAccountId: req.WithdrawAccountId,
		Lang:              ghttp.RequestFromCtx(ctx).GetHeader("lang"),
	}
	err = x.Exec(ctx)
	return
}
func (c cWallet) ChangePayPass(ctx context.Context, req *wallet.ChangePayPassReq) (res *model.CommonRes, err error) {
	x := usersvc.ChangePayPass{
		PayPass: req.PayPass,
		Pass:    req.Pass,
		Title:   req.Title,
		Address: req.Address,
	}
	err = x.Exec(ctx)
	return
}
