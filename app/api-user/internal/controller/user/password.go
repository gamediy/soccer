package user

import (
	"context"
	"star_net/app/api-user/api/user"
	"star_net/app/api-user/internal/service"
	"star_net/app/api-user/internal/service/usersvc"
	"star_net/model"
)

func (c cUser) PutLoginPass(ctx context.Context, req *user.PutLoginPassReq) (_ *model.CommonRes, _ error) {
	x := usersvc.PutLoginPass{
		Ctx:     ctx,
		OldPass: req.OldPass,
		NewPass: req.NewPass,
		Uid:     uint64(service.GetUserInfo(ctx).UidInt64),
	}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
