package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"star_net/app/admin/internal/cmd"
	_ "star_net/app/admin/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
