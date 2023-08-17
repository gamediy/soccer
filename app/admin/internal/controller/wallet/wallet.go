package wallet

import (
	"context"
	"star_net/app/admin/api/wallet"
	"star_net/app/admin/internal/service/walletsvc"
	"star_net/model"
	"star_net/utility/utils/xcrud"
)

var Wallet = cWallet{}

type cWallet struct{}

func (c cWallet) Get(ctx context.Context, req *wallet.GetReq) (_ *wallet.GetRes, _ error) {
	var d wallet.GetRes
	x := xcrud.Read{Ctx: ctx, Table: "u_wallet", Field: "uid", V: req.Uid}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cWallet) Update(ctx context.Context, req *wallet.UpdateBalanceReq) (_ *model.CommonRes, _ error) {
	x := walletsvc.UpdateWalletByAdmin{
		Ctx:         ctx,
		Uid:         req.Uid,
		Note:        req.Note,
		Amount:      req.Amount,
		BalanceCode: req.BalanceCode,
	}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cWallet) Freeze(ctx context.Context, req *wallet.FreezeReq) (_ *model.CommonRes, _ error) {
	x := walletsvc.Freeze{
		Uid:         req.Uid,
		Amount:      req.Amount,
		BalanceCode: req.BalanceCode,
		Note:        req.Note,
	}
	if err := x.Exec(ctx); err != nil {
		return nil, err
	}
	return
}
