package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateSoccerOrderSettleRes model.CommonRes
type UpdateSoccerOrderSettleReq struct {
	g.Meta `tags:"/足球/未结算订单" method:"put" path:"/soccerOrderSettle" dc:"修改未结算订单"`
	*entity.SoccerOrderSettle
}
type UpdateSoccerOrderSettleRes model.CommonRes
type ReadSoccerOrderSettleReq struct {
	g.Meta  `tags:"/足球/未结算订单" method:"get" path:"/soccerOrderSettle" dc:"查询未结算订单"`
	OrderNo uint64 `json:"orderNo"`
}
type ReadSoccerOrderSettleRes entity.SoccerOrderSettle
type ReadListSoccerOrderSettleReq struct {
	g.Meta `tags:"/足球/未结算订单" method:"get" path:"/soccerOrderSettle/list" dc:"查询未结算订单列表"`
	model.CommonPageReq
	OrderNo     string
	Uid         string
	Account     string
	LeagueTitle string
	EventsTitle string
	OddsTitle   string
	Status      string
}
type ReadListSoccerOrderSettleRes struct {
	Total int                         `json:"total"`
	List  []*entity.SoccerOrderSettle `json:"list"`
}
type CreateSoccerOrderSettleReq struct {
	g.Meta `tags:"/足球/未结算订单" method:"post" path:"/soccerOrderSettle" dc:"添加未结算订单"`
	*entity.SoccerOrderSettle
}
type DelSoccerOrderSettleReq struct {
	g.Meta  `tags:"/足球/未结算订单" method:"delete" path:"/soccerOrderSettle" dc:"删除未结算订单"`
	OrderNo uint64 `json:"orderNo"`
}
type DelSoccerOrderSettleRes model.CommonRes
