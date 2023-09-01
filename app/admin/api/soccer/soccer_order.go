package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateSoccerOrderRes model.CommonRes
type UpdateSoccerOrderReq struct {
	g.Meta `tags:"/足球/已结算订单" method:"put" path:"/soccerOrder" dc:"修改已结算订单"`
	*entity.SoccerOrder
}
type UpdateSoccerOrderRes model.CommonRes
type ReadSoccerOrderReq struct {
	g.Meta  `tags:"/足球/已结算订单" method:"get" path:"/soccerOrder" dc:"查询已结算订单"`
	OrderNo uint64 `json:"orderNo"`
}
type ReadSoccerOrderRes entity.SoccerOrder
type ReadListSoccerOrderReq struct {
	g.Meta `tags:"/足球/已结算订单" method:"get" path:"/soccerOrder/list" dc:"查询已结算订单列表"`
	model.CommonPageReq

	OrderNo     string
	Uid         string
	Account     string
	LeagueTitle string
	EventsTitle string
	OddsTitle   string
	Status      string
}
type ReadListSoccerOrderRes struct {
	Total int                   `json:"total"`
	List  []*entity.SoccerOrder `json:"list"`
}
type CreateSoccerOrderReq struct {
	g.Meta `tags:"/足球/已结算订单" method:"post" path:"/soccerOrder" dc:"添加已结算订单"`
	*entity.SoccerOrder
}
type DelSoccerOrderReq struct {
	g.Meta  `tags:"/足球/已结算订单" method:"delete" path:"/soccerOrder" dc:"删除已结算订单"`
	OrderNo uint64 `json:"orderNo"`
}
type DelSoccerOrderRes model.CommonRes
