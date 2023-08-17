// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AmountItem is the golang structure for table amount_item.
type AmountItem struct {
	Id               int         `json:"id"               description:""`
	Title            string      `json:"title"            description:"标题"`
	Status           int         `json:"status"           description:"状态"`
	Detail           string      `json:"detail"           description:"详情"`
	AmountCategoryId int         `json:"amountCategoryId" description:""`
	Net              string      `json:"net"              description:"tron eth"`
	Min              int64       `json:"min"              description:"最小金额"`
	Max              int64       `json:"max"              description:"最大金额"`
	Fee              int64       `json:"fee"              description:"手续费"`
	Type             string      `json:"type"             description:"类型"`
	Logo             string      `json:"logo"             description:"Logo"`
	Sort             int         `json:"sort"             description:"排序大到小"`
	Country          string      `json:"country"          description:"国家地区"`
	Currency         string      `json:"currency"         description:"货币单位"`
	Protocol         string      `json:"protocol"         description:"协议"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:""`
	Address          string      `json:"address"          description:"地址或卡号"`
}
