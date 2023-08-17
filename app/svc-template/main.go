package main

import (
	_ "temple_of_red_earth_golang/app/svc-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"temple_of_red_earth_golang/app/svc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
