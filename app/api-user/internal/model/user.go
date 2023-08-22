package model

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Id           int64       `json:"id"           description:""`
	Account      string      `json:"account"      description:"账号"`
	Email        string      `json:"email"        description:"邮箱"`
	Status       int         `json:"status"       description:"1:正常，2：冻结"`
	Xid          string      `json:"xid"          description:"short code 邀请 短码"`
	Ip           string      `json:"ip"           description:"IP"`
	Phone        string      `json:"phone"        description:"电话"`
	LevelBet     int         `json:"levelBet"     description:"下注的等级"`
	LevelDeposit int         `json:"levelDeposit" description:"充值的等级"`
	LevelAgent   int         `json:"levelAgent"   description:"代理的等级"`
	Pid          int64       `json:"pid"          description:"上级ID"`
	ParentPath   string      `json:"parentPath"   description:"上级全路经 /1/2/3/"`
	Country      string      `json:"country"      description:"国家地区"`
	Lang         string      `json:"lang"         description:"用户语言"`
	City         string      `json:"city"         description:"城市"`
	RealName     string      `json:"realName"     description:"真是姓名"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`
}

type Wallet struct {
	Uid           int64   `json:"uid"           description:""`
	Balance       int64   `json:"balance"       description:"余额"`
	TotalBet      float64 `json:"totalBet"      description:"总投注"`
	TotalDeposit  float64 `json:"totalDeposit"  description:"总充值"`
	TotalWithdraw float64 `json:"totalWithdraw" description:"总提现"`
	Freeze        int64   `json:"freeze"        description:"冻结"`
	Account       string  `json:"account"       description:"账号"`
	ParentPath    string  `json:"parentPath"    description:"上级路经"`
	TotalProfit   float64 `json:"totalProfit"   description:"总盈利"`
	TotalGift     float64 `json:"totalGift"     description:"总赠送"`
}
