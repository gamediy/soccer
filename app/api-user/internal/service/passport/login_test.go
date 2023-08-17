package spassport

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	vpassport "star_net/app/api-user/api/passport"
	"testing"
)

func TestLogin(t *testing.T) {
	var (
		ctx = gctx.New()
	)
	var req = &vpassport.LoginReq{
		CaptchaId:   "234324",
		CaptchaCode: "2324",
		Account:     "sdfs4234",
		Password:    "sfsdf234234",
	}
	loginRes, err := Service.Login(ctx, req)
	if err != nil {
		g.Dump(err)
		return
	}
	g.Dump(loginRes)
}
