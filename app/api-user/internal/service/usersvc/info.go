package usersvc

import (
	"context"
	"fmt"
	"star_net/app/api-user/internal/model"
	"star_net/db/dao"
)

type GetInfo struct {
	Ctx context.Context
	Uid int64
}

func (m *GetInfo) Exec() (*model.User, *model.Wallet, error) {
	var (
		u model.User
		w model.Wallet
	)
	_ = dao.User.Ctx(m.Ctx).Scan(&u, "id", m.Uid)
	_ = dao.Wallet.Ctx(m.Ctx).Scan(&w, "uid", m.Uid)
	if u.Id == 0 || w.Uid == 0 {
		return nil, nil, fmt.Errorf("用户数据不存在")
	}
	return &u, &w, nil
}
