package soccer

import (
	"context"
	"star_net/app/admin/api/soccer"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	League = cLeague{}
)

type cLeague struct{}

func (c cLeague) Create(ctx context.Context, req *soccer.CreateLeagueReq) (_ *soccer.CreateLeagueRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "p_league", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cLeague) Update(ctx context.Context, req *soccer.UpdateLeagueReq) (_ *soccer.UpdateLeagueRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_league", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cLeague) Read(ctx context.Context, req *soccer.ReadLeagueReq) (_ *soccer.ReadLeagueRes, _ error) {
	var d soccer.ReadLeagueRes
	x := xcrud.Read{Ctx: ctx, Table: "p_league", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cLeague) ReadList(ctx context.Context, req *soccer.ReadListLeagueReq) (_ *soccer.ReadListLeagueRes, _ error) {
	var (
		d     = make([]*entity.League, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_league", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListLeagueRes{List: d, Total: total}, nil
}
func (c cLeague) Del(ctx context.Context, req *soccer.DelLeagueReq) (_ *soccer.DelLeagueRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_league", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
