package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"star_net/app/api-user/internal/controller/user"
	"star_net/app/api-user/internal/service/usersvc"
	"star_net/core/auth"
	"star_net/utility/utils/xpusher"

	"star_net/common"
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
			// init auth role
			initAuthRule(ctx)
			// init router
			initRouter(s, ctx)
			xpusher.InitFromCfg(ctx)
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
	gfToken.AuthExcludePaths = g.SliceStr{"/api/user/getCaptcha", "/api/user/register", "/api.json", "/api/dict/**"}
	gfToken.LoginBeforeFunc = usersvc.Login
	gfToken.AuthAfterFunc = usersvc.AuthAfterFunc
	auth.GFToken = gfToken
}

/*
统一路由注册
*/
func initRouter(s *ghttp.Server, ctx context.Context) {
	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS, common.MiddlewareRequestLimit, common.MiddlewareHandlerResponse)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Bind(user.User)
		})
	})

	// 启动gtoken
	if err := auth.GFToken.Start(); err != nil {
		panic(err)
	}
}
