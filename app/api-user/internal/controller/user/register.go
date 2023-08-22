package user

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	vpassport "star_net/app/api-user/api/user"
	"star_net/app/api-user/internal/service/usersvc"
	"star_net/model"
	"star_net/utility/utils/xformat"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *vpassport.RegisterReq) (res *model.CommonRes, err error) {
	if err = xformat.Account(ctx, req.Account); err != nil {
		return nil, err
	}
	x := usersvc.Register{
		Ctx:      ctx,
		Account:  req.Account,
		Password: req.Password,
		Xid:      req.Xid,
		Country:  req.Account,
		City:     req.City,
		Ip:       ghttp.RequestFromCtx(ctx).GetClientIp(),
		RealName: req.RealName,
	}
	_, err = x.Exec()
	if err != nil {
		return nil, err
	}
	return
}
