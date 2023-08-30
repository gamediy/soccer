// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventsOddsDao is the data access object for table p_events_odds.
type EventsOddsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns EventsOddsColumns // columns contains all the column names of Table for convenient usage.
}

// EventsOddsColumns defines and stores column names for table p_events_odds.
type EventsOddsColumns struct {
	Id           string //
	EventsId     string // 赛事编号
	Title        string // 标题
	CalcRule     string // 结算规则
	BoutStatus   string // 类型 1:上半场 2：下半场，3：全场
	TotalAmount  string // 投注金额
	TotalProfit  string // 赢利
	Header       string // 主球队 1：主队，2：客队，draw 平局
	Odds         string // 赔率
	CreatedAt    string //
	UpdatedAt    string //
	PlayCode     string //
	Status       string //
	PlayTypeCode string //
}

// eventsOddsColumns holds the columns for table p_events_odds.
var eventsOddsColumns = EventsOddsColumns{
	Id:           "id",
	EventsId:     "events_id",
	Title:        "title",
	CalcRule:     "calc_rule",
	BoutStatus:   "bout_status",
	TotalAmount:  "total_amount",
	TotalProfit:  "total_profit",
	Header:       "header",
	Odds:         "odds",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	PlayCode:     "play_code",
	Status:       "status",
	PlayTypeCode: "play_type_code",
}

// NewEventsOddsDao creates and returns a new DAO object for table data access.
func NewEventsOddsDao() *EventsOddsDao {
	return &EventsOddsDao{
		group:   "default",
		table:   "p_events_odds",
		columns: eventsOddsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventsOddsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventsOddsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventsOddsDao) Columns() EventsOddsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventsOddsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventsOddsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventsOddsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
