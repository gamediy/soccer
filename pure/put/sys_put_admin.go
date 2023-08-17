package put

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcasbin"
	"star_net/utility/utils/xpwd"
)

func Admin(ctx context.Context, in *entity.Admin) error {
	m := gconv.Map(in)
	delete(m, "uname")
	s := m["pwd"].(string)
	if len(s) < 20 {
		m["pwd"] = xpwd.GenPwd(s)
	} else {
		delete(m, "pwd")
	}
	if _, err := dao.Admin.Ctx(ctx).Update(m, "id", in.Id); err != nil {
		return err
	}
	one, err := dao.Role.Ctx(ctx).Where("id", in.Rid).One()
	if err != nil || one == nil {
		return errors.New("do not have this permission")
	}
	xcasbin.Cb.DeleteUser(in.Uname)
	_, err = xcasbin.Cb.AddGroupingPolicy(in.Uname, one.Map()["name"])
	if err != nil {
		return err
	}
	return nil
}
