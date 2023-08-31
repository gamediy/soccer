// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LeagueDao is the data access object for table p_league.
type LeagueDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns LeagueColumns // columns contains all the column names of Table for convenient usage.
}

// LeagueColumns defines and stores column names for table p_league.
type LeagueColumns struct {
	Id      string // 编号
	Country string // 国家
	Status  string // 状态
	ZhName  string // 中文名称
	EnName  string // 英文名称
	Icon    string //
	VdName  string //
}

// leagueColumns holds the columns for table p_league.
var leagueColumns = LeagueColumns{
	Id:      "id",
	Country: "country",
	Status:  "status",
	ZhName:  "zh_name",
	EnName:  "en_name",
	Icon:    "icon",
	VdName:  "vd_name",
}

// NewLeagueDao creates and returns a new DAO object for table data access.
func NewLeagueDao() *LeagueDao {
	return &LeagueDao{
		group:   "default",
		table:   "p_league",
		columns: leagueColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LeagueDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LeagueDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LeagueDao) Columns() LeagueColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LeagueDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LeagueDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LeagueDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
