// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SoccerOrderDao is the data access object for table o_soccer_order.
type SoccerOrderDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SoccerOrderColumns // columns contains all the column names of Table for convenient usage.
}

// SoccerOrderColumns defines and stores column names for table o_soccer_order.
type SoccerOrderColumns struct {
	OrderNo          string // 编号
	Uid              string // 用户编号
	Account          string // 用户账号
	CreatedAt        string // 时间
	EventsId         string // 赛事编号
	EventsTitle      string // 赛事名称
	OddsId           string // 赔率编号
	OddsTitle        string // 赔率标题
	Amount           string // 投注金额
	Profit           string // 赢利
	CalcAt           string // 结算时间
	Status           string // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	OddsCalcRule     string // 结算规则
	Odds             string // 赔率
	LeagueId         string // 联盟编号
	LeagueTitle      string // 联盟
	EventsStartTime  string // 赛事开始时间
	Fee              string // 手续费
	EventsOpenResult string // 赛事开奖结果
	PlayCode         string // PlayId
	BoutStatus       string // 哪个场次
	Pid              string //
	ParentPath       string //
}

// soccerOrderColumns holds the columns for table o_soccer_order.
var soccerOrderColumns = SoccerOrderColumns{
	OrderNo:          "order_no",
	Uid:              "uid",
	Account:          "account",
	CreatedAt:        "created_at",
	EventsId:         "events_id",
	EventsTitle:      "events_title",
	OddsId:           "odds_id",
	OddsTitle:        "odds_title",
	Amount:           "amount",
	Profit:           "profit",
	CalcAt:           "calc_at",
	Status:           "status",
	OddsCalcRule:     "odds_calc_rule",
	Odds:             "odds",
	LeagueId:         "league_id",
	LeagueTitle:      "league_title",
	EventsStartTime:  "events_start_time",
	Fee:              "fee",
	EventsOpenResult: "events_open_result",
	PlayCode:         "play_code",
	BoutStatus:       "bout_status",
	Pid:              "pid",
	ParentPath:       "parent_path",
}

// NewSoccerOrderDao creates and returns a new DAO object for table data access.
func NewSoccerOrderDao() *SoccerOrderDao {
	return &SoccerOrderDao{
		group:   "default",
		table:   "o_soccer_order",
		columns: soccerOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SoccerOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SoccerOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SoccerOrderDao) Columns() SoccerOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SoccerOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SoccerOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SoccerOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
