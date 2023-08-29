// =================================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DigitalAccount is the golang structure of table u_digital_account for DAO operations like Where/Data.
type DigitalAccount struct {
	g.Meta       `orm:"table:u_digital_account, do:true"`
	Id           interface{} //
	Address      interface{} // 地址
	Net          interface{} // 网络 TRON ETH
	Status       interface{} // 状态 0，1，打开
	CountDeposit interface{} // 充值次数
	PrivateKey   interface{} // 私密
	TotalDeposit interface{} // 总充值
	Uid          interface{} //
	Account      interface{} // 账户
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
