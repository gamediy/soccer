package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type BatchInsertReq struct {
	g.Meta `tags:"足球赔率管理" method:"post" path:"/eventsOdds/batch" dc:"批量添加"`
	Data   []*entity.EventsOdds `json:"data"`
}

type CreateEventsOddsRes model.CommonRes
type UpdateEventsOddsReq struct {
	g.Meta `tags:"足球赔率管理" method:"put" path:"/eventsOdds" dc:"修改足球配置"`
	*entity.EventsOdds
}
type UpdateEventsOddsRes model.CommonRes
type ReadEventsOddsReq struct {
	g.Meta `tags:"足球赔率管理" method:"get" path:"/eventsOdds" dc:"查询足球配置"`
	Id     uint64 `json:"id"`
}
type ReadEventsOddsRes entity.EventsOdds
type ReadListEventsOddsReq struct {
	g.Meta `tags:"足球赔率管理" method:"get" path:"/eventsOdds/list" dc:"查询足球配置列表"`
	model.CommonPageReq
	EventsId   string
	Status     string
	BoutStatus string
}
type ReadListEventsOddsRes struct {
	Total int                  `json:"total"`
	List  []*entity.EventsOdds `json:"list"`
}
type CreateEventsOddsReq struct {
	g.Meta `tags:"足球赔率管理" method:"post" path:"/eventsOdds" dc:"添加足球配置"`
	*entity.EventsOdds
}
type DelEventsOddsReq struct {
	g.Meta `tags:"足球赔率管理" method:"delete" path:"/eventsOdds" dc:"删除足球配置"`
	Id     uint64 `json:"id"`
}
type DelEventsOddsRes model.CommonRes
