package report

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type WalletLogReq struct {
	g.Meta  `tags:"/report/walletlog" method:"get" path:"/walletlog" dc:"账变报表"`
	Account string `json:"account"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}
type WalletLogRes struct {
	List  gdb.List `json:"list"`
	Sum   gdb.List `json:"sum"`
	Total int64    `json:"total"`
}
