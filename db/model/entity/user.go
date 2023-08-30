// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id           int64       `json:"id"           description:""`
	Account      string      `json:"account"      description:"账号"`
	Email        string      `json:"email"        description:"邮箱"`
	Password     string      `json:"password"     description:"密码"`
	Status       int         `json:"status"       description:"1:正常，2：冻结"`
	Xid          string      `json:"xid"          description:"short code 邀请 短码"`
	Ip           string      `json:"ip"           description:"IP"`
	ClientAgent  string      `json:"clientAgent"  description:"注册clientAgen头"`
	Phone        string      `json:"phone"        description:"电报"`
	LevelBet     int         `json:"levelBet"     description:"下注的等级"`
	LevelDeposit int         `json:"levelDeposit" description:"充值的等级"`
	LevelAgent   int         `json:"levelAgent"   description:"代理的等级"`
	Pid          int64       `json:"pid"          description:"上级ID"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"更新时间"`
	ParentPath   string      `json:"parentPath"   description:"上级全路经 /1/2/3/"`
	Country      string      `json:"country"      description:"国家地区"`
	Lang         string      `json:"lang"         description:"用户语言"`
	AreaCode     string      `json:"areaCode"     description:""`
	City         string      `json:"city"         description:""`
	RealName     string      `json:"realName"     description:"real name"`
	PayPass      string      `json:"payPass"      description:"交易密码"`
}
