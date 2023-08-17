// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventsOdds is the golang structure of table p_events_odds for DAO operations like Where/Data.
type EventsOdds struct {
	g.Meta     `orm:"table:p_events_odds, do:true"`
	Id         interface{} //
	EventsId   interface{} // 赛事编号
	Title      interface{} // 标题
	CalcRule   interface{} // 结算规则
	Type       interface{} // 类型 1:上半场 2：全场
	BetMoney   interface{} // 投注金额
	ProfitRate interface{} // 赢利
	Header     interface{} // 主球队 1：主队，2：客队，draw 平局
	Odds       interface{} // 赔率
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
