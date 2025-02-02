package soccer

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/admin/api/soccer"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xstr"
)

var (
	SoccerOrderSettle = cSoccerOrderSettle{}
)

type cSoccerOrderSettle struct{}

func (c cSoccerOrderSettle) Create(ctx context.Context, req *soccer.CreateSoccerOrderSettleReq) (_ *soccer.CreateSoccerOrderSettleRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "o_soccer_order_settle", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cSoccerOrderSettle) Update(ctx context.Context, req *soccer.UpdateSoccerOrderSettleReq) (_ *soccer.UpdateSoccerOrderSettleRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "o_soccer_order_settle", Field: "order_no", V: req.OrderNo, Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cSoccerOrderSettle) Read(ctx context.Context, req *soccer.ReadSoccerOrderSettleReq) (_ *soccer.ReadSoccerOrderSettleRes, _ error) {
	var d soccer.ReadSoccerOrderSettleRes
	x := xcrud.Read{Ctx: ctx, Table: "o_soccer_order_settle", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cSoccerOrderSettle) ReadList(ctx context.Context, req *soccer.ReadListSoccerOrderSettleReq) (_ *soccer.ReadListSoccerOrderSettleRes, _ error) {
	var (
		d     = make([]*entity.SoccerOrderSettle, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "o_soccer_order_settle", Page: req.Page, Size: req.Size, Order: "order_no desc", Where: func(db *gdb.Model) {
		if req.OrderNo != "" {
			db = db.Where("order_no", req.OrderNo)
		}
		if req.Uid != "" {
			db = db.Where("uid", req.Uid)
		}
		if req.Account != "" {
			db = db.WhereLike("account", xstr.Like(req.Account))
		}
		if req.LeagueTitle != "" {
			db = db.WhereLike("league_title", xstr.Like(req.LeagueTitle))
		}
		if req.EventsTitle != "" {
			db = db.WhereLike("events_title", xstr.Like(req.EventsTitle))
		}
		if req.OddsTitle != "" {
			db = db.WhereLike("odds_title", xstr.Like(req.OddsTitle))
		}
		if req.Status != "" {
			db = db.Where("status", req.Status)
		}
	}}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListSoccerOrderSettleRes{List: d, Total: total}, nil
}
func (c cSoccerOrderSettle) Del(ctx context.Context, req *soccer.DelSoccerOrderSettleReq) (_ *soccer.DelSoccerOrderSettleRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "o_soccer_order_settle", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
