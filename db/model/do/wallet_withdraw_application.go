// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletWithdrawApplication is the golang structure of table u_wallet_withdraw_application for DAO operations like Where/Data.
type WalletWithdrawApplication struct {
	g.Meta         `orm:"table:u_wallet_withdraw_application, do:true"`
	Id             interface{} //
	TransId        interface{} //
	Uid            interface{} //
	ChangeType     interface{} //
	Money          interface{} //
	ReceiptAmount  interface{} //
	Fee            interface{} //
	Ip             interface{} //
	Desc           interface{} //
	Aid            interface{} //
	Status         interface{} //
	PayAddress     interface{} //
	ReceiveAddress interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DescTwo        interface{} //
}
