package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"star_net/app/api-user/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}