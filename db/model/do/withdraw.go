// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Withdraw is the golang structure of table o_withdraw for DAO operations like Where/Data.
type Withdraw struct {
	g.Meta        `orm:"table:o_withdraw, do:true"`
	OrderNo       interface{} // 订单号
	Account       interface{} // 账号
	Uid           interface{} // UID
	Pid           interface{} // 上级ID
	Status        interface{} // 状态 0 提现中，1提现成功，-1提现失败
	FinishAt      *gtime.Time // 完成时间
	Detail        interface{} // 详情
	StatusRemark  interface{} // 状态说明
	Amount        interface{} // 提现金额
	SysRemark     interface{} // 系统说明
	Address       interface{} // 地址（如果是银行卡，就是账号）
	AmountFinally interface{} // 最后倒账金额
	Fee           interface{} // 手续费
	Net           interface{} // 网络（数字货币和银行卡代号）
	AmountItemId  interface{} // 编号
	Currency      interface{} // 货币单位（USDT PHP CNY）
	Protocol      interface{} // 协议
	Title         interface{} // 标题
	Note          interface{} // 说明
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time //
	AdminOperate  interface{} // 操作用户
	ParentPath    interface{} //
}
