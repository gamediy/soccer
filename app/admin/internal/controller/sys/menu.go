package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/consts"
	"star_net/utility/utils/xcrud"
)

var (
	Menu = cMenu{}
)

type cMenu struct {
}

func (c cMenu) GetById(ctx context.Context, req *sys.GetMenuByIdReq) (_ sys.GetMenuByIdRes, _ error) {
	var (
		d sys.GetMenuByIdRes
	)
	x := xcrud.Read{
		Ctx:   ctx,
		Table: "s_menu",
		V:     req.Id,
	}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	if d.Icon != "" {
		d.Icon = consts.ImgPrefix + d.Icon
	}
	if d.BgImg != "" {
		d.Icon = consts.ImgPrefix + d.Icon
	}
	return d, nil
}

func (c cMenu) GetMenuByPath(ctx context.Context, req *sys.GetMenuByPathReq) (_ *sys.GetMenuByPathRes, _ error) {
	var d sys.GetMenuByPathRes
	x := xcrud.Read{Ctx: ctx, Table: "s_menu", Field: "path", V: req.Path}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	//d.AddImgPrefix()
	return &d, nil
}
func (c cMenu) ListMenuWithChildren(ctx context.Context, _ *sys.ListMenusReq) (_ sys.ListMenusRes, _ error) {
	return syssvc.Menu.List(ctx)
}
func (c *cMenu) Add(ctx context.Context, req *sys.AddMenuReq) (res sys.AddMenuRes, err error) {
	err = syssvc.Menu.Add(ctx, req.Menu, req.Api)
	return nil, err
}

func (c *cMenu) Edit(ctx context.Context, req *sys.EditMenuReq) (res sys.EditMenuRes, err error) {
	err = syssvc.Menu.Edit(ctx, req.Menu, req.Api)
	return nil, err
}

func (c *cMenu) Del(ctx context.Context, req *sys.DelMenuReq) (res *sys.DelMenuRes, err error) {
	err = syssvc.Menu.Del(ctx, req.Id)
	return nil, err
}
