package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreatePlayTypeRes model.CommonRes
type UpdatePlayTypeReq struct {
	g.Meta `tags:"/足球/玩法类型" method:"put" path:"/playType" dc:"修改玩法类型"`
	*entity.PlayType
}
type UpdatePlayTypeRes model.CommonRes
type ReadPlayTypeReq struct {
	g.Meta `tags:"/足球/玩法类型" method:"get" path:"/playType" dc:"查询玩法类型"`
	Id     uint64 `json:"id"`
}
type ReadPlayTypeRes entity.PlayType
type ReadListPlayTypeReq struct {
	g.Meta `tags:"/足球/玩法类型" method:"get" path:"/playType/list" dc:"查询玩法类型列表"`
	model.CommonPageReq
}
type ReadListPlayTypeRes struct {
	Total int                `json:"total"`
	List  []*entity.PlayType `json:"list"`
}
type CreatePlayTypeReq struct {
	g.Meta `tags:"/足球/玩法类型" method:"post" path:"/playType" dc:"添加玩法类型"`
	*entity.PlayType
}
type DelPlayTypeReq struct {
	g.Meta `tags:"/足球/玩法类型" method:"delete" path:"/playType" dc:"删除玩法类型"`
	Id     uint64 `json:"id"`
}
type DelPlayTypeRes model.CommonRes
