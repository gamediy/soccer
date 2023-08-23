package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
)

type BanksReq struct {
	g.Meta `tags:"钱包" sm:"银行列表" method:"get" path:"/banks"`
}
type BanksRes []*entity.Bank
