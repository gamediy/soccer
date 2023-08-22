package user

import (
	"context"
	"star_net/app/api-user/api/user"
	"star_net/app/api-user/internal/service"
	"star_net/app/api-user/internal/service/usersvc"
)

func (c cUser) Info(ctx context.Context, _ *user.InfoReq) (_ *user.InfoRes, _ error) {
	info := service.GetUserInfo(ctx)
	x := usersvc.GetInfo{
		Ctx: ctx,
		Uid: info.UidInt64,
	}
	u, w, err := x.Exec()
	if err != nil {
		return nil, err
	}
	return &user.InfoRes{
		User:   u,
		Wallet: w,
	}, nil
}
