// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DigitalAccountDao is the data access object for table u_digital_account.
type DigitalAccountDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns DigitalAccountColumns // columns contains all the column names of Table for convenient usage.
}

// DigitalAccountColumns defines and stores column names for table u_digital_account.
type DigitalAccountColumns struct {
	Id           string //
	Address      string // 地址
	Net          string // 网络 TRON ETH
	Status       string // 状态 0，1，打开
	CountDeposit string // 充值次数
	PrivateKey   string // 私密
	TotalDeposit string // 总充值
	Uid          string //
	Account      string // 账户
	CreatedAt    string //
	UpdatedAt    string //
}

// digitalAccountColumns holds the columns for table u_digital_account.
var digitalAccountColumns = DigitalAccountColumns{
	Id:           "id",
	Address:      "address",
	Net:          "net",
	Status:       "status",
	CountDeposit: "count_deposit",
	PrivateKey:   "private_key",
	TotalDeposit: "total_deposit",
	Uid:          "uid",
	Account:      "account",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewDigitalAccountDao creates and returns a new DAO object for table data access.
func NewDigitalAccountDao() *DigitalAccountDao {
	return &DigitalAccountDao{
		group:   "default",
		table:   "u_digital_account",
		columns: digitalAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DigitalAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DigitalAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DigitalAccountDao) Columns() DigitalAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DigitalAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DigitalAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DigitalAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
