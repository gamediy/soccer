package event

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-soccer/internal/service/eventsvc"
)

type PlayReq struct {
	g.Meta  `tags:"赛事" method:"get" path:"/play" dc:"赛事玩法"`
	EventId int ` dc:"id"  json:"eventId"`
}
type PlayRes struct {
	List []*eventsvc.PlayOddsOutput `json:"list"`
}
