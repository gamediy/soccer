// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Events is the golang structure of table p_events for DAO operations like Where/Data.
type Events struct {
	g.Meta           `orm:"table:p_events, do:true"`
	Id               interface{} // 编号
	HomeTeamName     interface{} // 主队名称
	AwayTeamName     interface{} // 客队名称
	HomeTeamId       interface{} // 主队Id
	AwayTeamId       interface{} // 客队Id
	RestTime         interface{} // 进行时间
	StartTime        *gtime.Time // 开始时间
	EndTime          *gtime.Time // 结束时间
	LeagueId         interface{} // 联盟编号
	LeagueTitle      interface{} // 联盟
	EnLeagueTitle    interface{} // 联盟英文
	FirstStatus      interface{} // 上半场状态 0未开始，1已开始，2已结束
	StartDate        *gtime.Time // 开始日期
	FirstOpenResult  interface{} // 上半场开奖结果
	FirstOpenTime    *gtime.Time // 上半场开奖时间
	SecondOpenResult interface{} // 下半场开奖结果
	SecondOpenTime   *gtime.Time // 下半场开奖结果
	SecondStatus     interface{} // 0未开始，1已开始，2已结束
	AddTime          *gtime.Time // 添加时间
	IsHot            interface{} // 热门
	Status           interface{} // 0：未开始 1接受下注，2：停止下注，3 已结算
	Remark           interface{} // 注释
	BetMoney         interface{} // 交易量
	AllOpenResult    interface{} // 全场开奖结果
	AllOpenTime      *gtime.Time // 全场开奖时间
	Handicap         interface{} // 让球
}
