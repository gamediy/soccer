// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table s_admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:s_admin, do:true"`
	Id        interface{} //
	Rid       interface{} // 角色ID
	Uname     interface{} //
	Pwd       interface{} //
	Nickname  interface{} //
	Email     interface{} //
	Phone     interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	KeySecret interface{} //
}
