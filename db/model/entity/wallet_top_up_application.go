// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletTopUpApplication is the golang structure for table wallet_top_up_application.
type WalletTopUpApplication struct {
	Id             uint64      `json:"id"             description:""`
	TransId        string      `json:"transId"        description:""`
	Uid            uint64      `json:"uid"            description:""`
	ChangeType     uint        `json:"changeType"     description:""`
	Money          float64     `json:"money"          description:""`
	ReceiptAmount  float64     `json:"receiptAmount"  description:""`
	Fee            float64     `json:"fee"            description:""`
	Ip             string      `json:"ip"             description:""`
	Description    string      `json:"description"    description:""`
	Aid            uint64      `json:"aid"            description:""`
	Status         uint        `json:"status"         description:"1 wait 2 success 3 fail"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	PayAddress     string      `json:"payAddress"     description:""`
	ReceiveAddress string      `json:"receiveAddress" description:""`
	Hash           string      `json:"hash"           description:"tx_id"`
	Icon           string      `json:"icon"           description:""`
}
