// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Events is the golang structure for table events.
type Events struct {
	Id               int64       `json:"id"               description:"编号"`
	HomeTeamName     string      `json:"homeTeamName"     description:"主队名称"`
	AwayTeamName     string      `json:"awayTeamName"     description:"客队名称"`
	HomeTeamId       string      `json:"homeTeamId"       description:"主队Id"`
	AwayTeamId       string      `json:"awayTeamId"       description:"客队Id"`
	RestTime         string      `json:"restTime"         description:"进行时间"`
	StartTime        *gtime.Time `json:"startTime"        description:"开始时间"`
	EndTime          *gtime.Time `json:"endTime"          description:"结束时间"`
	LeagueId         int64       `json:"leagueId"         description:"联盟编号"`
	LeagueTitle      string      `json:"leagueTitle"      description:"联盟"`
	EnLeagueTitle    string      `json:"enLeagueTitle"    description:"联盟英文"`
	FirstStatus      int         `json:"firstStatus"      description:"上半场状态 0未开始，1已开始，2已结束"`
	StartDate        *gtime.Time `json:"startDate"        description:"开始日期"`
	FirstOpenResult  string      `json:"firstOpenResult"  description:"上半场开奖结果"`
	FirstOpenTime    *gtime.Time `json:"firstOpenTime"    description:"上半场开奖时间"`
	SecondOpenResult string      `json:"secondOpenResult" description:"下半场开奖结果"`
	SecondOpenTime   *gtime.Time `json:"secondOpenTime"   description:"下半场开奖结果"`
	SecondStatus     int         `json:"secondStatus"     description:"0未开始，1已开始，2已结束"`
	AddTime          *gtime.Time `json:"addTime"          description:"添加时间"`
	IsHot            int         `json:"isHot"            description:"热门"`
	Status           int         `json:"status"           description:"0：未开始 1接受下注，2：停止下注，3 已结算"`
	Remark           string      `json:"remark"           description:"注释"`
	BetMoney         float64     `json:"betMoney"         description:"交易量"`
	AllOpenResult    string      `json:"allOpenResult"    description:"全场开奖结果"`
	AllOpenTime      *gtime.Time `json:"allOpenTime"      description:"全场开奖时间"`
	Handicap         int         `json:"handicap"         description:"让球"`
}
