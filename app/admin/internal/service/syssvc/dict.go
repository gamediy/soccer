package syssvc

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/logic/sysLogic"
	"star_net/db/model/entity"
)

func ListDict(ctx context.Context, req *sys.ListDictReq) (int, []*entity.Dict, error) {
	return sysLogic.ListDict(ctx, req.Page, req.Size)
}
