package spassport

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	vpassport "star_net/app/api-user/api/passport"
	"testing"
)

func TestGetCaptcha(t *testing.T) {
	var (
		ctx = gctx.New()
	)
	var req = &vpassport.CaptchaReq{
		CaptchaId: "234324",
	}
	captchaRes, err := Service.GetCaptcha(ctx, req)
	if err != nil {
		g.Dump(err)
		return
	}
	g.Dump(captchaRes)
}
