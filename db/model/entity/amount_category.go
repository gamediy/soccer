// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AmountCategory is the golang structure for table amount_category.
type AmountCategory struct {
	Id        int         `json:"id"        description:""`
	Title     string      `json:"title"     description:"标题"`
	Category  string      `json:"category"  description:"1:区块链，银行卡"`
	Status    int         `json:"status"    description:""`
	Type      string      `json:"type"      description:"order,withdraw"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
