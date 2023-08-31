// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeamDao is the data access object for table p_team.
type TeamDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TeamColumns // columns contains all the column names of Table for convenient usage.
}

// TeamColumns defines and stores column names for table p_team.
type TeamColumns struct {
	Id     string //
	ZhName string //
	EnName string //
	VdName string //
	Status string //
	Icon   string //
}

// teamColumns holds the columns for table p_team.
var teamColumns = TeamColumns{
	Id:     "id",
	ZhName: "zh_name",
	EnName: "en_name",
	VdName: "vd_name",
	Status: "status",
	Icon:   "icon",
}

// NewTeamDao creates and returns a new DAO object for table data access.
func NewTeamDao() *TeamDao {
	return &TeamDao{
		group:   "default",
		table:   "p_team",
		columns: teamColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeamDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TeamDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TeamDao) Columns() TeamColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TeamDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TeamDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TeamDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
