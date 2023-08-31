// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventsDao is the data access object for table p_events.
type EventsDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns EventsColumns // columns contains all the column names of Table for convenient usage.
}

// EventsColumns defines and stores column names for table p_events.
type EventsColumns struct {
	Id               string // 编号
	HomeTeamName     string // 主队名称
	AwayTeamName     string // 客队名称
	HomeTeamId       string // 主队Id
	AwayTeamId       string // 客队Id
	RestTime         string // 进行时间
	StartTime        string // 开始时间
	EndTime          string // 结束时间
	LeagueId         string // 联盟编号
	LeagueName       string // 联盟
	FirstStatus      string // 上半场状态 0未开始，1已开始，2已结束
	StartDate        string // 开始日期
	FirstOpenResult  string // 上半场开奖结果
	FirstOpenTime    string // 上半场开奖时间
	SecondOpenResult string // 下半场开奖结果
	SecondOpenTime   string // 下半场开奖结果
	SecondStatus     string // 0未开始，1已开始，2已结束
	AddTime          string // 添加时间
	IsHot            string // 热门
	Status           string // 0：未开始 1接受下注，2：停止下注，3 已结算
	Remark           string // 注释
	BetMoney         string // 交易量
	AllOpenResult    string // 全场开奖结果
	AllOpenTime      string // 全场开奖时间
	Handicap         string // 让球
}

// eventsColumns holds the columns for table p_events.
var eventsColumns = EventsColumns{
	Id:               "id",
	HomeTeamName:     "home_team_name",
	AwayTeamName:     "away_team_name",
	HomeTeamId:       "home_team_id",
	AwayTeamId:       "away_team_id",
	RestTime:         "rest_time",
	StartTime:        "start_time",
	EndTime:          "end_time",
	LeagueId:         "league_id",
	LeagueName:       "league_name",
	FirstStatus:      "first_status",
	StartDate:        "start_date",
	FirstOpenResult:  "first_open_result",
	FirstOpenTime:    "first_open_time",
	SecondOpenResult: "second_open_result",
	SecondOpenTime:   "second_open_time",
	SecondStatus:     "second_status",
	AddTime:          "add_time",
	IsHot:            "is_hot",
	Status:           "status",
	Remark:           "remark",
	BetMoney:         "bet_money",
	AllOpenResult:    "all_open_result",
	AllOpenTime:      "all_open_time",
	Handicap:         "handicap",
}

// NewEventsDao creates and returns a new DAO object for table data access.
func NewEventsDao() *EventsDao {
	return &EventsDao{
		group:   "default",
		table:   "p_events",
		columns: eventsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventsDao) Columns() EventsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
