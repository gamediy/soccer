// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MenuApiRule is the golang structure of table s_menu_api_rule for DAO operations like Where/Data.
type MenuApiRule struct {
	g.Meta `orm:"table:s_menu_api_rule, do:true"`
	Id     interface{} //
	MenuId interface{} //
	ApiId  interface{} //
}
