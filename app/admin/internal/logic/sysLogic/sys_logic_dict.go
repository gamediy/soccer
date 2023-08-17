package sysLogic

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

func ListDict(ctx context.Context, page, size int) (int, []*entity.Dict, error) {
	var d = make([]*entity.Dict, 0)
	db := dao.Dict.Ctx(ctx)
	count, err := db.Count()
	if err != nil {
		return 0, nil, err
	}
	if err = db.Page(page, size).OrderDesc("group,id").Scan(&d); err != nil {
		return 0, nil, err
	}

	return count, d, nil
}
