package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/pure/del"
	"star_net/pure/get"
	"star_net/pure/put"
	"star_net/utility/utils/xcrud"
)

var (
	Api = cApi{}
)

type cApi struct{}

func (c cApi) GetApi(ctx context.Context, req *sys.GetApiReq) (_ *sys.GetApiRes, err error) {
	var (
		d sys.GetApiRes
	)
	byId := get.ById{
		Table: "s_api",
		Id:    req.Id,
		Ctx:   ctx,
	}
	if err = byId.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cApi) ListApi(ctx context.Context, req *sys.ListApiReq) (res *sys.ListApiRes, err error) {
	total, apis, err := syssvc.ListApi(ctx, req)
	if err != nil {
		return nil, err
	}
	return &sys.ListApiRes{Total: total, List: apis}, nil
}
func (c cApi) AddApi(ctx context.Context, req *sys.AddApiReq) (res *sys.AddApiRes, err error) {
	x := xcrud.Create{Ctx: ctx, Table: "s_api", Data: req}
	if err = x.Exec(); err != nil {
		return nil, err
	}
	return nil, err
}

func (c cApi) UpdateApi(ctx context.Context, req *sys.UpdateApiReq) (_ *sys.UpdateApiRes, _ error) {
	if err := put.DataByIdOmitEmpty(ctx, "s_api", req.Id, req.Api); err != nil {
		return nil, err
	}
	return
}
func (c cApi) DelApi(ctx context.Context, req *sys.DelApiReq) (_ *sys.DelApiRes, _ error) {
	if err := del.DataById(ctx, "s_api", req.Id); err != nil {
		return
	}
	return
}
