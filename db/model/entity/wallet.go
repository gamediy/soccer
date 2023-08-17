// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallet is the golang structure for table wallet.
type Wallet struct {
	Id            int64       `json:"id"            description:""`
	Uid           int64       `json:"uid"           description:""`
	Balance       int64       `json:"balance"       description:"余额"`
	TotalBet      float64     `json:"totalBet"      description:"总投注"`
	TotalDeposit  float64     `json:"totalDeposit"  description:"总充值"`
	TotalWithdraw float64     `json:"totalWithdraw" description:"总提现"`
	Freeze        int64       `json:"freeze"        description:"冻结"`
	Account       string      `json:"account"       description:"账号"`
	ParentPath    string      `json:"parentPath"    description:"上级路经"`
	TotalProfit   float64     `json:"totalProfit"   description:"总盈利"`
	TotalGift     float64     `json:"totalGift"     description:"总赠送"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
	TotalFreeze   float64     `json:"totalFreeze"   description:""`
}
