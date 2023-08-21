// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SoccerOrderSettle is the golang structure of table o_soccer_order_settle for DAO operations like Where/Data.
type SoccerOrderSettle struct {
	g.Meta           `orm:"table:o_soccer_order_settle, do:true"`
	Id               interface{} //
	OrderNo          interface{} // 编号
	Uid              interface{} // 用户编号
	Account          interface{} // 用户账号
	CreatedAt        *gtime.Time // 时间
	EventsId         interface{} // 赛事编号
	EventsTitle      interface{} // 赛事名称
	OddsId           interface{} // 赔率编号
	OddsTitle        interface{} // 赔率标题
	Amount           interface{} // 投注金额
	Profit           interface{} // 赢利
	CalcAt           *gtime.Time // 结算时间
	Status           interface{} // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	OddsCalcRule     interface{} // 结算规则
	OddsProfitRate   interface{} // 利率
	LeagueId         interface{} // 联盟编号
	LeagueTitle      interface{} // 联盟
	EventsStartTime  *gtime.Time // 赛事开始时间
	Fee              interface{} // 手续费
	EventsOpenResult interface{} // 赛事开奖结果
	PlayCode         interface{} // PlayId
	BoutStatus       interface{} // 哪个场次
	Pid              interface{} //
	ParentPath       interface{} //
}
