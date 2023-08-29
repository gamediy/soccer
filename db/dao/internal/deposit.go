// ==========================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DepositDao is the data access object for table o_deposit.
type DepositDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns DepositColumns // columns contains all the column names of Table for convenient usage.
}

// DepositColumns defines and stores column names for table o_deposit.
type DepositColumns struct {
	OrderNo         string // 订单号
	Account         string // 账号
	Uid             string // UID
	Pid             string // 上级ID
	Status          string // 状态
	FinishAt        string // 完成时间
	Detail          string // 详情
	StatusRemark    string // 状态说明
	Amount          string // 充值金额
	SysRemark       string // 系统说明
	Address         string // 地址或者是卡号
	Net             string // 类型
	AmountItemId    string // 充值的编号
	Currency        string // 货币单位
	Protocol        string // 协议
	ParentPath      string // 上级代理全路经
	CreatedAt       string //
	UpdatedAt       string //
	AdminOperate    string // 操作用户
	TransferOrderNo string // 用户转账订单号
	TransferImg     string // 用户转账图片
}

// depositColumns holds the columns for table o_deposit.
var depositColumns = DepositColumns{
	OrderNo:         "order_no",
	Account:         "account",
	Uid:             "uid",
	Pid:             "pid",
	Status:          "status",
	FinishAt:        "finish_at",
	Detail:          "detail",
	StatusRemark:    "status_remark",
	Amount:          "amount",
	SysRemark:       "sys_remark",
	Address:         "address",
	Net:             "net",
	AmountItemId:    "amount_item_id",
	Currency:        "currency",
	Protocol:        "protocol",
	ParentPath:      "parent_path",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	AdminOperate:    "admin_operate",
	TransferOrderNo: "transfer_order_no",
	TransferImg:     "transfer_img",
}

// NewDepositDao creates and returns a new DAO object for table data access.
func NewDepositDao() *DepositDao {
	return &DepositDao{
		group:   "default",
		table:   "o_deposit",
		columns: depositColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DepositDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DepositDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DepositDao) Columns() DepositColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DepositDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DepositDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DepositDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
