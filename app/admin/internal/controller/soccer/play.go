package soccer

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/admin/api/soccer"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	Play = cPlay{}
)

type cPlay struct{}

func (c cPlay) Create(ctx context.Context, req *soccer.CreatePlayReq) (_ *soccer.CreatePlayRes, _ error) {
	x := xcrud.Create{Ctx: ctx, Table: "p_play", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cPlay) Update(ctx context.Context, req *soccer.UpdatePlayReq) (_ *soccer.UpdatePlayRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_play", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cPlay) Read(ctx context.Context, req *soccer.ReadPlayReq) (_ *soccer.ReadPlayRes, _ error) {
	var d soccer.ReadPlayRes
	x := xcrud.Read{Ctx: ctx, Table: "p_play", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cPlay) ReadList(ctx context.Context, req *soccer.ReadListPlayReq) (_ *soccer.ReadListPlayRes, _ error) {
	var (
		d     = make([]*entity.Play, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_play", Page: req.Page, Size: req.Size, Where: func(db *gdb.Model) {
		if req.TypeName != "" {
			db = db.Where("type_name", req.TypeName)
		}
		if req.TypeCode != "" {
			db = db.Where("type_code", req.TypeCode)
		}
	}}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListPlayRes{List: d, Total: total}, nil
}
func (c cPlay) Del(ctx context.Context, req *soccer.DelPlayReq) (_ *soccer.DelPlayRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_play", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cPlay) GetPlayAll(ctx gctx.Ctx, req *soccer.GetPlayAllReq) (res []*soccer.GetPlayAllRes, err error) {
	typeList := []entity.PlayType{}
	dao.PlayType.Ctx(ctx).Scan(&typeList)
	res = []*soccer.GetPlayAllRes{}
	for _, playType := range typeList {
		item := soccer.GetPlayAllRes{}
		gconv.Struct(playType, &item)
		dao.Play.Ctx(ctx).Scan(&item.Children, "type_code", playType.Code)
		res = append(res, &item)
	}
	return res, nil
}
