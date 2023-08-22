package usersvc

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type InertLoginLog struct {
	Ctx     context.Context
	Account string
	Uid     uint64
	Ip      string
}

func (m *InertLoginLog) Exec() error {
	d := entity.UserLoginLog{}
	d.Account = m.Account
	d.Uid = m.Uid
	d.Ip = m.Ip
	_, err := dao.UserLoginLog.Ctx(m.Ctx).Insert(d)
	if err != nil {
		return err
	}
	return nil
}
