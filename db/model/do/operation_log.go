// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure of table s_operation_log for DAO operations like Where/Data.
type OperationLog struct {
	g.Meta    `orm:"table:s_operation_log, do:true"`
	Id        interface{} //
	Uid       interface{} //
	Content   interface{} //
	Response  interface{} //
	Method    interface{} //
	Uri       interface{} //
	Ip        interface{} //
	UseTime   interface{} //
	CreatedAt *gtime.Time //
	MenuName  interface{} //
}
