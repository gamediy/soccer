package soccer

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type CreatePlayRes model.CommonRes
type UpdatePlayReq struct {
	g.Meta `tags:"/足球/玩法" method:"put" path:"/play" dc:"修改玩法"`
	*entity.Play
}
type UpdatePlayRes model.CommonRes
type ReadPlayReq struct {
	g.Meta `tags:"/足球/玩法" method:"get" path:"/play" dc:"查询玩法"`
	Id     uint64 `json:"id"`
}
type ReadPlayRes entity.Play
type ReadListPlayReq struct {
	g.Meta `tags:"/足球/玩法" method:"get" path:"/play/list" dc:"查询玩法列表"`
	model.CommonPageReq
	TypeName string
	TypeCode string
}
type ReadListPlayRes struct {
	Total int            `json:"total"`
	List  []*entity.Play `json:"list"`
}
type CreatePlayReq struct {
	g.Meta `tags:"/足球/玩法" method:"post" path:"/play" dc:"添加玩法"`
	*entity.Play
}
type DelPlayReq struct {
	g.Meta `tags:"/足球/玩法" method:"delete" path:"/play" dc:"删除玩法"`
	Id     uint64 `json:"id"`
}
type DelPlayRes model.CommonRes
