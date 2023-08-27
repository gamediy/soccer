package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"star_net/app/api-soccer/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
