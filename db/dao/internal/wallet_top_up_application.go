// ==========================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WalletTopUpApplicationDao is the data access object for table u_wallet_top_up_application.
type WalletTopUpApplicationDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns WalletTopUpApplicationColumns // columns contains all the column names of Table for convenient usage.
}

// WalletTopUpApplicationColumns defines and stores column names for table u_wallet_top_up_application.
type WalletTopUpApplicationColumns struct {
	Id             string //
	TransId        string //
	Uid            string //
	ChangeType     string //
	Money          string //
	ReceiptAmount  string //
	Fee            string //
	Ip             string //
	Description    string //
	Aid            string //
	Status         string // 1 wait 2 success 3 fail
	CreatedAt      string //
	UpdatedAt      string //
	PayAddress     string //
	ReceiveAddress string //
	Hash           string // tx_id
	Icon           string //
}

// walletTopUpApplicationColumns holds the columns for table u_wallet_top_up_application.
var walletTopUpApplicationColumns = WalletTopUpApplicationColumns{
	Id:             "id",
	TransId:        "trans_id",
	Uid:            "uid",
	ChangeType:     "change_type",
	Money:          "money",
	ReceiptAmount:  "receipt_amount",
	Fee:            "fee",
	Ip:             "ip",
	Description:    "description",
	Aid:            "aid",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	PayAddress:     "pay_address",
	ReceiveAddress: "receive_address",
	Hash:           "hash",
	Icon:           "icon",
}

// NewWalletTopUpApplicationDao creates and returns a new DAO object for table data access.
func NewWalletTopUpApplicationDao() *WalletTopUpApplicationDao {
	return &WalletTopUpApplicationDao{
		group:   "default",
		table:   "u_wallet_top_up_application",
		columns: walletTopUpApplicationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WalletTopUpApplicationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WalletTopUpApplicationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WalletTopUpApplicationDao) Columns() WalletTopUpApplicationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WalletTopUpApplicationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WalletTopUpApplicationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WalletTopUpApplicationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
