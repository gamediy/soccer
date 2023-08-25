package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"star_net/app/admin/internal/controller/report"
	"star_net/app/admin/internal/controller/setting"
	"star_net/app/admin/internal/controller/share"
	"star_net/app/admin/internal/controller/soccer"
	"star_net/app/admin/internal/controller/sys"
	"star_net/app/admin/internal/controller/user"
	"star_net/app/admin/internal/controller/wallet"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/common"
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
			initRouter(s, ctx)

			gfToken := syssvc.GFToken{}
			if err := gfToken.New(ctx).Start(); err != nil {
				panic(err)
			}
			s.Run()
			return nil
		},
	}
)

/*
统一路由注册
*/
func initRouter(s *ghttp.Server, ctx context.Context) {
	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS)
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(common.MiddlewareHandlerResponse)
		g.Group("/api", func(g *ghttp.RouterGroup) {
			g.Group("soccer", func(g *ghttp.RouterGroup) {
				g.Bind(soccer.Events)
			})
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
}
