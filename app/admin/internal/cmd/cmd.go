package cmd

import (
	"context"
	"fmt"
	"star_net/app/admin/internal/controller/report"
	"star_net/app/admin/internal/controller/setting"
	"star_net/app/admin/internal/controller/share"
	"star_net/app/admin/internal/controller/sys"
	"star_net/app/admin/internal/controller/user"
	"star_net/app/admin/internal/controller/wallet"
	"star_net/app/admin/internal/model"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/common"
	"star_net/utility/utils/xcasbin"
	"star_net/utility/utils/xelastic"
	"strings"
	"time"

	"github.com/goflyfox/gtoken/gtoken"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			init := syssvc.InitD{}
			init.Exec(ctx)
			s := g.Server("star_net")
			initRouter(s)
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
	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS)
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(common.MiddlewareHandlerResponse)
		g.Group("/api", func(g *ghttp.RouterGroup) {
			g.Group("/user", func(g *ghttp.RouterGroup) {
				g.Bind(user.User)
				g.Bind(user.UserLoginLog)
			})
			g.Group("/wallet", func(g *ghttp.RouterGroup) {
				g.Bind(wallet.Wallet)
				g.Bind(wallet.Log)
				g.Bind(wallet.Deposit)
				g.Bind(wallet.Bank)
				g.Bind(wallet.Withdraw)
			})
			g.Group("/setting", func(g *ghttp.RouterGroup) {
				g.Bind(setting.BalanceCode)
				g.Bind(setting.AmountCategory)
				g.Bind(setting.AmountItem)
			})
			g.Group("/sys", func(g *ghttp.RouterGroup) {
				g.Bind(sys.Role)
				g.Bind(sys.Admin)
				g.Bind(sys.Dict)
				g.Bind(sys.Menu)
				g.Bind(sys.Ws)
				g.Bind(sys.Api)
				g.Bind(sys.AdminLoginLog)
				g.Bind(sys.Pusher)
				g.Bind(sys.File)
			})
			g.Group("/share", func(g *ghttp.RouterGroup) {
				g.Bind(share.Banner)
				g.Bind(share.Language)
			})
			g.Group("/report", func(g *ghttp.RouterGroup) {
				g.Bind(report.Report)
				g.Bind(share.Language)
			})
		})
	})
	loginFunc := syssvc.Login
	// 启动gtoken
	gfToken = &gtoken.GfToken{
		ServerName: "star_net",
		//Timeout:         10 * 1000,
		CacheMode: 2, //1gcache 2redis 3file
		CacheKey:  "start_net_token:",
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
				ctx := r.Context()
				timeStart := ctx.Value("time").(time.Time)
				s2 := respData.Data.(map[string]interface{})
				userInfo := s2["data"].(map[string]interface{})
				u := model.UserInfo{
					Uid:      userInfo["uid"].(float64),
					Account:  userInfo["account"].(string),
					RuleName: userInfo["ruleName"].(string),
					RuleId:   gconv.Int(userInfo["ruleId"].(float64)),
					ClientIp: userInfo["clientIp"].(string),
				}
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
		},
		MultiLogin:       true,
		LoginPath:        "/api/user/login",
		LoginBeforeFunc:  loginFunc,
		LogoutPath:       "/api/user/logout",
		AuthPaths:        g.SliceStr{"/api"},                                                        // 这里是按照前缀拦截，拦截/user /user/list /user/add ...
		AuthExcludePaths: g.SliceStr{"/api/user/login", "/api.json", "/api/sys/menu/getMenuByPath"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		GlobalMiddleware: true,                                                                      // 开启全局拦截
	}
	err := gfToken.Start()
	if err != nil {
		panic(err)
	}
}
