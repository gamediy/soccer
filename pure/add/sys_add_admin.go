package add

import (
	"context"
	"errors"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcasbin"
	"star_net/utility/utils/xpwd"
)

func Admin(ctx context.Context, d *entity.Admin) error {

	d.Pwd = xpwd.GenPwd(d.Pwd)
	if _, err := dao.Admin.Ctx(ctx).Insert(d); err != nil {
		return err
	}
	one, err := dao.Role.Ctx(ctx).Where("id", d.Rid).One()
	if err != nil || one == nil {
		return errors.New("do not have this permission")
	}
	_, err = xcasbin.Cb.AddGroupingPolicy(d.Uname, one.Map()["name"])
	if err != nil {
		return err
	}
	return nil
}
