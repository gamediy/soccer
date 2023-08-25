// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SoccerOrderSettle is the golang structure for table soccer_order_settle.
type SoccerOrderSettle struct {
	OrderNo          int64       `json:"orderNo"          description:"编号"`
	Uid              int64       `json:"uid"              description:"用户编号"`
	Account          string      `json:"account"          description:"用户账号"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:"时间"`
	EventsId         int64       `json:"eventsId"         description:"赛事编号"`
	EventsTitle      string      `json:"eventsTitle"      description:"赛事名称"`
	OddsId           int64       `json:"oddsId"           description:"赔率编号"`
	OddsTitle        string      `json:"oddsTitle"        description:"赔率标题"`
	Amount           float64     `json:"amount"           description:"投注金额"`
	Profit           float64     `json:"profit"           description:"赢利"`
	CalcAt           *gtime.Time `json:"calcAt"           description:"结算时间"`
	Status           int         `json:"status"           description:"状态 0 撤单，1:投注成功，2：中奖，3：未中奖"`
	OddsCalcRule     string      `json:"oddsCalcRule"     description:"结算规则"`
	Odds             float64     `json:"odds"             description:"赔率"`
	LeagueId         int64       `json:"leagueId"         description:"联盟编号"`
	LeagueTitle      string      `json:"leagueTitle"      description:"联盟"`
	EventsStartTime  *gtime.Time `json:"eventsStartTime"  description:"赛事开始时间"`
	Fee              float64     `json:"fee"              description:"手续费"`
	EventsOpenResult string      `json:"eventsOpenResult" description:"赛事开奖结果"`
	PlayCode         int         `json:"playCode"         description:"PlayId"`
	BoutStatus       int         `json:"boutStatus"       description:"哪个场次"`
	Pid              int64       `json:"pid"              description:""`
	ParentPath       string      `json:"parentPath"       description:""`
}
