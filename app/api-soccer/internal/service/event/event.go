package event

import "github.com/gogf/gf/v2/os/gctx"

var (
	Service = &event{}
)

type event struct {
}

type Event interface {
	GetEventList(ctx gctx.Ctx)
}
