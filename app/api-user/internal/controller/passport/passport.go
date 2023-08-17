package passport

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/app/api-user/api/passport"
	"star_net/app/api-user/consts"
	"star_net/app/api-user/internal/service/passport"
)

var (
	Ctrl = &cPassport{}
)

type cPassport struct {
}

func (c *cPassport) Register(ctx context.Context, req *vpassport.RegisterReq) (res *vpassport.RegisterRes, err error) {
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
	r := ghttp.RequestFromCtx(ctx)
	req.Ip = r.GetClientIp()
	req.UserAgent = r.UserAgent()
	result, err := spassport.Service.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *cPassport) Login(ctx context.Context, req *vpassport.LoginReq) (res *vpassport.LoginRes, err error) {
	if err = g.Validator().Rules("required").Data(req.Account).Run(ctx); err != nil {
		return nil, consts.ErrAreaCode
	}
	if err = g.Validator().Rules("required|password").Data(req.Password).Run(ctx); err != nil {
		return nil, consts.ErrPassFormat
	}
	result, err := spassport.Service.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *cPassport) GetCaptcha(ctx context.Context, req *vpassport.CaptchaReq) (res *vpassport.CaptchaRes, err error) {
	return spassport.Service.GetCaptcha(ctx, req)
}
