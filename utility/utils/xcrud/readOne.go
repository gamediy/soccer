package xcrud

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
)

type Read struct {
	Ctx   context.Context
	DB    gdb.DB
	Table string
	Field string
	V     interface{}
	Group string
}

func (my *Read) Exec(pointer interface{}) error {
	my.beforeExec()
	one, err := my.DB.Model(my.Table).Where(my.Field, my.V).One()
	if err != nil {
		return err
	}
	if one.IsEmpty() {
		return consts.ErrDataNotFound
	}
	if err = one.Struct(pointer); err != nil {
		return err
	}
	return nil
}

func (my *Read) beforeExec() {
	if my.DB == nil {
		my.DB = g.DB()
	}
	if my.Field == "" {
		my.Field = "id"
	}

}
