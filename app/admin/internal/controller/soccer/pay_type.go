package soccer

import (
	"context"
	"star_net/app/admin/api/soccer"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	PlayType = cPlayType{}
)

type cPlayType struct{}

func (c cPlayType) Create(ctx context.Context, req *soccer.CreatePlayTypeReq) (_ *soccer.CreatePlayTypeRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "p_play_type", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cPlayType) Update(ctx context.Context, req *soccer.UpdatePlayTypeReq) (_ *soccer.UpdatePlayTypeRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_play_type", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cPlayType) Read(ctx context.Context, req *soccer.ReadPlayTypeReq) (_ *soccer.ReadPlayTypeRes, _ error) {
	var d soccer.ReadPlayTypeRes
	x := xcrud.Read{Ctx: ctx, Table: "p_play_type", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cPlayType) ReadList(ctx context.Context, req *soccer.ReadListPlayTypeReq) (_ *soccer.ReadListPlayTypeRes, _ error) {
	var (
		d     = make([]*entity.PlayType, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_play_type", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListPlayTypeRes{List: d, Total: total}, nil
}
func (c cPlayType) Del(ctx context.Context, req *soccer.DelPlayTypeReq) (_ *soccer.DelPlayTypeRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_play_type", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
