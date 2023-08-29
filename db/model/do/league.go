// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// League is the golang structure of table p_league for DAO operations like Where/Data.
type League struct {
	g.Meta  `orm:"table:p_league, do:true"`
	Id      interface{} // 编号
	Title   interface{} // 标题
	Country interface{} // 国家
	Status  interface{} // 状态
	Name    interface{} // 中文名称
	EnName  interface{} // 英文名称
}
