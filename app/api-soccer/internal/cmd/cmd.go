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
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI"/>
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@latest/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@latest/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api/soccer/doc/api.json',
			dom_id: '#swagger-ui',
			   requestInterceptor: (req) => {
                    req.headers['Authorization'] = 'Bearer your-auth-token';
                    return req;
                },
		});


	};
</script>
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
