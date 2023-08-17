package setting

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type ReadBalanceCodeReq struct {
	g.Meta `tags:"/sys/balanceCode" method:"get" path:"/balanceCode" dc:"查询账变类型"`
	Id     uint64 `json:"id"`
}
type ReadBalanceCodeRes entity.BalanceCode
type ReadListBalanceCodeReq struct {
	g.Meta `tags:"/sys/balanceCode" method:"get" path:"/balanceCode/list" dc:"查询账变类型列表"`
	model.CommonPageReq
	Account     string
	OrderNo     string
	BalanceCode string
	Node        string
}
type ReadListBalanceCodeRes struct {
	Total int                   `json:"total"`
	List  []*entity.BalanceCode `json:"list"`
}
type CreateBalanceCodeReq struct {
	g.Meta `tags:"/sys/balanceCode" method:"post" path:"/balanceCode" dc:"添加账变类型"`
	*entity.BalanceCode
}
type CreateBalanceCodeRes model.CommonRes
type UpdateBalanceCodeReq struct {
	g.Meta `tags:"/sys/balanceCode" method:"put" path:"/balanceCode" dc:"修改账变类型"`
	*entity.BalanceCode
}
type UpdateBalanceCodeRes model.CommonRes
type DelBalanceCodeReq struct {
	g.Meta `tags:"/sys/balanceCode" method:"delete" path:"/balanceCode" dc:"删除账变类型"`
	Id     uint64 `json:"id"`
}
type DelBalanceCodeRes model.CommonRes
