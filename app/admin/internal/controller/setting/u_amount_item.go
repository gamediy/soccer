package setting

import (
	"context"
	"star_net/app/admin/api/setting"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	AmountItem = cAmountItem{}
)

type cAmountItem struct{}

func (c cAmountItem) Read(ctx context.Context, req *setting.ReadAmountItemReq) (_ *setting.ReadAmountItemRes, _ error) {
	var d setting.ReadAmountItemRes
	x := xcrud.Read{Ctx: ctx, Table: "u_amount_item", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cAmountItem) ReadList(ctx context.Context, req *setting.ReadListAmountItemReq) (_ *setting.ReadListAmountItemRes, _ error) {
	var (
		d     = make([]*entity.AmountItem, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "u_amount_item", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &setting.ReadListAmountItemRes{List: d, Total: total}, nil
}
func (c cAmountItem) Create(ctx context.Context, req *setting.CreateAmountItemReq) (_ *setting.CreateAmountItemRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "u_amount_item", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cAmountItem) Del(ctx context.Context, req *setting.DelAmountItemReq) (_ *setting.DelAmountItemRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "u_amount_item", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cAmountItem) Update(ctx context.Context, req *setting.UpdateAmountItemReq) (_ *setting.UpdateAmountItemRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "u_amount_item", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
