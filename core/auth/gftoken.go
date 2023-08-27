package auth

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/consts"
	"star_net/model"
	"star_net/utility/utils/xtrans"
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
	gfToken.AuthAfterFunc = func(r *ghttp.Request, respData gtoken.Resp) {
		if respData.Code != 0 {
			switch r.Method {
			case "PUT", "DELETE", "GET", "POST":
				r.Response.WriteJsonExit(respData)
			}
			return
		}
		var u model.UserInfo
		if err := gjson.New(respData.Data).Get("data").Struct(&u); err != nil {
			r.Response.WriteJsonExit(err)
		}
		u.Lang = r.Request.Header.Get("lang")
		u.ClientIP = r.GetClientIp()
		r.SetCtxVar("account", u.Account)
		r.SetCtxVar("uid", u.UidInt64)
		u.I18n = xtrans.New(u.Lang)
		r.SetCtxVar(consts.UserInfo, u)
		r.Middleware.Next()
	}
	GFToken = &gfToken
	return &gfToken
}
