// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AmountItem is the golang structure of table u_amount_item for DAO operations like Where/Data.
type AmountItem struct {
	g.Meta           `orm:"table:u_amount_item, do:true"`
	Id               interface{} //
	Title            interface{} // 标题
	Status           interface{} // 状态
	Detail           interface{} // 详情
	AmountCategoryId interface{} //
	Net              interface{} // tron eth
	Min              interface{} // 最小金额
	Max              interface{} // 最大金额
	Fee              interface{} // 手续费
	Type             interface{} // 类型
	Logo             interface{} // Logo
	Sort             interface{} // 排序大到小
	Country          interface{} // 国家地区
	Currency         interface{} // 货币单位
	Protocol         interface{} // 协议
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	Address          interface{} // 地址或卡号
}
