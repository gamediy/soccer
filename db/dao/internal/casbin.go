// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CasbinDao is the data access object for table s_casbin.
type CasbinDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns CasbinColumns // columns contains all the column names of Table for convenient usage.
}

// CasbinColumns defines and stores column names for table s_casbin.
type CasbinColumns struct {
	Ptype string //
	V0    string //
	V1    string //
	V2    string //
	V3    string //
	V4    string //
	V5    string //
}

// casbinColumns holds the columns for table s_casbin.
var casbinColumns = CasbinColumns{
	Ptype: "ptype",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

// NewCasbinDao creates and returns a new DAO object for table data access.
func NewCasbinDao() *CasbinDao {
	return &CasbinDao{
		group:   "default",
		table:   "s_casbin",
		columns: casbinColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CasbinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CasbinDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CasbinDao) Columns() CasbinColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CasbinDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CasbinDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CasbinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
