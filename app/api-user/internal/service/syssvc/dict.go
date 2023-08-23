package syssvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type GetDict struct {
	K string
}

func (s GetDict) GetString(ctx context.Context) (string, error) {
	var d entity.Dict
	_ = dao.Dict.Ctx(ctx).Scan(&d, "k", s.K)
	if d.Id == 0 {
		return "", fmt.Errorf("数据不存在")
	}
	return d.V, nil
}
func (s GetDict) GetAll(ctx context.Context) (g.Map, error) {
	all, err := dao.Dict.Ctx(ctx).All("status = 1 and `group` = 2")
	if err != nil {
		return nil, err
	}
	d := g.Map{}
	for _, i := range all {
		d[i["k"].String()] = g.Map{
			"v":    i["v"].String(),
			"desc": i["desc"].String(),
		}
	}
	return d, nil
}
