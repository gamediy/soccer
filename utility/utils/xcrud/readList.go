package xcrud

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type ReadList struct {
	Ctx    context.Context
	Table  string
	Page   int
	Size   int
	Fields string
	Order  string
	Db     gdb.DB
	Where  func(db *gdb.Model)
}

func (x *ReadList) Exec(pointer interface{}, totalCount *int) error {
	x.beforeExec()
	db := x.Db.Model(x.Table).Ctx(x.Ctx)
	if x.Where != nil {
		x.Where(db)
	}
	if err := db.Ctx(x.Ctx).Offset(x.Page*x.Size).Limit(x.Size).Order(x.Order).ScanAndCount(pointer, totalCount, false); err != nil {
		return err
	}
	return nil
}

func (x *ReadList) beforeExec() {
	if x.Db == nil {
		x.Db = g.DB()
	}
	if x.Order == "" {
		x.Order = "id desc"
	}
	if x.Page < 0 {
		x.Page = 0
	}
	if x.Size < 0 || x.Size > 10000 {
		x.Size = 10
	}
}
