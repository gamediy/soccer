package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreateTeamRes model.CommonRes
type UpdateTeamReq struct {
	g.Meta `tags:"/足球/队伍" method:"put" path:"/team" dc:"修改队伍"`
	*entity.Team
}
type UpdateTeamRes model.CommonRes
type ReadTeamReq struct {
	g.Meta `tags:"/足球/队伍" method:"get" path:"/team" dc:"查询队伍"`
	Id     uint64 `json:"id"`
}
type ReadTeamRes entity.Team
type ReadListTeamReq struct {
	g.Meta `tags:"/足球/队伍" method:"get" path:"/team/list" dc:"查询队伍列表"`
	model.CommonPageReq
}
type ReadListTeamRes struct {
	Total int            `json:"total"`
	List  []*entity.Team `json:"list"`
}
type CreateTeamReq struct {
	g.Meta `tags:"/足球/队伍" method:"post" path:"/team" dc:"添加队伍"`
	*entity.Team
}
type DelTeamReq struct {
	g.Meta `tags:"/足球/队伍" method:"delete" path:"/team" dc:"删除队伍"`
	Id     uint64 `json:"id"`
}
type DelTeamRes model.CommonRes
