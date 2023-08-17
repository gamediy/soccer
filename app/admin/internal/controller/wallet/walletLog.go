package wallet

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/admin/api/wallet"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xstr"
)

var Log = cWalletLog{}

type cWalletLog struct{}

func (c cWalletLog) Create(ctx context.Context, req *wallet.CreateWalletLogReq) (_ *wallet.CreateWalletLogRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "u_wallet_log", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cWalletLog) Update(ctx context.Context, req *wallet.UpdateWalletLogReq) (_ *wallet.UpdateWalletLogRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "u_wallet_log", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cWalletLog) Read(ctx context.Context, req *wallet.ReadWalletLogReq) (_ *wallet.ReadWalletLogRes, _ error) {
	var d wallet.ReadWalletLogRes
	x := xcrud.Read{Ctx: ctx, Table: "u_wallet_log", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cWalletLog) ReadList(ctx context.Context, req *wallet.ReadListWalletLogReq) (_ *wallet.ReadListWalletLogRes, _ error) {
	var (
		d     = make([]*entity.WalletLog, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx,
		Table: "u_wallet_log",
		Page:  req.Page, Size: req.Size,
		Where: func(db *gdb.Model) {
			if req.Account != "" {
				db = db.WhereLike("account", xstr.Like(req.Account))
			}
			if req.Note != "" {
				db = db.WhereLike("note", xstr.Like(req.Note))
			}
			if req.OrderNo != "" {
				db = db.WhereLike("order_no", xstr.Like(req.OrderNo))
			}
			if req.BalanceCode != "" {
				db = db.Where("balance_code", req.BalanceCode)
			}
		},
	}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &wallet.ReadListWalletLogRes{List: d, Total: total}, nil
}
func (c cWalletLog) Del(ctx context.Context, req *wallet.DelWalletLogReq) (_ *wallet.DelWalletLogRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "u_wallet_log", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
