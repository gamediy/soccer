package put

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func DataById(ctx context.Context, table string, id uint64, d interface{}, group ...string) error {
	flagGroup := "default"
	if len(group) > 0 {
		flagGroup = group[0]
	}
	if _, err := g.DB(flagGroup).Ctx(ctx).Model(table).Update(d, "id", id); err != nil {
		return err
	}

	return nil
}

func DataByIdOmitEmpty(ctx context.Context, table string, id uint64, d interface{}, group ...string) error {
	flagGroup := "default"
	if len(group) > 0 {
		flagGroup = group[0]
	}
	if _, err := g.DB(flagGroup).Ctx(ctx).Model(table).OmitEmpty().Update(d, "id", id); err != nil {
		return err
	}

	return nil
}
