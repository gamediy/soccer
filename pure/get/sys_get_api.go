package get

import (
	"context"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

func ApiByGroupAndUrl(ctx context.Context, group, url string, method string) (*entity.Api, error) {
	var d entity.Api
	one, err := dao.Api.Ctx(ctx).One("`group` = ? and url = ? and method=?", group, url, method)
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
