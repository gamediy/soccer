package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
	g.Meta     `tags:"/足球/赛事" method:"post" path:"/events" dc:"添加赛事"`
	HomeTeamId int         `json:"homeTeamId"       description:"主队Id" v:"required:主队ID不能为空"`
	AwayTeamId int         `json:"awayTeamId"       description:"客队Id" v:"required:客队ID不能为空"`
	StartTime  *gtime.Time `json:"startTime"        description:"开始时间" v:"required#开始时间不能为空"`
	EndTime    *gtime.Time `json:"endTime"          description:"结束时间" v:"required#结束时间不能为空"`
	LeagueId   int64       `json:"leagueId"         description:"联盟编号" v:"required#联盟"`
	IsHot      int         `json:"isHot"            description:"热门"`
	Handicap   int         `json:"handicap"         description:"让球"`
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
