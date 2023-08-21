package auth

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	GFToken *AGfToken
)

type AGfToken struct {
	gtoken.GfToken
}

func NewGFTokenFromCtx(ctx context.Context) *AGfToken {
	x := AGfToken{}
	x.ServerName = g.Cfg().MustGet(ctx, "gfToken.name").String()
	x.CacheMode = g.Cfg().MustGet(ctx, "gfToken.cacheMode").Int8()
	x.CacheKey = g.Cfg().MustGet(ctx, "gfToken.cacheMode").String()
	x.EncryptKey = g.Cfg().MustGet(ctx, "gfToken.encryptKey").Bytes()
	x.AuthFailMsg = g.Cfg().MustGet(ctx, "gfToken.authFailMsg").String()
	x.MultiLogin = g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool()
	x.GlobalMiddleware = g.Cfg().MustGet(ctx, "gfToken.globalMiddleware").Bool()
	x.LoginPath = "/api/user/login"
	x.LogoutPath = "/api/user/logout"
	return &x
}
