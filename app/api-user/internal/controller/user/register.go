package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	vpassport "star_net/app/api-user/api/user"
	"star_net/app/api-user/consts"
	"star_net/app/api-user/internal/service/usersvc"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *vpassport.RegisterReq) (res *vpassport.RegisterRes, err error) {
	if err = g.Validator().Rules("required").Data(req.Account).Run(ctx); err != nil {
		return nil, consts.ErrAreaCode
	}
	if err = g.Validator().Rules("required").Data(req.Phone).Run(ctx); err != nil {
		return nil, consts.ErrPhoneEmpty
	}
	if err = g.Validator().Rules("min-length:5").Data(req.Phone).Run(ctx); err != nil {
		return nil, consts.ErrPhoneLength5
	}
	if err = g.Validator().Rules("required|password").Data(req.Password).Run(ctx); err != nil {
		return nil, consts.ErrPassFormat
	}
	x := usersvc.Register{
		Ctx:      ctx,
		Account:  req.Account,
		Password: req.Password,
		Xid:      req.Xid,
		Country:  req.Account,
		City:     req.City,
		Ip:       ghttp.RequestFromCtx(ctx).GetClientIp(),
	}
	if err = x.Exec(); err != nil {
		return nil, err
	}
	return
}
