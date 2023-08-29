// ==========================================================================
// TypeCode generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LanguageDao is the data access object for table c_language.
type LanguageDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LanguageColumns // columns contains all the column names of Table for convenient usage.
}

// LanguageColumns defines and stores column names for table c_language.
type LanguageColumns struct {
	Id        string //
	Name      string //
	Desc      string //
	En        string // 英语
	Fr        string // 法语
	Ja        string // 日语
	Hi        string // 印度语
	Zh        string // 中文
	Ko        string // 汉语
	Pt        string // 葡萄牙语
	Es        string // 西班牙语
	CreatedAt string //
	UpdatedAt string //
	Hu        string // 匈牙利语
	Ar        string // 阿拉伯语
}

// languageColumns holds the columns for table c_language.
var languageColumns = LanguageColumns{
	Id:        "id",
	Name:      "name",
	Desc:      "desc",
	En:        "en",
	Fr:        "fr",
	Ja:        "ja",
	Hi:        "hi",
	Zh:        "zh",
	Ko:        "ko",
	Pt:        "pt",
	Es:        "es",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Hu:        "hu",
	Ar:        "ar",
}

// NewLanguageDao creates and returns a new DAO object for table data access.
func NewLanguageDao() *LanguageDao {
	return &LanguageDao{
		group:   "default",
		table:   "c_language",
		columns: languageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LanguageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LanguageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LanguageDao) Columns() LanguageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LanguageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LanguageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LanguageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
