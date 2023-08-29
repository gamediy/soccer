// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletLog is the golang structure of table u_wallet_log for DAO operations like Where/Data.
type WalletLog struct {
	g.Meta          `orm:"table:u_wallet_log, do:true"`
	Id              interface{} //
	OrderNo         interface{} // 订单号，（有可能是充值提现的订单号）
	Uid             interface{} // UID
	Account         interface{} // 账号
	Pid             interface{} // 上级ID
	BalanceCode     interface{} // 余额编号
	Title           interface{} // 标题
	BalanceBefore   interface{} // 之前余额
	BalanceAfter    interface{} // 之后余额
	Note            interface{} // 说明
	OrderNoRelation interface{} // 关联订单号，（有可能是充值提现的订单号）
	ParentPath      interface{} // 上级全路经
	Amount          interface{} // 金额
	CreatedAt       *gtime.Time // 创建时间
	UpdateAt        *gtime.Time //
}
