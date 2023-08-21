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
	gfToken := AGfToken{}
	gfToken.ServerName = g.Cfg().MustGet(ctx, "gfToken.name").String()
	gfToken.CacheMode = g.Cfg().MustGet(ctx, "gfToken.cacheMode").Int8()
	gfToken.CacheKey = g.Cfg().MustGet(ctx, "gfToken.cacheMode").String()
	gfToken.EncryptKey = g.Cfg().MustGet(ctx, "gfToken.encryptKey").Bytes()
	gfToken.AuthFailMsg = g.Cfg().MustGet(ctx, "gfToken.authFailMsg").String()
	gfToken.MultiLogin = g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool()
	gfToken.GlobalMiddleware = g.Cfg().MustGet(ctx, "gfToken.globalMiddleware").Bool()
	return &gfToken
}
