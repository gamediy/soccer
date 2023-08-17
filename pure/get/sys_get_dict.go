package get

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

func DictById(ctx context.Context, id uint64) (*entity.Dict, error) {
	var d *entity.Dict
	one, err := dao.Dict.Ctx(ctx).One("id", id)
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

func DictByKey(ctx context.Context, k string) (*entity.Dict, error) {
	var data entity.Dict
	one, err := dao.Dict.Ctx(ctx).One("k", k)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
