package syssvc

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/internal/model"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

var (
	Menu = sMenu{}
)

type sMenu struct {
}

func (s *sMenu) List(ctx context.Context) ([]*model.Menu, error) {

	var list []*model.Menu
	err := dao.Menu.Ctx(ctx).Scan(&list)
	all, _ := g.Model(" s_menu_api_rule m").LeftJoin("s_api a", "a.id=m.api_id").Fields("a.id,a.url,a.desc,a.method,m.menu_id").All()
	if err != nil {
		return list, err
	}
	//for _, i := range list {
	//	i.AddImgPrefix()
	//}
	menus := s.child(ctx, -1, list, all.List())
	return menus, nil
}
func (s *sMenu) Del(ctx context.Context, menuId int) error {
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Ctx(ctx).Delete("s_menu_api_rule", "menu_id", menuId)
		if err != nil {
			return err
		}
		_, err = tx.Ctx(ctx).Delete("s_menu", "id", menuId)
		return err
	})
	return err
}
func (s *sMenu) Add(ctx context.Context, menu entity.Menu, api []*entity.MenuApiRule) error {

	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		insert, err := tx.Ctx(ctx).Insert("s_menu", menu)
		if err != nil {
			return err
		}
		id, err := insert.LastInsertId()
		if err != nil {
			return err
		}
		if api != nil {
			for index, _ := range api {
				api[index].MenuId = id

			}
			_, err = tx.Ctx(ctx).Insert("s_menu_api_rule", api)
		}

		return err

	})
	return err
}
func (s *sMenu) Edit(ctx context.Context, menu entity.Menu, api []*entity.MenuApiRule) error {
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Update("s_menu", menu, "id", menu.Id)
		if err != nil {
			return err
		}
		_, err = tx.Delete("s_menu_api_rule", "menu_id", menu.Id)
		if err != nil {
			return err
		}
		if api != nil {
			for index, _ := range api {
				api[index].MenuId = int64(menu.Id)

			}
			_, err = tx.Insert("s_menu_api_rule", api)
		}
		return err
	})
	return err
}

func (s *sMenu) ListMenuByRoleId(ctx context.Context, roleId int) ([]*model.Menu, error) {
	result, err := dao.RoleMenu.Ctx(ctx).Where("role_id", roleId).All()
	if err != nil {
		return nil, err
	}
	menuId := make([]int32, 0)
	for _, record := range result {
		menuId = append(menuId, record.Map()["menu_id"].(int32))
	}
	var list []*model.Menu
	err = dao.Menu.Ctx(ctx).Where("status=? and id in(?) and type<?", 1, menuId, 3).Scan(&list)
	all, _ := g.Model(" s_menu_api_rule m").LeftJoin("s_api a", "a.id=m.api_id").Fields("a.id,a.url,a.desc,a.method,m.menu_id").All()
	if err != nil {
		return list, err
	}
	menus := s.child(ctx, -1, list, all.List())
	return menus, nil
}

func (s *sMenu) ListButtonByRoleId(ctx context.Context, roleId int) ([]*model.Menu, error) {
	result, err := dao.RoleMenu.Ctx(ctx).Where("role_id", roleId).All()
	if err != nil {
		return nil, err
	}
	menuId := make([]int32, 0)
	for _, record := range result {
		menuId = append(menuId, record.Map()["menu_id"].(int32))
	}
	var list []*model.Menu
	err = dao.Menu.Ctx(ctx).Where("status=? and id in(?) and type=?", 1, menuId, 3).Scan(&list)
	return list, err
}

func (s *sMenu) child(ctx context.Context, pid int, menus []*model.Menu, apiList gdb.List) []*model.Menu {
	var result []*model.Menu
	for _, menu := range menus {
		if menu.Pid == pid {
			menu.Children = s.child(context.Background(), menu.Id, menus, apiList)
			for _, api := range apiList {
				i := api["menu_id"].(int64)
				if i == int64(menu.Id) && api["id"] != nil {
					api := entity.Api{
						Url:    api["url"].(string),
						Id:     uint64(api["id"].(int64)),
						Desc:   api["desc"].(string),
						Method: api["method"].(string),
					}
					menu.Api = append(menu.Api, &api)
				}

			}
			result = append(result, menu)
		}
	}
	return result

}
