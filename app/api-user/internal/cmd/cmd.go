package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"star_net/app/api-user/internal/controller/user"
	"star_net/app/api-user/internal/service/usersvc"
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
			initRouter(s, ctx)
			xpusher.InitFromCfg(ctx)
			s.Run()
			return nil
		},
	}
)

/*
统一路由注册
*/
func initRouter(s *ghttp.Server, ctx context.Context) {
	gfToken := usersvc.NewGFToken(ctx)
	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS, common.MiddlewareRequestLimit, common.MiddlewareHandlerResponse)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Bind(user.User)
		})
	})

	// 启动gtoken
	if err := gfToken.Start(); err != nil {
		panic(err)
	}
}
