// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletWithdrawApplication is the golang structure for table wallet_withdraw_application.
type WalletWithdrawApplication struct {
	Id             uint64      `json:"id"             description:""`
	TransId        string      `json:"transId"        description:""`
	Uid            uint64      `json:"uid"            description:""`
	ChangeType     uint        `json:"changeType"     description:""`
	Money          float64     `json:"money"          description:""`
	ReceiptAmount  float64     `json:"receiptAmount"  description:""`
	Fee            float64     `json:"fee"            description:""`
	Ip             string      `json:"ip"             description:""`
	Desc           string      `json:"desc"           description:""`
	Aid            uint64      `json:"aid"            description:""`
	Status         int         `json:"status"         description:""`
	PayAddress     string      `json:"payAddress"     description:""`
	ReceiveAddress string      `json:"receiveAddress" description:""`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	DescTwo        string      `json:"descTwo"        description:""`
}
