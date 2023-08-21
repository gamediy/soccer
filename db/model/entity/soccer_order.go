// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SoccerOrder is the golang structure for table soccer_order.
type SoccerOrder struct {
	Id               int64       `json:"id"               description:""`
	OrderNo          int64       `json:"orderNo"          description:"编号"`
	UserId           int64       `json:"userId"           description:"用户编号"`
	UserAccount      string      `json:"userAccount"      description:"用户账号"`
	AddTime          *gtime.Time `json:"addTime"          description:"时间"`
	EventsId         int64       `json:"eventsId"         description:"赛事编号"`
	EventsTitle      string      `json:"eventsTitle"      description:"赛事名称"`
	OddsId           int64       `json:"oddsId"           description:"赔率编号"`
	OddsTitle        string      `json:"oddsTitle"        description:"赔率标题"`
	Amount           float64     `json:"amount"           description:"投注金额"`
	Profit           float64     `json:"profit"           description:"赢利"`
	CalcTime         *gtime.Time `json:"calcTime"         description:"结算时间"`
	Status           int         `json:"status"           description:"状态 0 撤单，1:投注成功，2：中奖，3：未中奖"`
	CompanyCode      int         `json:"companyCode"      description:"公司编号"`
	OddsCalcRule     string      `json:"oddsCalcRule"     description:"结算规则"`
	OddsProfitRate   float64     `json:"oddsProfitRate"   description:"利率"`
	LeagueId         int64       `json:"leagueId"         description:"联盟编号"`
	LeagueTitle      string      `json:"leagueTitle"      description:"联盟"`
	OddsType         int         `json:"oddsType"         description:"1:半场，2：全场"`
	EventsStartTime  *gtime.Time `json:"eventsStartTime"  description:"赛事开始时间"`
	Fee              float64     `json:"fee"              description:"手续费"`
	EventsOpenResult string      `json:"eventsOpenResult" description:"赛事开奖结果"`
	PlayId           int64       `json:"playId"           description:"PlayId"`
	BoutStatus       int         `json:"boutStatus"       description:"哪个回合"`
}
