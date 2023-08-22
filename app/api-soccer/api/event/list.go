package event

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EventListReq struct {
	g.Meta `tags:"eventsvc" method:"get" path:"/eventsvc/list" dc:"eventsvc"`
}
type EventListRes struct {
}
