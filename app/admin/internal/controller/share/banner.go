package share

import (
	"context"
	"star_net/app/admin/api/share"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xcrud"
)

var (
	Banner = cBanner{}
)

type cBanner struct{}

func (c cBanner) Read(ctx context.Context, req *vshare.GetBannerReq) (_ *vshare.GetBannerRes, _ error) {
	var d vshare.GetBannerRes
	x := xcrud.Read{Ctx: ctx, Table: "c_banner", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cBanner) ReadList(ctx context.Context, req *vshare.ListBannerReq) (_ *vshare.ListBannerRes, _ error) {
	var (
		d     = make([]*entity.Banner, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "c_banner", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &vshare.ListBannerRes{List: d, Total: total}, nil
}
func (c cBanner) Create(ctx context.Context, req *vshare.AddBannerReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "c_banner", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cBanner) Del(ctx context.Context, req *vshare.DelBannerReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "c_banner", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cBanner) Update(ctx context.Context, req *vshare.UpdateBannerReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "c_banner", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
