package xcrud

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type Create struct {
	Db     gdb.DB
	Ctx    context.Context
	Table  string
	Data   interface{}
	Before func(context.Context, *gdb.Model) error
}

func (x *Create) Exec() error {
	if x.Db == nil {
		x.Db = g.DB()
	}
	db := x.Db.Model(x.Table).Ctx(x.Ctx)
	if x.Before != nil {
		if err := x.Before(x.Ctx, db); err != nil {
			return err
		}
	}
	if _, err := db.Insert(x.Data); err != nil {
		return err
	}
	return nil
}
