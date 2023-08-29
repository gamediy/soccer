// ==========================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReportWalletDayDao is the data access object for table r_report_wallet_day.
type ReportWalletDayDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ReportWalletDayColumns // columns contains all the column names of Table for convenient usage.
}

// ReportWalletDayColumns defines and stores column names for table r_report_wallet_day.
type ReportWalletDayColumns struct {
	Id          string //
	Uid         string //
	Account     string //
	Pid         string //
	ParentPath  string // 上级路经
	BalanceCode string // 余额编号
	Date        string // 日期
	Amount      string // 余额
	CreatedAt   string //
	UpdatedAt   string //
	Count       string // 数量
	Title       string //
}

// reportWalletDayColumns holds the columns for table r_report_wallet_day.
var reportWalletDayColumns = ReportWalletDayColumns{
	Id:          "id",
	Uid:         "uid",
	Account:     "account",
	Pid:         "pid",
	ParentPath:  "parent_path",
	BalanceCode: "balance_code",
	Date:        "date",
	Amount:      "amount",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Count:       "count",
	Title:       "title",
}

// NewReportWalletDayDao creates and returns a new DAO object for table data access.
func NewReportWalletDayDao() *ReportWalletDayDao {
	return &ReportWalletDayDao{
		group:   "default",
		table:   "r_report_wallet_day",
		columns: reportWalletDayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReportWalletDayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReportWalletDayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReportWalletDayDao) Columns() ReportWalletDayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReportWalletDayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReportWalletDayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReportWalletDayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
