package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	model2 "star_net/app/admin/internal/model"
	"star_net/consts"
	"star_net/db/model/entity"
	"star_net/model"
)

type GetMenuByIdReq struct {
	g.Meta `tags:"/sys/menu" path:"/menu" method:"get" dc:"菜单根据ID查询"`
	Id     uint64 `json:"id"`
}
type GetMenuByIdRes *entity.Menu
type GetMenuByPathReq struct {
	g.Meta `tags:"/sys/menu" path:"/menu/getMenuByPath" method:"get" dc:"菜单获取根据路经"`
	Path   string `json:"path"`
}
type GetMenuByPathRes entity.Menu

func (d *GetMenuByPathRes) AddImgPrefix() {
	if d.Icon != "" {
		d.Icon = consts.ImgPrefix + d.Icon
	}
	if d.BgImg != "" {
		d.Icon = consts.ImgPrefix + d.Icon
	}
}

type ListMenusReq struct {
	g.Meta `tags:"/sys/menu" path:"/menu/listWithChildren" method:"get" dc:"列表查询菜单带子级"`
}
type ListMenusRes []*model2.Menu

type AddMenuReq struct {
	g.Meta `tags:"/sys/menu" path:"/menu"  method:"post" dc:"菜单列表添加"`
	Menu   entity.Menu           `json:"menu"`
	Api    []*entity.MenuApiRule `json:"api"`
}
type AddMenuRes *model.CommonRes
type EditMenuReq struct {
	g.Meta `tags:"/sys/menu" path:"/menu"  method:"put" dc:"菜单修该 "`
	Menu   entity.Menu           `json:"menu"`
	Api    []*entity.MenuApiRule `json:"api"`
}
type EditMenuRes *model.CommonRes
type DelMenuReq struct {
	g.Meta `tags:"/sys/menu" path:"/menu"  method:"delete" dc:"菜单列表删除"`
	Id     int `json:"id"`
}
type DelMenuRes model.CommonRes
