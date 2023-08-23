package wallet

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/admin/api/wallet"
	adminModel "star_net/app/admin/internal/model"
	"star_net/app/admin/internal/service/walletsvc"
	"star_net/consts"
	"star_net/model"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xstr"
)

var (
	Deposit = cDeposit{}
)

type cDeposit struct{}

func (c cDeposit) Create(ctx context.Context, req *wallet.CreateDepositReq) (_ *wallet.CreateDepositRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "o_deposit", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cDeposit) Update(ctx context.Context, req *wallet.UpdateDepositReq) (_ *wallet.UpdateDepositRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "o_deposit", Data: req, Field: "order_no", V: req.OrderNo}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cDeposit) Read(ctx context.Context, req *wallet.ReadDepositReq) (_ *wallet.ReadDepositRes, _ error) {
	var d wallet.ReadDepositRes
	x := xcrud.Read{Ctx: ctx, Table: "o_deposit", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cDeposit) ReadList(ctx context.Context, req *wallet.ReadListDepositReq) (_ *wallet.ReadListDepositRes, _ error) {
	var (
		d     = make([]*adminModel.Deposit, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "o_deposit", Page: req.Page, Size: req.Size, Order: "order_no desc", Where: func(db *gdb.Model) {
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
	}}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	for _, i := range d {
		if i.TransferImg != "" {
			i.TransferImg = consts.ImgPrefix + i.TransferImg
		}
	}
	return &wallet.ReadListDepositRes{List: d, Total: total}, nil
}
func (c cDeposit) Del(ctx context.Context, req *wallet.DelDepositReq) (_ *wallet.DelDepositRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "o_deposit", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (cDeposit) CheckOrderByAdmin(ctx context.Context, req *wallet.CheckDepositByAdminReq) (_ *model.CommonRes, _ error) {
	x := walletsvc.Deposit{
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
