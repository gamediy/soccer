package del

import (
	"context"
	"star_net/db/dao"
)

func Dict(ctx context.Context, id uint64) error {
	if _, err := dao.Dict.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
