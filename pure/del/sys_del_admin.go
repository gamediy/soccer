package del

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcasbin"
)

func Admin(ctx context.Context, id uint64) error {
	admin := entity.Admin{}
	dao.Admin.Ctx(ctx).Scan(&admin, id)
	if _, err := dao.Admin.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}

	_, err := xcasbin.Cb.DeleteUser(admin.Uname)
	if err != nil {
		return err
	}
	return nil
}
