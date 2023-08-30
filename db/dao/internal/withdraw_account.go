// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WithdrawAccountDao is the data access object for table u_withdraw_account.
type WithdrawAccountDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns WithdrawAccountColumns // columns contains all the column names of Table for convenient usage.
}

// WithdrawAccountColumns defines and stores column names for table u_withdraw_account.
type WithdrawAccountColumns struct {
	Id       string //
	Net      string //
	Protocol string //
	Uid      string //
	Account  string //
	Address  string //
	Currency string // currency
	Status   string //
	Default  string // 是否默认的
	Title    string //
}

// withdrawAccountColumns holds the columns for table u_withdraw_account.
var withdrawAccountColumns = WithdrawAccountColumns{
	Id:       "id",
	Net:      "net",
	Protocol: "protocol",
	Uid:      "uid",
	Account:  "account",
	Address:  "address",
	Currency: "currency",
	Status:   "status",
	Default:  "default",
	Title:    "title",
}

// NewWithdrawAccountDao creates and returns a new DAO object for table data access.
func NewWithdrawAccountDao() *WithdrawAccountDao {
	return &WithdrawAccountDao{
		group:   "default",
		table:   "u_withdraw_account",
		columns: withdrawAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WithdrawAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WithdrawAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WithdrawAccountDao) Columns() WithdrawAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WithdrawAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WithdrawAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WithdrawAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
