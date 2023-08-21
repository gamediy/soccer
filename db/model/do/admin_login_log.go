// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLoginLog is the golang structure of table s_admin_login_log for DAO operations like Where/Data.
type AdminLoginLog struct {
	g.Meta    `orm:"table:s_admin_login_log, do:true"`
	Id        interface{} //
	Uid       interface{} //
	Account   interface{} //
	Ip        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
