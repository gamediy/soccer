package wallet

import (
	"context"
	"star_net/app/admin/api/wallet"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	Bank = cBank{}
)

type cBank struct{}

func (c cBank) Create(ctx context.Context, req *wallet.CreateBankReq) (_ *wallet.CreateBankRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "b_bank", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cBank) Update(ctx context.Context, req *wallet.UpdateBankReq) (_ *wallet.UpdateBankRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "b_bank", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cBank) Read(ctx context.Context, req *wallet.ReadBankReq) (_ *wallet.ReadBankRes, _ error) {
	var d wallet.ReadBankRes
	x := xcrud.Read{Ctx: ctx, Table: "b_bank", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cBank) ReadList(ctx context.Context, req *wallet.ReadListBankReq) (_ *wallet.ReadListBankRes, _ error) {
	var (
		d     = make([]*entity.Bank, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "b_bank", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &wallet.ReadListBankRes{List: d, Total: total}, nil
}
func (c cBank) Del(ctx context.Context, req *wallet.DelBankReq) (_ *wallet.DelBankRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "b_bank", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
