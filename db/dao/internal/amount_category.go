// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AmountCategoryDao is the data access object for table u_amount_category.
type AmountCategoryDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AmountCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// AmountCategoryColumns defines and stores column names for table u_amount_category.
type AmountCategoryColumns struct {
	Id        string //
	Title     string // 标题
	Category  string // 1:区块链，银行卡
	Status    string //
	Type      string // order,withdraw
	CreatedAt string //
	UpdatedAt string //
}

// amountCategoryColumns holds the columns for table u_amount_category.
var amountCategoryColumns = AmountCategoryColumns{
	Id:        "id",
	Title:     "title",
	Category:  "category",
	Status:    "status",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAmountCategoryDao creates and returns a new DAO object for table data access.
func NewAmountCategoryDao() *AmountCategoryDao {
	return &AmountCategoryDao{
		group:   "default",
		table:   "u_amount_category",
		columns: amountCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AmountCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AmountCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AmountCategoryDao) Columns() AmountCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AmountCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AmountCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AmountCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
