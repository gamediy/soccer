package xcasbin

import (
	"context"
	casbin "github.com/dobyte/gf-casbin"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Cb *casbin.Enforcer
)

func InitFromCfg(ctx context.Context) {
	get, err := g.Cfg().Get(ctx, "database.default.link")
	if err != nil {
		panic(err)
	}
	Cb, err = casbin.NewEnforcer(&casbin.Options{
		Model:    "./manifest/config/casbin.conf",
		Debug:    false,
		Enable:   true,
		AutoLoad: true,
		Table:    "s_casbin",
		Link:     get.String(),
	})
	if err != nil {
		panic(err)
	}
}
