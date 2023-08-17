// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Withdraw is the golang structure for table withdraw.
type Withdraw struct {
	OrderNo       int64       `json:"orderNo"       description:"订单号"`
	Account       string      `json:"account"       description:"账号"`
	Uid           int64       `json:"uid"           description:"UID"`
	Pid           int64       `json:"pid"           description:"上级ID"`
	Status        int         `json:"status"        description:"状态 0 提现中，1提现成功，-1提现失败"`
	FinishAt      *gtime.Time `json:"finishAt"      description:"完成时间"`
	Detail        string      `json:"detail"        description:"详情"`
	StatusRemark  string      `json:"statusRemark"  description:"状态说明"`
	Amount        float64     `json:"amount"        description:"提现金额"`
	SysRemark     string      `json:"sysRemark"     description:"系统说明"`
	Address       string      `json:"address"       description:"地址（如果是银行卡，就是账号）"`
	AmountFinally int64       `json:"amountFinally" description:"最后倒账金额"`
	Fee           int64       `json:"fee"           description:"手续费"`
	Net           string      `json:"net"           description:"网络（数字货币和银行卡代号）"`
	AmountItemId  int         `json:"amountItemId"  description:"编号"`
	Currency      string      `json:"currency"      description:"货币单位（USDT PHP CNY）"`
	Protocol      string      `json:"protocol"      description:"协议"`
	Title         string      `json:"title"         description:"标题"`
	Note          string      `json:"note"          description:"说明"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
	AdminOperate  string      `json:"adminOperate"  description:"操作用户"`
	ParentPath    string      `json:"parentPath"    description:""`
}
