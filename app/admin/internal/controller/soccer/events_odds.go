package soccer

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/admin/api/soccer"
	"star_net/app/admin/internal/service/soccersvc"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xcrud"
)

var (
	EventsOdds = cEventsOdds{}
)

type cEventsOdds struct{}

func (c cEventsOdds) BatchInsertReq(ctx context.Context, req *soccer.BatchInsertReq) (_ *model.CommonRes, err error) {
	x := soccersvc.BatchInsert{
		Data: req.Data,
	}
	if err = x.Exec(ctx); err != nil {
		return nil, err
	}
	return
}

func (c cEventsOdds) Create(ctx context.Context, req *soccer.CreateEventsOddsReq) (_ *soccer.CreateEventsOddsRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "p_events_odds", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cEventsOdds) Update(ctx context.Context, req *soccer.UpdateEventsOddsReq) (_ *soccer.UpdateEventsOddsRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_events_odds", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cEventsOdds) Read(ctx context.Context, req *soccer.ReadEventsOddsReq) (_ *soccer.ReadEventsOddsRes, _ error) {
	var d soccer.ReadEventsOddsRes
	x := xcrud.Read{Ctx: ctx, Table: "p_events_odds", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cEventsOdds) ReadList(ctx context.Context, req *soccer.ReadListEventsOddsReq) (_ *soccer.ReadListEventsOddsRes, _ error) {
	var (
		d     = make([]*entity.EventsOdds, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_events_odds", Page: req.Page, Size: req.Size, Where: func(db *gdb.Model) {
		if req.EventsId != "" {
			db = db.Where("events_id", req.EventsId)
		}
		if req.Status != "" {
			db = db.Where("status", req.Status)
		}
		if req.BoutStatus != "" {
			db = db.Where("bout_status", req.BoutStatus)
		}
	}}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListEventsOddsRes{List: d, Total: total}, nil
}
func (c cEventsOdds) Del(ctx context.Context, req *soccer.DelEventsOddsReq) (_ *soccer.DelEventsOddsRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_events_odds", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
