package setting

import (
	"context"
	"star_net/app/admin/api/setting"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	BalanceCode = cBalanceCode{}
)

type cBalanceCode struct{}

func (c cBalanceCode) Read(ctx context.Context, req *setting.ReadBalanceCodeReq) (_ *setting.ReadBalanceCodeRes, _ error) {
	var d setting.ReadBalanceCodeRes
	x := xcrud.Read{Ctx: ctx, Table: "u_balance_code", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cBalanceCode) ReadList(ctx context.Context, req *setting.ReadListBalanceCodeReq) (_ *setting.ReadListBalanceCodeRes, _ error) {
	var (
		d     = make([]*entity.BalanceCode, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "u_balance_code", Page: req.Page, Size: req.Size, Order: "code desc"}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &setting.ReadListBalanceCodeRes{List: d, Total: total}, nil
}
func (c cBalanceCode) Create(ctx context.Context, req *setting.CreateBalanceCodeReq) (_ *setting.CreateBalanceCodeRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "u_balance_code", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cBalanceCode) Del(ctx context.Context, req *setting.DelBalanceCodeReq) (_ *setting.DelBalanceCodeRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "u_balance_code", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cBalanceCode) Update(ctx context.Context, req *setting.UpdateBalanceCodeReq) (_ *setting.UpdateBalanceCodeRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "u_balance_code", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
