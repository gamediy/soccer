// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletLog is the golang structure for table wallet_log.
type WalletLog struct {
	Id              int64       `json:"id"              description:""`
	OrderNo         int64       `json:"orderNo"         description:"订单号，（有可能是充值提现的订单号）"`
	Uid             int64       `json:"uid"             description:"UID"`
	Account         string      `json:"account"         description:"账号"`
	Pid             int64       `json:"pid"             description:"上级ID"`
	BalanceCode     int         `json:"balanceCode"     description:"余额编号"`
	Title           string      `json:"title"           description:"标题"`
	BalanceBefore   float64     `json:"balanceBefore"   description:"之前余额"`
	BalanceAfter    float64     `json:"balanceAfter"    description:"之后余额"`
	Note            string      `json:"note"            description:"说明"`
	OrderNoRelation int64       `json:"orderNoRelation" description:"关联订单号，（有可能是充值提现的订单号）"`
	ParentPath      string      `json:"parentPath"      description:"上级全路经"`
	Amount          float64     `json:"amount"          description:"金额"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"        description:""`
}
