// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table s_menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta     `orm:"table:s_menu, do:true"`
	Id         interface{} //
	Pid        interface{} //
	Icon       interface{} //
	BgImg      interface{} //
	Name       interface{} //
	Path       interface{} //
	Sort       interface{} //
	Type       interface{} // 1normal 2menu 3 button
	Desc       interface{} //
	FilePath   interface{} //
	Status     interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	Permission interface{} // 权限标识
}
