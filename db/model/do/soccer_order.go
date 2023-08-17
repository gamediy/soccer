// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SoccerOrder is the golang structure of table o_soccer_order for DAO operations like Where/Data.
type SoccerOrder struct {
	g.Meta           `orm:"table:o_soccer_order, do:true"`
	Id               interface{} //
	OrderNo          interface{} // 编号
	UserId           interface{} // 用户编号
	UserAccount      interface{} // 用户账号
	AddTime          *gtime.Time // 时间
	EventsId         interface{} // 赛事编号
	EventsTitle      interface{} // 赛事名称
	OddsId           interface{} // 赔率编号
	OddsTitle        interface{} // 赔率标题
	BetMoney         interface{} // 投注金额
	Profit           interface{} // 赢利
	CalcTime         *gtime.Time // 结算时间
	Status           interface{} // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	CompanyCode      interface{} // 公司编号
	OddsCalcRule     interface{} // 结算规则
	OddsProfitRate   interface{} // 利率
	LeagueId         interface{} // 联盟编号
	LeagueTitle      interface{} // 联盟
	OddsType         interface{} // 1:半场，2：全场
	EventsStartTime  *gtime.Time // 赛事开始时间
	Fee              interface{} // 手续费
	EventsOpenResult interface{} // 赛事开奖结果
	PlayId           interface{} // PlayId
}
