// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventsOdds is the golang structure for table events_odds.
type EventsOdds struct {
	Id           int64       `json:"id"           description:""`
	EventsId     int64       `json:"eventsId"     description:"赛事编号"`
	Title        string      `json:"title"        description:"标题"`
	CalcRule     string      `json:"calcRule"     description:"结算规则"`
	BoutStatus   int         `json:"boutStatus"   description:"类型 1:上半场 2：全场"`
	TotalAmount  float64     `json:"totalAmount"  description:"投注金额"`
	TotalProfit  float64     `json:"totalProfit"  description:"赢利"`
	Header       string      `json:"header"       description:"主球队 1：主队，2：客队，draw 平局"`
	Odds         float64     `json:"odds"         description:"赔率"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
	PlayCode     int         `json:"playCode"     description:""`
	Status       int         `json:"status"       description:""`
	PlayTypeCode int         `json:"playTypeCode" description:""`
}
