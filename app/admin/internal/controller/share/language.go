package share

import (
	"context"
	"star_net/app/admin/api/share"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xcrud"
)

var (
	Language = cLanguage{}
)

type cLanguage struct{}

func (c cLanguage) Read(ctx context.Context, req *vshare.GetLanguageReq) (_ *vshare.GetLanguageRes, _ error) {
	var d vshare.GetLanguageRes
	x := xcrud.Read{Ctx: ctx, Table: "c_language", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cLanguage) ReadList(ctx context.Context, req *vshare.ListLanguageReq) (_ *vshare.ListLanguageRes, _ error) {
	var (
		d     = make([]*entity.Language, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "c_language", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &vshare.ListLanguageRes{List: d, Total: total}, nil
}
func (c cLanguage) Create(ctx context.Context, req *vshare.AddLanguageReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "c_language", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cLanguage) Del(ctx context.Context, req *vshare.DelLanguageReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "c_language", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cLanguage) Update(ctx context.Context, req *vshare.UpdateLanguageReq) (_ *model.CommonRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "c_language", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
