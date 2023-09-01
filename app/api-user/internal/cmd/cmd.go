package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	ctrlComon "star_net/app/api-user/internal/controller/common"
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
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server(g.Cfg().MustGet(ctx, "gfToken.name").String())
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
	gfToken.LoginPath = "/api/user/user/login"
	gfToken.LogoutPath = "/api/user/user/logout"
	gfToken.AuthPaths = g.SliceStr{"/api"}
	gfToken.AuthExcludePaths = g.SliceStr{
		"/api/user/user/getCaptcha",
		"/api/user/user/register",
		"/api/user/user/login",
		"/api.json",
		"/api/user/sys/dict/*",
		"/api/user/doc",
		"/api/user/doc/*",
		"/api/user/api.json",
		"/api/user/common/banners",
	}
	gfToken.LoginBeforeFunc = usersvc.UserLogin
	gfToken.AuthAfterFunc = usersvc.AuthAfterFunc

}

/*
统一路由注册
*/
func initRouter(s *ghttp.Server) {
	s.BindMiddlewareDefault(common.MiddlewareDefaultCORS, common.MiddlewareRequestLimit, common.MiddlewareHandlerResponse)
	s.Group("/api/user", func(group *ghttp.RouterGroup) {

		group.GET("/doc", func(r *ghttp.Request) {
			r.Response.Write(swaggerUIPageContent)
		})
		group.Group("/common", func(group *ghttp.RouterGroup) {
			group.Bind(ctrlComon.Common)
		})
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

	s.SetOpenApiPath("/api/user/doc/api.json")

}

const (
	swaggerUIPageContent = `

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta
    name="description"
    content="SwaggerUI"
  />
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@4.5.0/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@4.5.0/swagger-ui-bundle.js" crossorigin></script>
<script>
  window.onload = () => {
    window.ui = SwaggerUIBundle({
      url: '/api/user/doc/api.json',
      dom_id: '#swagger-ui',
    });
  };
</script>
</body>
</html>



`
)
