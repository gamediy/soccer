package sys

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/admin/api/sys"
	"star_net/consts"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcrud"
)

var (
	Dict = cDict{}
)

type cDict struct{}

func (c *cDict) GetDict(ctx context.Context, req *sys.GetDictReq) (_ *sys.GetDictRes, _ error) {
	var d sys.GetDictRes
	x := xcrud.Read{Ctx: ctx, Table: "s_dict", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c *cDict) List(ctx context.Context, req *sys.ListDictReq) (_ *sys.ListDictRes, _ error) {
	var (
		d     = make([]*entity.Dict, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Page: req.Page, Size: req.Size, Table: "s_dict"}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &sys.ListDictRes{Total: total, List: d}, nil
}
func (c *cDict) Add(ctx context.Context, req *sys.AddDictReq) (res *sys.AddDictRes, err error) {
	x := xcrud.Create{Ctx: ctx, Table: "s_dict", Data: req, Before: func(ctx context.Context, db *gdb.Model) error {
		count, err := db.Count("k", req.K)
		if err != nil {
			return err
		}
		if count != 0 {
			return consts.ErrKeyAlreadyExist
		}
		return nil
	},
	}
	if err = x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c *cDict) Del(ctx context.Context, req *sys.DelDictReq) (_ *sys.DelDictRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "s_dict", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c *cDict) Update(ctx context.Context, req *sys.UpdateDictReq) (_ *sys.UpdateDictRes, _ error) {
	m := gconv.Map(req)
	delete(m, "k")
	x := xcrud.Update{Ctx: ctx, Data: m, Table: "s_dict"}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
