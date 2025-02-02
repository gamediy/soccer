// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletTopUpApplication is the golang structure of table u_wallet_top_up_application for DAO operations like Where/Data.
type WalletTopUpApplication struct {
	g.Meta         `orm:"table:u_wallet_top_up_application, do:true"`
	Id             interface{} //
	TransId        interface{} //
	Uid            interface{} //
	ChangeType     interface{} //
	Money          interface{} //
	ReceiptAmount  interface{} //
	Fee            interface{} //
	Ip             interface{} //
	Description    interface{} //
	Aid            interface{} //
	Status         interface{} // 1 wait 2 success 3 fail
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	PayAddress     interface{} //
	ReceiveAddress interface{} //
	Hash           interface{} // tx_id
	Icon           interface{} //
}
