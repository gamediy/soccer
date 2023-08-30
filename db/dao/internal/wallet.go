// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WalletDao is the data access object for table u_wallet.
type WalletDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns WalletColumns // columns contains all the column names of Table for convenient usage.
}

// WalletColumns defines and stores column names for table u_wallet.
type WalletColumns struct {
	Id            string //
	Uid           string //
	Balance       string // 余额
	TotalBet      string // 总投注
	TotalDeposit  string // 总充值
	TotalWithdraw string // 总提现
	Freeze        string // 冻结
	Account       string // 账号
	ParentPath    string // 上级路经
	TotalProfit   string // 总盈利
	TotalGift     string // 总赠送
	CreatedAt     string //
	UpdatedAt     string //
}

// walletColumns holds the columns for table u_wallet.
var walletColumns = WalletColumns{
	Id:            "id",
	Uid:           "uid",
	Balance:       "balance",
	TotalBet:      "total_bet",
	TotalDeposit:  "total_deposit",
	TotalWithdraw: "total_withdraw",
	Freeze:        "freeze",
	Account:       "account",
	ParentPath:    "parent_path",
	TotalProfit:   "total_profit",
	TotalGift:     "total_gift",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewWalletDao creates and returns a new DAO object for table data access.
func NewWalletDao() *WalletDao {
	return &WalletDao{
		group:   "default",
		table:   "u_wallet",
		columns: walletColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WalletDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WalletDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WalletDao) Columns() WalletColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WalletDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WalletDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WalletDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
