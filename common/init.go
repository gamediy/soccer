package common

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

func init() {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath("../app/admin/manifest/config")
}
