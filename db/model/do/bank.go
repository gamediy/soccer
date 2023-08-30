// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Bank is the golang structure of table b_bank for DAO operations like Where/Data.
type Bank struct {
	g.Meta   `orm:"table:b_bank, do:true"`
	Id       interface{} //
	Icon     interface{} // 图标
	Currency interface{} //
	Net      interface{} // 类型
	Protocol interface{} // 名称
	Summary  interface{} // 简介
	Url      interface{} // 官网
	Status   interface{} // 状态 1开启 2关闭
	Class    interface{} //
}
