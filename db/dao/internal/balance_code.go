// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceCodeDao is the data access object for table u_balance_code.
type BalanceCodeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BalanceCodeColumns // columns contains all the column names of Table for convenient usage.
}

// BalanceCodeColumns defines and stores column names for table u_balance_code.
type BalanceCodeColumns struct {
	Id        string //
	Title     string // 标题
	Code      string //
	Remark    string // 说明
	Type      string // 类型
	Class     string // default,success,primary,info,warning,error,secondary,
	Status    string // 状态
	CreatedAt string //
	UpdatedAt string //
}

// balanceCodeColumns holds the columns for table u_balance_code.
var balanceCodeColumns = BalanceCodeColumns{
	Id:        "id",
	Title:     "title",
	Code:      "code",
	Remark:    "remark",
	Type:      "type",
	Class:     "class",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewBalanceCodeDao creates and returns a new DAO object for table data access.
func NewBalanceCodeDao() *BalanceCodeDao {
	return &BalanceCodeDao{
		group:   "default",
		table:   "u_balance_code",
		columns: balanceCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BalanceCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BalanceCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BalanceCodeDao) Columns() BalanceCodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BalanceCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BalanceCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BalanceCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
