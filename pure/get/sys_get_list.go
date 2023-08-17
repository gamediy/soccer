package get

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type List struct {
	Page   int
	Size   int
	Table  string
	Fields string
	Order  string
	Db     gdb.DB
	Ctx    context.Context
	Where  func(db *gdb.Model)
}

func (my *List) Exec(pointer interface{}, totalCount *int) error {
	my.beforeExec()
	db := my.Db.Model(my.Table).Ctx(my.Ctx)
	if my.Where != nil {
		my.Where(db)
	}
	if err := db.Ctx(my.Ctx).Page(my.Page, my.Size).Order(my.Order).ScanAndCount(pointer, totalCount, false); err != nil {
		return err
	}
	return nil
}

func (my *List) beforeExec() {
	if my.Db == nil {
		my.Db = g.DB()
	}
	if my.Order == "" {
		my.Order = "id desc"
	}
	if my.Page < 0 {
		my.Page = 0
	}
	if my.Size < 0 || my.Size > 10000 {
		my.Size = 10
	}
}
