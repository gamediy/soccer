// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Deposit is the golang structure for table deposit.
type Deposit struct {
	OrderNo         int64       `json:"orderNo"         description:"订单号"`
	Account         string      `json:"account"         description:"账号"`
	Uid             int64       `json:"uid"             description:"UID"`
	Pid             int64       `json:"pid"             description:"上级ID"`
	Status          int         `json:"status"          description:"状态"`
	FinishAt        *gtime.Time `json:"finishAt"        description:"完成时间"`
	Detail          string      `json:"detail"          description:"详情"`
	StatusRemark    string      `json:"statusRemark"    description:"状态说明"`
	Amount          float64     `json:"amount"          description:"充值金额"`
	SysRemark       string      `json:"sysRemark"       description:"系统说明"`
	Address         string      `json:"address"         description:"地址或者是卡号"`
	Net             string      `json:"net"             description:"类型"`
	AmountItemId    int         `json:"amountItemId"    description:"充值的编号"`
	Currency        string      `json:"currency"        description:"货币单位"`
	Protocol        string      `json:"protocol"        description:"协议"`
	ParentPath      string      `json:"parentPath"      description:"上级代理全路经"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       description:""`
	AdminOperate    string      `json:"adminOperate"    description:"操作用户"`
	TransferOrderNo string      `json:"transferOrderNo" description:"用户转账订单号"`
	TransferImg     string      `json:"transferImg"     description:"用户转账图片"`
}
