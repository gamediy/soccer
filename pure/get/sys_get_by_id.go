package get

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
)

type ById struct {
	Table string
	Id    uint64
	Group string
	Ctx   context.Context
}

func (my *ById) Exec(pointer interface{}) error {
	one, err := g.DB().Model(my.Table).Where("id", my.Id).One()
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
