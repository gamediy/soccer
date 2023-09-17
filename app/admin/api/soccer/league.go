package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateLeagueRes model.CommonRes
type CreateLeagueReq struct {
	g.Meta `tags:"/soccer/league" method:"post" path:"/league" dc:"add league"`
	*entity.League
}
type ReadLeagueRes entity.League
type ReadListLeagueReq struct {
	g.Meta `tags:"/soccer/league" method:"get" path:"/league/list" dc:"query league list"`
	model.CommonPageReq
}
type ReadListLeagueRes struct {
	Total int              `json:"total"`
	List  []*entity.League `json:"list"`
}
type UpdateLeagueReq struct {
	g.Meta `tags:"/soccer/league" method:"put" path:"/league" dc:"update league"`
	*entity.League
}
type UpdateLeagueRes model.CommonRes
type ReadLeagueReq struct {
	g.Meta `tags:"/soccer/league" method:"get" path:"/league" dc:"query league"`
	Id     uint64 `json:"id"`
}
type DelLeagueReq struct {
	g.Meta `tags:"/soccer/league" method:"delete" path:"/league" dc:"del league"`
	Id     uint64 `json:"id"`
}
type DelLeagueRes model.CommonRes
