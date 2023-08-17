package syssvc

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xstr"
	"strings"
)

func ListApi(ctx context.Context, req *sys.ListApiReq) (int, []*entity.Api, error) {
	var (
		d     = make([]*entity.Api, 0)
		total int
		page  = req.Page
	)

	db := dao.Api.Ctx(ctx)
	if req.Group != "" {
		db = db.Where("group", req.Group)
	}
	if req.Method != "" {
		db = db.Where("method", strings.ToLower(req.Method))
	}
	if req.Url != "" {
		db = db.WhereLike("url", xstr.Like(req.Url))
	}
	if req.Desc != "" {
		db = db.WhereLike("desc", xstr.Like(req.Desc))
	}
	if err := db.Page(page, req.Size).OrderDesc("id").ScanAndCount(&d, &total, false); err != nil {
		return 0, nil, err
	}
	return total, d, nil
}
