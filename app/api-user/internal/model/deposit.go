package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DepositSubmitInput struct {
	PayId           int     `json:"payId"`           //充值Id
	Amount          float64 `json:"amount"`          //充值金额
	TransferOrderNo string  `json:"transferOrderNo"` //成功订单号
	TranserImg      string  `json:"transerImg"`      //成功图片

}

// DepositRecordInput 用户充值记录请求参数
type DepositRecordInput struct {
	Page int
	Size int
}

// DepositRecordOutput 用户充值记录输出参数
type DepositRecordOutput struct {
	OrderNo      string     `json:"orderNo"`
	Status       string     `json:"status"`
	StatusRemark string     `json:"statusRemark"`
	Amount       float64    `json:"amount"`
	Note         string     `json:"note"`
	CreatedAt    gtime.Time `json:"createdAt"`
}

type GetDepositItemInput struct {
}

type GetDepositItemOutput struct {
	Tips string
	Item g.Map
}
