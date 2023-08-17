package get

import (
	"context"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

func AdminById(ctx context.Context, id uint64) (*entity.Admin, error) {
	var d entity.Admin
	one, err := dao.Admin.Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&d); err != nil {
		return nil, err
	}

	return &d, nil
}

func AdminByUname(ctx context.Context, uname string) (*entity.Admin, error) {
	var d *entity.Admin
	one, err := dao.Admin.Ctx(ctx).One("uname", uname)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&d); err != nil {
		return nil, err
	}
	return d, nil
}
