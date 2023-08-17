package xcrud

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type Update struct {
	Db        gdb.DB
	Ctx       context.Context
	Table     string
	Data      interface{}
	Field     string
	V         interface{}
	OmitEmpty bool
	Before    func(context.Context, *gdb.Model) error
}

func (x *Update) Exec() (err error) {
	if x.Db == nil {
		x.Db = g.DB()
	}
	db := x.Db.Model(x.Table).Ctx(x.Ctx)
	if x.Before != nil {
		if err = x.Before(x.Ctx, db); err != nil {
			return err
		}
	}

	if x.Field == "" {
		x.Field = "id"
	}
	if x.V == nil {
		m := gconv.Map(x.Data)
		x.V = m["id"]
	}

	if x.OmitEmpty {
		if _, err = db.Where(x.Field, x.V).OmitEmpty().Update(x.Data); err != nil {
			return err
		}
	} else {
		if _, err = db.Where(x.Field, x.V).Update(x.Data); err != nil {
			return err
		}
	}
	return nil
}
