package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"star_net/app/job-soccer/internal/cmd"
	_ "star_net/app/job-soccer/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
