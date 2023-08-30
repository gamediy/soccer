// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceCode is the golang structure of table u_balance_code for DAO operations like Where/Data.
type BalanceCode struct {
	g.Meta    `orm:"table:u_balance_code, do:true"`
	Id        interface{} //
	Title     interface{} // 标题
	Code      interface{} //
	Remark    interface{} // 说明
	Type      interface{} // 类型
	Class     interface{} // default,success,primary,info,warning,error,secondary,
	Status    interface{} // 状态
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
