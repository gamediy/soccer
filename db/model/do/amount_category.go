// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AmountCategory is the golang structure of table u_amount_category for DAO operations like Where/Data.
type AmountCategory struct {
	g.Meta    `orm:"table:u_amount_category, do:true"`
	Id        interface{} //
	Title     interface{} // 标题
	Category  interface{} // 1:区块链，银行卡
	Status    interface{} //
	Type      interface{} // deposit,withdraw
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
