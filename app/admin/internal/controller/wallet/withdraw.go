package wallet

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/admin/api/wallet"
	"star_net/app/admin/internal/service/walletsvc"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xstr"
)

var (
	Withdraw = cWithdraw{}
)

type cWithdraw struct{}

func (c cWithdraw) Create(ctx context.Context, req *wallet.CreateWithdrawReq) (_ *wallet.CreateWithdrawRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "o_withdraw", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cWithdraw) Update(ctx context.Context, req *wallet.UpdateWithdrawReq) (_ *wallet.UpdateWithdrawRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "o_withdraw", Field: "order_no", V: req.OrderNo, Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cWithdraw) Read(ctx context.Context, req *wallet.ReadWithdrawReq) (_ *wallet.ReadWithdrawRes, _ error) {
	var d wallet.ReadWithdrawRes
	x := xcrud.Read{Ctx: ctx, Table: "o_withdraw", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cWithdraw) ReadList(ctx context.Context, req *wallet.ReadListWithdrawReq) (_ *wallet.ReadListWithdrawRes, _ error) {
	var (
		d     = make([]*entity.Withdraw, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "o_withdraw", Page: req.Page, Size: req.Size, Order: "order_no desc", Where: func(db *gdb.Model) {
		if req.Protocol != "" {
			db = db.Where("protocol", req.Protocol)
		}
		if req.OrderNo != "" {
			db = db.WhereLike("order_no", xstr.Like(req.OrderNo))
		}
		if req.Account != "" {
			db = db.WhereLike("account", xstr.Like(req.Account))
		}
		if req.Address != "" {
			db = db.WhereLike("address", xstr.Like(req.Address))
		}
		if req.Net != "" {
			db = db.Where("net", req.Net)
		}
		if req.TransferOrderNo != "" {
			db = db.WhereLike("transfer_order_no", xstr.Like(req.TransferOrderNo))
		}
		if req.Detail != "" {
			db = db.WhereLike("detail", xstr.Like(req.Detail))
		}
		if req.Status != "" {
			db = db.Where("status", req.Status)
		}
		if req.AdminOperate != "" {
			db = db.WhereLike("admin_operate", xstr.Like(req.AdminOperate))
		}
		if req.Note != "" {
			db = db.WhereLike("note", xstr.Like(req.Note))
		}
	}}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &wallet.ReadListWithdrawRes{List: d, Total: total}, nil
}
func (c cWithdraw) Del(ctx context.Context, req *wallet.DelWithdrawReq) (_ *wallet.DelWithdrawRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "o_withdraw", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cWithdraw) CheckOrderByAdmin(ctx context.Context, req *wallet.CheckOrderByAdminReq) (_ *model.CommonRes, _ error) {
	x := walletsvc.Withdraw{
		OrderNo:      req.OrderNo,
		Status:       req.Status,
		StatusRemark: req.StatusRemark,
		SysRemark:    req.SysRemark,
	}
	if err := x.Update(ctx); err != nil {
		return nil, err
	}
	return
}
