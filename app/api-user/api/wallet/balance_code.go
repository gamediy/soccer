package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
)

type BalanceCodesReq struct {
	g.Meta `tags:"钱包" sm:"账变类型" method:"get" path:"/balanceCodes"`
}
type BalanceCodesRes []*entity.BalanceCode
