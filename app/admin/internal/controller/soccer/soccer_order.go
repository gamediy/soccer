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
	SoccerOrder = cSoccerOrder{}
)

type cSoccerOrder struct{}

func (c cSoccerOrder) Create(ctx context.Context, req *soccer.CreateSoccerOrderReq) (_ *soccer.CreateSoccerOrderRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "o_soccer_order", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cSoccerOrder) Update(ctx context.Context, req *soccer.UpdateSoccerOrderReq) (_ *soccer.UpdateSoccerOrderRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "o_soccer_order", Field: "orderNo", V: req.OrderNo, Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cSoccerOrder) Read(ctx context.Context, req *soccer.ReadSoccerOrderReq) (_ *soccer.ReadSoccerOrderRes, _ error) {
	var d soccer.ReadSoccerOrderRes
	x := xcrud.Read{Ctx: ctx, Table: "o_soccer_order", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cSoccerOrder) ReadList(ctx context.Context, req *soccer.ReadListSoccerOrderReq) (_ *soccer.ReadListSoccerOrderRes, _ error) {
	var (
		d     = make([]*entity.SoccerOrder, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "o_soccer_order", Page: req.Page, Size: req.Size, Order: "order_no desc", Where: func(db *gdb.Model) {

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
	return &soccer.ReadListSoccerOrderRes{List: d, Total: total}, nil
}
func (c cSoccerOrder) Del(ctx context.Context, req *soccer.DelSoccerOrderReq) (_ *soccer.DelSoccerOrderRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "o_soccer_order", Field: "order_no", V: req.OrderNo}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
