// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallet is the golang structure of table u_wallet for DAO operations like Where/Data.
type Wallet struct {
	g.Meta        `orm:"table:u_wallet, do:true"`
	Id            interface{} //
	Uid           interface{} //
	Balance       interface{} // 余额
	TotalBet      interface{} // 总投注
	TotalDeposit  interface{} // 总充值
	TotalWithdraw interface{} // 总提现
	Freeze        interface{} // 冻结
	Account       interface{} // 账号
	ParentPath    interface{} // 上级路经
	TotalProfit   interface{} // 总盈利
	TotalGift     interface{} // 总赠送
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
