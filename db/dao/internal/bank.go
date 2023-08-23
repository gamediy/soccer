// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BankDao is the data access object for table b_bank.
type BankDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns BankColumns // columns contains all the column names of Table for convenient usage.
}

// BankColumns defines and stores column names for table b_bank.
type BankColumns struct {
	Id       string //
	Icon     string // 图标
	Currency string //
	Net      string // 类型
	Protocol string // 名称
	Summary  string // 简介
	Url      string // 官网
	Status   string // 状态 1开启 2关闭
	Class    string //
}

// bankColumns holds the columns for table b_bank.
var bankColumns = BankColumns{
	Id:       "id",
	Icon:     "icon",
	Currency: "currency",
	Net:      "net",
	Protocol: "protocol",
	Summary:  "summary",
	Url:      "url",
	Status:   "status",
	Class:    "class",
}

// NewBankDao creates and returns a new DAO object for table data access.
func NewBankDao() *BankDao {
	return &BankDao{
		group:   "default",
		table:   "b_bank",
		columns: bankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BankDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BankDao) Columns() BankColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BankDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
