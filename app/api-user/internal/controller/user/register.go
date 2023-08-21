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
	token, err := x.Exec()
	if err != nil {
		return nil, err
	}
	return &vpassport.RegisterRes{
		Token: token,
	}, nil
}
