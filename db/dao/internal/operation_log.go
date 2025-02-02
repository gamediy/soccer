// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OperationLogDao is the data access object for table s_operation_log.
type OperationLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns OperationLogColumns // columns contains all the column names of Table for convenient usage.
}

// OperationLogColumns defines and stores column names for table s_operation_log.
type OperationLogColumns struct {
	Id        string //
	Uid       string //
	Content   string //
	Response  string //
	Method    string //
	Uri       string //
	Ip        string //
	UseTime   string //
	CreatedAt string //
	MenuName  string //
}

// operationLogColumns holds the columns for table s_operation_log.
var operationLogColumns = OperationLogColumns{
	Id:        "id",
	Uid:       "uid",
	Content:   "content",
	Response:  "response",
	Method:    "method",
	Uri:       "uri",
	Ip:        "ip",
	UseTime:   "use_time",
	CreatedAt: "created_at",
	MenuName:  "menu_name",
}

// NewOperationLogDao creates and returns a new DAO object for table data access.
func NewOperationLogDao() *OperationLogDao {
	return &OperationLogDao{
		group:   "default",
		table:   "s_operation_log",
		columns: operationLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OperationLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OperationLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OperationLogDao) Columns() OperationLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OperationLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OperationLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OperationLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
