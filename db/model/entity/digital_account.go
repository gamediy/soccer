// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DigitalAccount is the golang structure for table digital_account.
type DigitalAccount struct {
	Id           int         `json:"id"           description:""`
	Address      string      `json:"address"      description:"地址"`
	Net          string      `json:"net"          description:"网络 TRON ETH"`
	Status       int         `json:"status"       description:"状态 0，1，打开"`
	CountDeposit int         `json:"countDeposit" description:"充值次数"`
	PrivateKey   string      `json:"privateKey"   description:"私密"`
	TotalDeposit int64       `json:"totalDeposit" description:"总充值"`
	Uid          int64       `json:"uid"          description:""`
	Account      string      `json:"account"      description:"账户"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
}
