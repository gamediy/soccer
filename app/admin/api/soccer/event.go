package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateEventsRes model.CommonRes
type UpdateEventsReq struct {
	g.Meta `tags:"/足球/赛事" method:"put" path:"/events" dc:"修改赛事"`
	*entity.Events
}
type UpdateEventsRes model.CommonRes
type ReadEventsReq struct {
	g.Meta `tags:"/足球/赛事" method:"get" path:"/events" dc:"查询赛事"`
	Id     uint64 `json:"id"`
}
type ReadEventsRes entity.Events
type ReadListEventsReq struct {
	g.Meta `tags:"/足球/赛事" method:"get" path:"/events/list" dc:"查询赛事列表"`
	model.CommonPageReq
}
type ReadListEventsRes struct {
	Total int              `json:"total"`
	List  []*entity.Events `json:"list"`
}
type CreateEventsReq struct {
	g.Meta `tags:"/足球/赛事" method:"post" path:"/events" dc:"添加赛事"`
	*entity.Events
}
type DelEventsReq struct {
	g.Meta `tags:"/足球/赛事" method:"delete" path:"/events" dc:"删除赛事"`
	Id     uint64 `json:"id"`
}
type DelEventsRes model.CommonRes

type OpenResultReq struct {
	g.Meta     `tags:"/足球/赛事" method:"post" path:"/open" dc:"赛事开奖"`
	Result     string `json:"result" dc:"1-1"`
	BoutStatus int    `json:"boutStatus" dc:"1:上，2：下，3：全"`
	EventsId   int64  `json:"eventsId"`
}
