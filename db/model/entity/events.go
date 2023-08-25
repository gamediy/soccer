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
	HomeTeam         string      `json:"homeTeam"         description:"主队名称"`
	AwayTeam         string      `json:"awayTeam"         description:"客队名称"`
	EnHomeTeam       string      `json:"enHomeTeam"       description:"主队英文"`
	EnAwayTeam       string      `json:"enAwayTeam"       description:"客队英文"`
	RestTime         string      `json:"restTime"         description:"进行时间"`
	StartTime        *gtime.Time `json:"startTime"        description:"开始时间"`
	EndTime          *gtime.Time `json:"endTime"          description:"结束时间"`
	LeagueId         int64       `json:"leagueId"         description:"联盟编号"`
	LeagueTitle      string      `json:"leagueTitle"      description:"联盟"`
	EnLeagueTitle    string      `json:"enLeagueTitle"    description:"联盟英文"`
	HalfStatus       int         `json:"halfStatus"       description:"半场状态 0未开始，1已开始，2已结束"`
	FirstStatus      int         `json:"firstStatus"      description:"上半场状态 0未开始，1已开始，2已结束"`
	StartDate        *gtime.Time `json:"startDate"        description:"开始日期"`
	FirstOpenResult  string      `json:"firstOpenResult"  description:"上半场开奖结果"`
	FirstOpenTime    *gtime.Time `json:"firstOpenTime"    description:"上半场开奖时间"`
	SecondOpenResult string      `json:"secondOpenResult" description:"下半场开奖结果"`
	SecondOpenTime   *gtime.Time `json:"secondOpenTime"   description:"下半场开奖结果"`
	SecondStatus     int         `json:"secondStatus"     description:"全场状态"`
	AddTime          *gtime.Time `json:"addTime"          description:"添加时间"`
	IsHot            int         `json:"isHot"            description:"热门"`
	Status           int         `json:"status"           description:"状态"`
	Remark           string      `json:"remark"           description:"注释"`
	BetMoney         float64     `json:"betMoney"         description:"交易量"`
	OpenResult       string      `json:"openResult"       description:""`
	HalfOpenResult   string      `json:"halfOpenResult"   description:"上半场开奖结果"`
	HalfOpenTime     *gtime.Time `json:"halfOpenTime"     description:"上半场开奖时间"`
	AllOpenResult    string      `json:"allOpenResult"    description:"全场开奖结果"`
	AllOpenTime      *gtime.Time `json:"allOpenTime"      description:"全场开奖时间"`
	AllStatus        int         `json:"allStatus"        description:"全场状态"`
	FiId             int64       `json:"fiId"             description:"FI编号"`
	Handicap         string      `json:"handicap"         description:"让球"`
}
