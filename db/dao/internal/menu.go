// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for table s_menu.
type MenuDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns MenuColumns // columns contains all the column names of Table for convenient usage.
}

// MenuColumns defines and stores column names for table s_menu.
type MenuColumns struct {
	Id         string //
	Pid        string //
	Icon       string //
	BgImg      string //
	Name       string //
	Path       string //
	Sort       string //
	Type       string // 1normal 2menu 3 button
	Desc       string //
	FilePath   string //
	Status     string //
	CreatedAt  string //
	UpdatedAt  string //
	Permission string // 权限标识
}

// menuColumns holds the columns for table s_menu.
var menuColumns = MenuColumns{
	Id:         "id",
	Pid:        "pid",
	Icon:       "icon",
	BgImg:      "bg_img",
	Name:       "name",
	Path:       "path",
	Sort:       "sort",
	Type:       "type",
	Desc:       "desc",
	FilePath:   "file_path",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Permission: "permission",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   "default",
		table:   "s_menu",
		columns: menuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
