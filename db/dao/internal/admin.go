// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminDao is the data access object for table s_admin.
type AdminDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns AdminColumns // columns contains all the column names of Table for convenient usage.
}

// AdminColumns defines and stores column names for table s_admin.
type AdminColumns struct {
	Id        string //
	Rid       string // 角色ID
	Uname     string //
	Pwd       string //
	Nickname  string //
	Email     string //
	Phone     string //
	Status    string //
	CreatedAt string //
	UpdatedAt string //
	KeySecret string //
}

// adminColumns holds the columns for table s_admin.
var adminColumns = AdminColumns{
	Id:        "id",
	Rid:       "rid",
	Uname:     "uname",
	Pwd:       "pwd",
	Nickname:  "nickname",
	Email:     "email",
	Phone:     "phone",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	KeySecret: "key_secret",
}

// NewAdminDao creates and returns a new DAO object for table data access.
func NewAdminDao() *AdminDao {
	return &AdminDao{
		group:   "default",
		table:   "s_admin",
		columns: adminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminDao) Columns() AdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
