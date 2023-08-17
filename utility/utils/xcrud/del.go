package xcrud

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type Del struct {
	DB     gdb.DB
	Ctx    context.Context
	Table  string
	Field  string
	V      interface{}
	Before func(ctx context.Context, model *gdb.Model) error
}

func (x *Del) Exec() error {
	if x.DB == nil {
		x.DB = g.DB()
	}
	if x.Field == "" {
		x.Field = "id"
	}
	if _, err := x.DB.Ctx(x.Ctx).Model(x.Table).Delete(x.Field, x.V); err != nil {
		return err
	}
	return nil
}
