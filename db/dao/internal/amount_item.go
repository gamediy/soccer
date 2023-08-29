// ==========================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AmountItemDao is the data access object for table u_amount_item.
type AmountItemDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AmountItemColumns // columns contains all the column names of Table for convenient usage.
}

// AmountItemColumns defines and stores column names for table u_amount_item.
type AmountItemColumns struct {
	Id               string //
	Title            string // 标题
	Status           string // 状态
	Detail           string // 详情
	AmountCategoryId string //
	Net              string // tron eth
	Min              string // 最小金额
	Max              string // 最大金额
	Fee              string // 手续费
	Type             string // 类型
	Logo             string // Logo
	Sort             string // 排序大到小
	Country          string // 国家地区
	Currency         string // 货币单位
	Protocol         string // 协议
	CreatedAt        string //
	UpdatedAt        string //
	Address          string // 地址或卡号
}

// amountItemColumns holds the columns for table u_amount_item.
var amountItemColumns = AmountItemColumns{
	Id:               "id",
	Title:            "title",
	Status:           "status",
	Detail:           "detail",
	AmountCategoryId: "amount_category_id",
	Net:              "net",
	Min:              "min",
	Max:              "max",
	Fee:              "fee",
	Type:             "type",
	Logo:             "logo",
	Sort:             "sort",
	Country:          "country",
	Currency:         "currency",
	Protocol:         "protocol",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	Address:          "address",
}

// NewAmountItemDao creates and returns a new DAO object for table data access.
func NewAmountItemDao() *AmountItemDao {
	return &AmountItemDao{
		group:   "default",
		table:   "u_amount_item",
		columns: amountItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AmountItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AmountItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AmountItemDao) Columns() AmountItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AmountItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AmountItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AmountItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
