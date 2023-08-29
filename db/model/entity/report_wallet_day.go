// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReportWalletDay is the golang structure for table report_wallet_day.
type ReportWalletDay struct {
	Id          int64       `json:"id"          description:""`
	Uid         int64       `json:"uid"         description:""`
	Account     string      `json:"account"     description:""`
	Pid         int64       `json:"pid"         description:""`
	ParentPath  string      `json:"parentPath"  description:"上级路经"`
	BalanceCode int         `json:"balanceCode" description:"余额编号"`
	Date        *gtime.Time `json:"date"        description:"日期"`
	Amount      float64     `json:"amount"      description:"余额"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	Count       int         `json:"count"       description:"数量"`
	Title       string      `json:"title"       description:""`
}
