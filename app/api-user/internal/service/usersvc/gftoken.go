package usersvc

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xpwd"
	"time"
)

var GFToken *gtoken.GfToken

type sGFToken struct{}

func NewGFToken(ctx context.Context) *gtoken.GfToken {
	ss := sGFToken{}
	gfToken := &gtoken.GfToken{
		ServerName:       g.Cfg().MustGet(ctx, "gfToken.name").String(),
		CacheMode:        g.Cfg().MustGet(ctx, "gfToken.cacheMode").Int8(),
		CacheKey:         g.Cfg().MustGet(ctx, "gfToken.cacheMode").String(),
		EncryptKey:       g.Cfg().MustGet(ctx, "gfToken.encryptKey").Bytes(),
		AuthFailMsg:      g.Cfg().MustGet(ctx, "gfToken.authFailMsg").String(),
		MultiLogin:       g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool(),
		GlobalMiddleware: g.Cfg().MustGet(ctx, "gfToken.globalMiddleware").Bool(),
		LoginPath:        "/api/user/login",
		LogoutPath:       "/api/user/logout",
		LoginBeforeFunc:  ss.Login,
		AuthPaths:        g.SliceStr{"/api"},
		AuthExcludePaths: g.SliceStr{"/api/user/getCaptcha", "/api/user/register", "/api.json", "/api/dict/**"},
		AuthBeforeFunc:   ss.AuthBeforeFunc,
		AuthAfterFunc:    ss.AuthAfterFunc,
	}
	GFToken = gfToken
	return gfToken
}

func (s *sGFToken) Login(r *ghttp.Request) (string, interface{}) {
	var (
		ctx   = r.Context()
		uname = r.Get("uname").String()
		pass  = r.Get("pass").String()
	)
	if uname == "" || pass == "" {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT OR PASSWORD CANNOT BE EMPTY."))
		r.ExitAll()
	}
	// check admin info
	var admin entity.Admin
	x := xcrud.Read{Ctx: ctx, Table: "s_admin", Field: "uname", V: uname}
	if err := x.Exec(&admin); err != nil {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ERROR."))
		r.ExitAll()
	}
	if !xpwd.ComparePassword(admin.Pwd, pass) {
		r.Response.WriteJson(gtoken.Fail("WRONG PASSWORD."))
		r.ExitAll()
	}

	// get role
	var role entity.Role
	x = xcrud.Read{Ctx: ctx, Table: "s_role", V: admin.Rid}
	if err := x.Exec(&role); err != nil {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ROLE ERROR."))
		r.ExitAll()
	}
	info := model.UserInfo{
		Account: uname,
		Uid:     gconv.Float64(admin.Id),
	}
	// 唯一标识，扩展参数user data
	//loginLog := LoginLog{
	//	Uid:     uint64(admin.Id),
	//	Ip:      r.GetClientIp(),
	//	Account: uname,
	//}
	//if err := loginLog.Save(); err != nil {
	//	r.Response.WriteJson(gtoken.Fail("SAVE LOGIN LOG  ERROR."))
	//	r.ExitAll()
	//}
	return uname, info
}

func (s *sGFToken) AuthBeforeFunc(r *ghttp.Request) bool {
	r.SetCtxVar("time", time.Now())
	return true
}
func (s *sGFToken) AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Code != 0 {
		switch r.Method {
		case "PUT", "DELETE", "GET", "POST":
			r.Response.WriteJsonExit(respData)
		}
		return
	}
	var (
		s2       = respData.Data.(map[string]interface{})
		userInfo = s2["data"].(map[string]interface{})
		u        = model.UserInfo{
			Uid:     userInfo["uid"].(float64),
			Account: userInfo["account"].(string),
		}
	)
	r.SetCtxVar("userName", s2["userKey"])
	r.SetCtxVar("roleName", userInfo["ruleName"])
	r.SetCtxVar("uid", userInfo["uid"])
	r.SetCtxVar("account", userInfo["account"])
	r.SetCtxVar("userInfo", u)
}
