package del

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func DataById(ctx context.Context, table string, id uint64, group ...string) error {
	flagGroup := "default"
	if len(group) > 0 {
		flagGroup = group[0]
	}
	if _, err := g.DB(flagGroup).Model(table).Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
