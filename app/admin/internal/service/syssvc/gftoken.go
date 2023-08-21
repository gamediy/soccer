package syssvc

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcasbin"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xelastic"
	"star_net/utility/utils/xpwd"
	"strings"
	"time"
)

type GFToken struct{}

func (s *GFToken) New(ctx context.Context) *gtoken.GfToken {
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
		LoginBeforeFunc:  s.Login,
		AuthPaths:        g.SliceStr{"/api"},
		AuthExcludePaths: g.SliceStr{"/api/user/login", "/api.json", "/api/sys/admin/otp/check", "/api/sys/admin/otp/validate"},
		AuthBeforeFunc:   s.AuthBeforeFunc,
		AuthAfterFunc:    s.AuthAfterFunc,
	}
	return gfToken
}

func (s *GFToken) Login(r *ghttp.Request) (string, interface{}) {
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
		Account:   uname,
		Uid:       gconv.Float64(admin.Id),
		ClientIp:  r.GetClientIp(),
		RuleId:    admin.Rid,
		RuleName:  role.Name,
		Email:     admin.Email,
		Phone:     admin.Phone,
		LoginTime: time.Now(),
	}
	// 唯一标识，扩展参数user data
	loginLog := LoginLog{
		Uid:     uint64(admin.Id),
		Ip:      r.GetClientIp(),
		Account: uname,
	}
	if err := loginLog.Save(); err != nil {
		r.Response.WriteJson(gtoken.Fail("SAVE LOGIN LOG  ERROR."))
		r.ExitAll()
	}
	return uname, info
}

func (s *GFToken) AuthBeforeFunc(r *ghttp.Request) bool {
	r.SetCtxVar("time", time.Now())
	return true
}
func (s *GFToken) AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Code != 0 {
		switch r.Method {
		case "PUT", "DELETE", "GET", "POST":
			r.Response.WriteJsonExit(respData)
		}
		return
	}
	var (
		ctx       = r.Context()
		timeStart = ctx.Value("time").(time.Time)
		s2        = respData.Data.(map[string]interface{})
		userInfo  = s2["data"].(map[string]interface{})
		u         = model.UserInfo{
			Uid:      userInfo["uid"].(float64),
			Account:  userInfo["account"].(string),
			RuleName: userInfo["ruleName"].(string),
			RuleId:   gconv.Int(userInfo["ruleId"].(float64)),
			ClientIp: userInfo["clientIp"].(string),
		}
	)
	r.SetCtxVar("userName", s2["userKey"])
	r.SetCtxVar("roleName", userInfo["ruleName"])
	r.SetCtxVar("uid", userInfo["uid"])
	r.SetCtxVar("account", userInfo["account"])
	u.AdminId = int64(u.Uid)
	r.SetCtxVar("userInfo", u)

	if !strings.Contains(r.URL.Path, "/api/sys/admin/user_info") && !strings.Contains(r.URL.Path, "/api/sys/menu/getMenuByPath") {
		enforce, _ := xcasbin.Cb.Enforce(u.Account, r.URL.Path, r.Method)
		if !enforce {
			respData.Code = -403
			respData.Data = "未授权"
			r.Response.Status = 403
			r.Response.WriteJsonExit(respData)
			return
		}
	}

	r.Middleware.Next()
	elapsedTime := time.Since(timeStart)
	requestBody := r.GetBodyString()
	//写入操作日志
	go func(ctx1 context.Context) {
		log := model.OperateLog{
			Ip:          u.ClientIp,
			Path:        r.URL.Path,
			Method:      r.Method,
			Account:     u.Account,
			RoleName:    u.RuleName,
			Request:     fmt.Sprint(requestBody),
			Response:    r.Response.BufferString(),
			CreatedAt:   time.Now(),
			ElapsedTime: elapsedTime.Milliseconds(),
		}
		format := time.Now().Format("2006-01-02")
		sprint := fmt.Sprint("admin_operate_log_", format)
		xelastic.Create(ctx1, sprint, log)
	}(context.Background())
}
