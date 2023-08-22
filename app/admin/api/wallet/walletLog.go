package wallet

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"star_net/model"
)

type ReadWalletLogReq struct {
	g.Meta `tags:"/order/walletLog" method:"get" path:"/walletLog" dc:"查询账变日志"`
	Id     uint64 `json:"id"`
}
type ReadWalletLogRes entity.WalletLog
type ReadListWalletLogReq struct {
	g.Meta `tags:"/order/walletLog" method:"get" path:"/walletLog/list" dc:"查询账变日志列表"`
	model.CommonPageReq
	Account     string
	OrderNo     string
	BalanceCode string
	Note        string
}
type ReadListWalletLogRes struct {
	Total int                 `json:"total"`
	List  []*entity.WalletLog `json:"list"`
}
type CreateWalletLogReq struct {
	g.Meta `tags:"/order/walletLog" method:"post" path:"/walletLog" dc:"添加账变日志"`
	*entity.WalletLog
}
type CreateWalletLogRes model.CommonRes
type UpdateWalletLogReq struct {
	g.Meta `tags:"/order/walletLog" method:"put" path:"/walletLog" dc:"修改账变日志"`
	*entity.WalletLog
}
type UpdateWalletLogRes model.CommonRes
type DelWalletLogReq struct {
	g.Meta `tags:"/order/walletLog" method:"delete" path:"/walletLog" dc:"删除账变日志"`
	Id     uint64 `json:"id"`
}
type DelWalletLogRes model.CommonRes
