// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceCode is the golang structure for table balance_code.
type BalanceCode struct {
	Id        int         `json:"id"        description:""`
	Title     string      `json:"title"     description:"标题"`
	Code      int         `json:"code"      description:""`
	Remark    string      `json:"remark"    description:"说明"`
	Type      string      `json:"type"      description:"类型"`
	Class     string      `json:"class"     description:"default,success,primary,info,warning,error,secondary,"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
