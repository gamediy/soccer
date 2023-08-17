package add

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/api/sys"
	"star_net/db/dao"
)

func Dict(ctx context.Context, d *sys.AddDictReq) error {
	if _, err := dao.Dict.Ctx(ctx).Insert(d); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
