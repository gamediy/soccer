// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WithdrawDao is the data access object for table o_withdraw.
type WithdrawDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns WithdrawColumns // columns contains all the column names of Table for convenient usage.
}

// WithdrawColumns defines and stores column names for table o_withdraw.
type WithdrawColumns struct {
	OrderNo       string // 订单号
	Account       string // 账号
	Uid           string // UID
	Pid           string // 上级ID
	Status        string // 状态 0 提现中，1提现成功，-1提现失败
	FinishAt      string // 完成时间
	Detail        string // 详情
	StatusRemark  string // 状态说明
	Amount        string // 提现金额
	SysRemark     string // 系统说明
	Address       string // 地址（如果是银行卡，就是账号）
	AmountFinally string // 最后倒账金额
	Fee           string // 手续费
	Net           string // 网络（数字货币和银行卡代号）
	AmountItemId  string // 编号
	Currency      string // 货币单位（USDT PHP CNY）
	Protocol      string // 协议
	Title         string // 标题
	Note          string // 说明
	CreatedAt     string // 创建时间
	UpdatedAt     string //
	AdminOperate  string // 操作用户
	ParentPath    string //
}

// withdrawColumns holds the columns for table o_withdraw.
var withdrawColumns = WithdrawColumns{
	OrderNo:       "order_no",
	Account:       "account",
	Uid:           "uid",
	Pid:           "pid",
	Status:        "status",
	FinishAt:      "finish_at",
	Detail:        "detail",
	StatusRemark:  "status_remark",
	Amount:        "amount",
	SysRemark:     "sys_remark",
	Address:       "address",
	AmountFinally: "amount_finally",
	Fee:           "fee",
	Net:           "net",
	AmountItemId:  "amount_item_id",
	Currency:      "currency",
	Protocol:      "protocol",
	Title:         "title",
	Note:          "note",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	AdminOperate:  "admin_operate",
	ParentPath:    "parent_path",
}

// NewWithdrawDao creates and returns a new DAO object for table data access.
func NewWithdrawDao() *WithdrawDao {
	return &WithdrawDao{
		group:   "default",
		table:   "o_withdraw",
		columns: withdrawColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WithdrawDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WithdrawDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WithdrawDao) Columns() WithdrawColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WithdrawDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WithdrawDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WithdrawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
