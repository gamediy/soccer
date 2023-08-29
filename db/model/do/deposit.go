// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Deposit is the golang structure of table o_deposit for DAO operations like Where/Data.
type Deposit struct {
	g.Meta          `orm:"table:o_deposit, do:true"`
	OrderNo         interface{} // 订单号
	Account         interface{} // 账号
	Uid             interface{} // UID
	Pid             interface{} // 上级ID
	Status          interface{} // 状态
	FinishAt        *gtime.Time // 完成时间
	Detail          interface{} // 详情
	StatusRemark    interface{} // 状态说明
	Amount          interface{} // 充值金额
	SysRemark       interface{} // 系统说明
	Address         interface{} // 地址或者是卡号
	Net             interface{} // 类型
	AmountItemId    interface{} // 充值的编号
	Currency        interface{} // 货币单位
	Protocol        interface{} // 协议
	ParentPath      interface{} // 上级代理全路经
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
	AdminOperate    interface{} // 操作用户
	TransferOrderNo interface{} // 用户转账订单号
	TransferImg     interface{} // 用户转账图片
}
