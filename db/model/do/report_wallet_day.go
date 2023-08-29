// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReportWalletDay is the golang structure of table r_report_wallet_day for DAO operations like Where/Data.
type ReportWalletDay struct {
	g.Meta      `orm:"table:r_report_wallet_day, do:true"`
	Id          interface{} //
	Uid         interface{} //
	Account     interface{} //
	Pid         interface{} //
	ParentPath  interface{} // 上级路经
	BalanceCode interface{} // 余额编号
	Date        *gtime.Time // 日期
	Amount      interface{} // 余额
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	Count       interface{} // 数量
	Title       interface{} //
}
