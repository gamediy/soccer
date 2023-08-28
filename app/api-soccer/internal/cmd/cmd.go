package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"star_net/app/api-soccer/internal/controller/event"
	"star_net/app/api-soccer/internal/controller/order"
	"star_net/common"
	"star_net/core/auth"
	"star_net/utility/utils/xpusher"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server(g.Cfg().MustGet(ctx, "gfToken.name").String())
			auth.NewGFTokenFromCtx(ctx)
			initRouter(s)
			xpusher.InitFromCfg(ctx)
			s.SetPort(4102)
			auth.GFToken.AuthPaths = g.SliceStr{"/api"}
			auth.GFToken.AuthExcludePaths = g.SliceStr{
				"/api.json",
				"/api/soccer/doc",
				"/api/soccer/doc/*",
				"/api/soccer/api.json",
			}
			auth.GFToken.Start()
			s.Run()
			return nil
		},
	}
)

const (
	swaggerUIPageContent = `
<!DOCTYPE html>
<html>
  <head>
    <title>Redoc</title>
    <!-- needed for adaptive design -->
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

    <!--
    Redoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='/api/soccer/doc/api.json'></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
  </body>
</html>

`
)

/*
统一路由注册
*/
func initRouter(s *ghttp.Server) {

	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS, common.MiddlewareHandlerResponse, common.MiddlewareRequestLimit)
	s.Group("/api/soccer", func(group *ghttp.RouterGroup) {
		group.GET("/doc", func(r *ghttp.Request) {
			r.Response.Write(swaggerUIPageContent)
		})

		group.Group("/order", func(group *ghttp.RouterGroup) {
			auth.GFToken.Middleware(context.Background(), group)
			group.Bind(order.OrderCtrl)
		})
		group.Group("/event", func(group *ghttp.RouterGroup) {
			auth.GFToken.Middleware(context.Background(), group)
			group.Bind(event.EventCtrl)
		})
	})
	s.SetOpenApiPath("/api/soccer/doc/api.json")
	// 启动gtoken

}
