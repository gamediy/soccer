package sysLogic

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

func ListAdmin(ctx context.Context, page, size int) (int, []*entity.Admin, error) {
	var d = make([]*entity.Admin, 0)
	db := dao.Admin.Ctx(ctx)
	count, err := db.Count()
	if err != nil {
		return 0, nil, err
	}
	if err = db.Page(page, size).Order("id desc").Scan(&d); err != nil {
		return 0, nil, err
	}
	return count, d, nil
}
