package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"star_net/app/api-user/internal/controller/sys"
	"star_net/app/api-user/internal/controller/user"
	"star_net/app/api-user/internal/controller/wallet"
	"star_net/app/api-user/internal/service/syssvc"
	"star_net/app/api-user/internal/service/usersvc"
	"star_net/common"
	"star_net/core/auth"
	"star_net/utility/utils/xpusher"
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
			init := syssvc.InitD{}
			init.Exec(ctx)
			// init auth role
			initAuthRule(ctx)
			// init router
			initRouter(s)
			xpusher.InitFromCfg(ctx)
			// 启动gtoken
			if err := auth.GFToken.Start(); err != nil {
				panic(err)
			}
			s.Run()
			return nil
		},
	}
)

func initAuthRule(ctx context.Context) {
	gfToken := auth.NewGFTokenFromCtx(ctx)
	gfToken.LoginPath = "/api/user/login"
	gfToken.LogoutPath = "/api/user/logout"
	gfToken.AuthPaths = g.SliceStr{"/api"}
	gfToken.AuthExcludePaths = g.SliceStr{
		"/api/user/getCaptcha",
		"/api/user/register",
		"/api/user/login",
		"/api.json",
		"/api/sys/dict/*",
	}
	gfToken.LoginBeforeFunc = usersvc.UserLogin
	gfToken.AuthAfterFunc = usersvc.AuthAfterFunc
	auth.GFToken = gfToken
}

/*
统一路由注册
*/
func initRouter(s *ghttp.Server) {
	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS, common.MiddlewareRequestLimit, common.MiddlewareHandlerResponse)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/sys", func(group *ghttp.RouterGroup) {
			group.Bind(sys.Sys)
		})
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Bind(user.User)
		})
		group.Group("/wallet", func(group *ghttp.RouterGroup) {
			group.Bind(wallet.Wallet)
		})
	})

}
