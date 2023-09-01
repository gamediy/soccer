// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayTypeDao is the data access object for table p_play_type.
type PlayTypeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns PlayTypeColumns // columns contains all the column names of Table for convenient usage.
}

// PlayTypeColumns defines and stores column names for table p_play_type.
type PlayTypeColumns struct {
	Id     string //
	Code   string //
	EnName string //
	ZhName string //
	Status string //
	Class  string //
}

// playTypeColumns holds the columns for table p_play_type.
var playTypeColumns = PlayTypeColumns{
	Id:     "id",
	Code:   "code",
	EnName: "en_name",
	ZhName: "zh_name",
	Status: "status",
	Class:  "class",
}

// NewPlayTypeDao creates and returns a new DAO object for table data access.
func NewPlayTypeDao() *PlayTypeDao {
	return &PlayTypeDao{
		group:   "default",
		table:   "p_play_type",
		columns: playTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PlayTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PlayTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PlayTypeDao) Columns() PlayTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PlayTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PlayTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PlayTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
