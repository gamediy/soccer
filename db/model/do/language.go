// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Language is the golang structure of table c_language for DAO operations like Where/Data.
type Language struct {
	g.Meta    `orm:"table:c_language, do:true"`
	Id        interface{} //
	Name      interface{} //
	Desc      interface{} //
	En        interface{} // 英语
	Fr        interface{} // 法语
	Ja        interface{} // 日语
	Hi        interface{} // 印度语
	Zh        interface{} // 中文
	Ko        interface{} // 汉语
	Pt        interface{} // 葡萄牙语
	Es        interface{} // 西班牙语
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Hu        interface{} // 匈牙利语
	Ar        interface{} // 阿拉伯语
}
