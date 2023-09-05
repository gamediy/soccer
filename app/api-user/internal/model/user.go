package model

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Account string `json:"account"      description:"账号"`
	Email   string `json:"email"        description:"邮箱"`

	Xid string `json:"xid"          description:"short code 邀请 短码"`

	Phone string `json:"phone"        description:"电话"`

	RealName      string      `json:"realName"     description:"真实姓名"`
	CreatedAt     *gtime.Time `json:"createdAt"    description:"注册时间"`
	PayPassStatus int         `json:"payPassStatus" dc:"交易密码设置状态 0 未设置 1已设置"`
}

type Wallet struct {
	Balance       int64   `json:"balance"       description:"余额"`
	TotalBet      float64 `json:"totalBet"      description:"总投注"`
	TotalDeposit  float64 `json:"totalDeposit"  description:"总充值"`
	TotalWithdraw float64 `json:"totalWithdraw" description:"总提现"`
	Freeze        int64   `json:"freeze"        description:"冻结"`
	Account       string  `json:"account"       description:"账号"`

	TotalProfit float64 `json:"totalProfit"   description:"总盈利"`
	TotalGift   float64 `json:"totalGift"     description:"总赠送"`
}
