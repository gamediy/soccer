// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table u_user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:u_user, do:true"`
	Id           interface{} //
	Account      interface{} // 账号
	Email        interface{} // 邮箱
	Password     interface{} // 密码
	Status       interface{} // 1:正常，2：冻结
	Xid          interface{} // short code 邀请 短码
	Ip           interface{} // IP
	ClientAgent  interface{} // 注册clientAgen头
	Phone        interface{} // 电报
	LevelBet     interface{} // 下注的等级
	LevelDeposit interface{} // 充值的等级
	LevelAgent   interface{} // 代理的等级
	Pid          interface{} // 上级ID
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	ParentPath   interface{} // 上级全路经 /1/2/3/
	Country      interface{} // 国家地区
	Lang         interface{} // 用户语言
	AreaCode     interface{} //
	City         interface{} //
	RealName     interface{} // real name
	PayPass      interface{} // 交易密码
}
