package soccer

import (
	"context"
	"star_net/app/admin/api/soccer"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	Events = cEvents{}
)

type cEvents struct{}

func (c cEvents) Create(ctx context.Context, req *soccer.CreateEventsReq) (_ *soccer.CreateEventsRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "p_events", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cEvents) Update(ctx context.Context, req *soccer.UpdateEventsReq) (_ *soccer.UpdateEventsRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_events", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cEvents) Read(ctx context.Context, req *soccer.ReadEventsReq) (_ *soccer.ReadEventsRes, _ error) {
	var d soccer.ReadEventsRes
	x := xcrud.Read{Ctx: ctx, Table: "p_events", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cEvents) ReadList(ctx context.Context, req *soccer.ReadListEventsReq) (_ *soccer.ReadListEventsRes, _ error) {
	var (
		d     = make([]*entity.Events, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_events", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListEventsRes{List: d, Total: total}, nil
}
func (c cEvents) Del(ctx context.Context, req *soccer.DelEventsReq) (_ *soccer.DelEventsRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_events", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
