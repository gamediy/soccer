package event

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-soccer/internal/service/eventsvc"
)

type EventListReq struct {
	g.Meta `tags:"赛事" method:"get" path:"/list" dc:"赛事列表"`
	Status int ` dc:"0：未开始 1接受下注，2：停止下注，3 已结算 -1所有"  json:"status"`
}
type EventListRes struct {
	List []eventsvc.EventsListOutput `json:"list"`
}
