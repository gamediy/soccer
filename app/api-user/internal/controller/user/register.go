package user

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	vpassport "star_net/app/api-user/api/user"
	"star_net/app/api-user/internal/service/usersvc"
	"star_net/model"
)

func (c *cUser) Register(ctx context.Context, req *vpassport.RegisterReq) (res *model.CommonRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	x := usersvc.Register{
		Ctx:         ctx,
		Account:     req.Account,
		Password:    req.Password,
		Xid:         req.Xid,
		Country:     req.Account,
		City:        req.City,
		Ip:          r.GetClientIp(),
		RealName:    req.RealName,
		ClientAgent: r.UserAgent(),
	}
	_, err = x.Exec()
	if err != nil {
		return nil, err
	}
	return
}
