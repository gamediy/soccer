package setting

import (
	"context"
	"star_net/app/admin/api/setting"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	AmountCategory = cAmountCategory{}
)

type cAmountCategory struct{}

func (c cAmountCategory) Read(ctx context.Context, req *setting.ReadAmountCategoryReq) (_ *setting.ReadAmountCategoryRes, _ error) {
	var d setting.ReadAmountCategoryRes
	x := xcrud.Read{Ctx: ctx, Table: "u_amount_category", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cAmountCategory) ReadList(ctx context.Context, req *setting.ReadListAmountCategoryReq) (_ *setting.ReadListAmountCategoryRes, _ error) {
	var (
		d     = make([]*entity.AmountCategory, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "u_amount_category", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &setting.ReadListAmountCategoryRes{List: d, Total: total}, nil
}
func (c cAmountCategory) Create(ctx context.Context, req *setting.CreateAmountCategoryReq) (_ *setting.CreateAmountCategoryRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "u_amount_category", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cAmountCategory) Del(ctx context.Context, req *setting.DelAmountCategoryReq) (_ *setting.DelAmountCategoryRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "u_amount_category", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cAmountCategory) Update(ctx context.Context, req *setting.UpdateAmountCategoryReq) (_ *setting.UpdateAmountCategoryRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "u_amount_category", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
