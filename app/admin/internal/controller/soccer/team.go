package soccer

import (
	"context"
	"star_net/app/admin/api/soccer"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	Team = cTeam{}
)

type cTeam struct{}

func (c cTeam) Create(ctx context.Context, req *soccer.CreateTeamReq) (_ *soccer.CreateTeamRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "p_team", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cTeam) Update(ctx context.Context, req *soccer.UpdateTeamReq) (_ *soccer.UpdateTeamRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_team", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cTeam) Read(ctx context.Context, req *soccer.ReadTeamReq) (_ *soccer.ReadTeamRes, _ error) {
	var d soccer.ReadTeamRes
	x := xcrud.Read{Ctx: ctx, Table: "p_team", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cTeam) ReadList(ctx context.Context, req *soccer.ReadListTeamReq) (_ *soccer.ReadListTeamRes, _ error) {
	var (
		d     = make([]*entity.Team, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_team", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListTeamRes{List: d, Total: total}, nil
}
func (c cTeam) Del(ctx context.Context, req *soccer.DelTeamReq) (_ *soccer.DelTeamRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_team", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
