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
	g.Meta         `orm:"table:p_events, do:true"`
	Id             interface{} // 编号
	HomeTeam       interface{} // 主队名称
	AwayTeam       interface{} // 客队名称
	EnHomeTeam     interface{} // 主队英文
	EnAwayTeam     interface{} // 客队英文
	RestTime       interface{} // 进行时间
	StartTime      *gtime.Time // 开始时间
	EndTime        *gtime.Time // 结束时间
	LeagueId       interface{} // 联盟编号
	LeagueTitle    interface{} // 联盟
	EnLeagueTitle  interface{} // 联盟英文
	HalfStatus     interface{} // 半场状态 0未开始，1已开始，2已结束
	StartDate      *gtime.Time // 开始日期
	HalfOpenResult interface{} // 上半场开奖结果
	HalfOpenTime   *gtime.Time // 上半场开奖时间
	AllOpenResult  interface{} // 全场开奖结果
	AllOpenTime    *gtime.Time // 全场开奖时间
	AllStatus      interface{} // 全场状态
	FiId           interface{} // FI编号
	AddTime        *gtime.Time // 添加时间
	IsHot          interface{} // 热门
	Status         interface{} // 状态
	Remark         interface{} // 注释
	BetMoney       interface{} // 交易量
}
