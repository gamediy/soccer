package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"star_net/app/api-soccer/internal/controller/hello"
	"star_net/common"
	"star_net/consts"
	"star_net/model"
	"star_net/utility/utils/xpusher"
	"time"
)

var (
	serverName = "star_net_user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server(serverName)
			initRouter(s)
			xpusher.InitFromCfg(ctx)
			s.Run()
			return nil
		},
	}
)
var gfToken *gtoken.GfToken

/*
统一路由注册
*/
func initRouter(s *ghttp.Server) {
	gfToken = &gtoken.GfToken{
		ServerName: serverName,
		//Timeout:         10 * 1000,
		CacheMode: 2, //1gcache 2redis 3file
		CacheKey:  "start_net_user_token:",
		AuthBeforeFunc: func(r *ghttp.Request) bool {
			r.SetCtxVar("time", time.Now())
			return true
		},

		TokenDelimiter: "_",
		EncryptKey:     []byte("koi29a83idakguqjq29asd9asd8a7jhq"),
		AuthFailMsg:    "登录超时，请重新登录",
		AuthAfterFunc: func(r *ghttp.Request, respData gtoken.Resp) {
			if respData.Code != 0 {
				switch r.Method {
				case "PUT", "DELETE", "GET", "POST":
					r.Response.WriteJsonExit(respData)
				}
			} else {
				s2 := respData.Data.(map[string]interface{})

				userInfo := s2["data"].(map[string]interface{})
				u := model.UserInfo{
					Uid:      userInfo["uid"].(float64),
					Account:  userInfo["account"].(string),
					Lang:     r.Request.Header.Get("lang"),
					ClientIP: r.GetClientIp(),
				}
				u.UidInt64 = int64(u.Uid)
				i18n := gi18n.New()
				i18n.SetLanguage(u.Lang)
				u.I18n = i18n
				r.SetCtxVar(consts.UserInfo, u)
			}
		},
		MultiLogin: true,
		LoginPath:  "/api/user/login",
		LoginBeforeFunc: func(r *ghttp.Request) (string, interface{}) {
			return "", nil
		},
		LogoutPath: "/api/user/logout",
	}
	s.Group("/api/soccer", func(group *ghttp.RouterGroup) {
		group.Middleware(common.MiddlewareDefaultCORS, common.MiddlewareHandlerResponse, common.MiddlewareRequestLimit)
		group.Group("/hello", func(group *ghttp.RouterGroup) {
			//gfToken.Middleware(context.Background(), group)
			group.Bind(hello.Ctrl)
		})

		group.Group("/order", func(group *ghttp.RouterGroup) {
			//gfToken.Middleware(context.Background(), group)

		})

	})

	// 启动gtoken

}
