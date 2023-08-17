// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WithdrawAccount is the golang structure of table u_withdraw_account for DAO operations like Where/Data.
type WithdrawAccount struct {
	g.Meta   `orm:"table:u_withdraw_account, do:true"`
	Id       interface{} //
	Net      interface{} //
	Protocol interface{} //
	Uid      interface{} //
	Account  interface{} //
	Address  interface{} //
	Currency interface{} // currency
	Status   interface{} //
	Default  interface{} // 是否默认的
}
