package syssvc

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/internal/model"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcasbin"
)

var (
	Role = sRole{}
)

type sRole struct {
}

func GetRoleById(ctx context.Context, id uint64) (*entity.Role, []*entity.RoleMenu, error) {
	var d entity.Role
	one, err := dao.Role.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, nil, err
	}
	if one.IsEmpty() {
		return nil, nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&d); err != nil {
		return nil, nil, err
	}
	roleMenu := []*entity.RoleMenu{}
	err = dao.RoleMenu.Ctx(ctx).Where("role_id", id).Scan(&roleMenu)
	if err != nil {
		return nil, nil, err
	}
	return &d, roleMenu, nil
}

func (s *sRole) List(ctx context.Context) ([]*entity.Role, int, error) {
	var (
		d     = make([]*entity.Role, 0)
		total int
	)
	if err := dao.Role.Ctx(ctx).ScanAndCount(&d, &total, false); err != nil {
		return nil, 0, err
	}
	return d, total, nil
}
func (s *sRole) Add(ctx context.Context, role model.Role) (err error) {
	count, err := dao.Role.Ctx(ctx).Where("name", role.Name).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("角色名称以存在")
	}
	role.Status = 1
	cb := make([][]string, 0)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		last, err := tx.Insert("s_role", role)
		if err != nil {
			return err
		}
		id, err := last.LastInsertId()
		if err != nil {
			return err
		}
		menuIds := make([]int, 0)
		for index, m := range role.Menu {
			menuIds = append(menuIds, m.MenuId)
			role.Menu[index].Id = 0
			role.Menu[index].RoleId = int(id)
		}
		_, err = tx.Insert("s_role_menu", role.Menu)
		if err != nil {
			return err
		}
		all, err := tx.Model(" s_menu_api_rule m").LeftJoin("s_api a", "a.id=m.api_id").Fields("a.id,a.url,a.desc,a.method,m.menu_id").Where("m.menu_id IN (?)", menuIds).All()
		if err != nil {
			return err
		}
		for _, m := range all.List() {
			cb = append(cb, []string{role.Name, m["url"].(string), m["method"].(string)})
		}
		_, err = xcasbin.Cb.AddPolicies(cb)
		return err
	})
	return err
}
func (s *sRole) Del(ctx context.Context, roleId int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		one, err := tx.Model("s_role").Where("id", roleId).One()
		if err != nil {
			return err
		}
		_, err = tx.Delete("s_role", "id", roleId)

		if err != nil {
			return err
		}
		_, err = tx.Delete("s_role_menu", "role_id", roleId)
		_, err = xcasbin.Cb.DeleteRole(one.Map()["name"].(string))
		return err
	})
	return err

}

func (s *sRole) Edit(ctx context.Context, role model.Role) (err error) {
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		one, err := tx.Model("s_role").Where("id", role.Id).One()
		if err != nil {
			return err
		}
		if one == nil {
			return fmt.Errorf("数据为空")
		}
		list := []entity.Casbin{}
		// 要删除的策略规则

		tx.Model(dao.Casbin.Table()).Where("ptype='p' and v0=?", one.Map()["name"]).Scan(&list)
		for _, casbin := range list {

			xcasbin.Cb.RemovePolicy(casbin.V0, casbin.V1, casbin.V2)
		}

		if err != nil {
			return err
		}
		_, err = tx.Update("s_role", role, "id", role.Id)
		if err != nil {
			return err
		}
		_, err = tx.Delete("s_role_menu", "role_id", role.Id)
		if err != nil {
			return err
		}
		if role.Menu != nil {
			menuIds := make([]int, 0)
			for index, m := range role.Menu {
				menuIds = append(menuIds, m.MenuId)
				role.Menu[index].RoleId = role.Id
			}
			list := []entity.Api{}
			err := tx.Model(" s_menu_api_rule m").LeftJoin("s_api a", "a.id=m.api_id").Fields("a.id,a.url,a.desc,a.method,m.menu_id").Where("m.menu_id IN (?)", menuIds).Scan(&list)
			if err != nil {
				return err
			}
			cb := make([][]string, 0)
			for _, m := range list {
				if m.Method != "" && m.Url != "" {
					cb = append(cb, []string{role.Name, m.Url, m.Method})
				}

			}

			_, err = xcasbin.Cb.AddPolicies(cb)
			if err != nil {
				return err
			}

			_, err = tx.Insert("s_role_menu", role.Menu)
		}

		return err
	})
	return err
}
