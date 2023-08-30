// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table u_user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table u_user.
type UserColumns struct {
	Id           string //
	Account      string // 账号
	Email        string // 邮箱
	Password     string // 密码
	Status       string // 1:正常，2：冻结
	Xid          string // short code 邀请 短码
	Ip           string // IP
	ClientAgent  string // 注册clientAgen头
	Phone        string // 电报
	LevelBet     string // 下注的等级
	LevelDeposit string // 充值的等级
	LevelAgent   string // 代理的等级
	Pid          string // 上级ID
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	ParentPath   string // 上级全路经 /1/2/3/
	Country      string // 国家地区
	Lang         string // 用户语言
	AreaCode     string //
	City         string //
	RealName     string // real name
	PayPass      string // 交易密码
}

// userColumns holds the columns for table u_user.
var userColumns = UserColumns{
	Id:           "id",
	Account:      "account",
	Email:        "email",
	Password:     "password",
	Status:       "status",
	Xid:          "xid",
	Ip:           "ip",
	ClientAgent:  "client_agent",
	Phone:        "phone",
	LevelBet:     "level_bet",
	LevelDeposit: "level_deposit",
	LevelAgent:   "level_agent",
	Pid:          "pid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ParentPath:   "parent_path",
	Country:      "country",
	Lang:         "lang",
	AreaCode:     "area_code",
	City:         "city",
	RealName:     "real_name",
	PayPass:      "pay_pass",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "u_user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
