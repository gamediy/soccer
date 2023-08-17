package model

import "github.com/gogf/gf/v2/frame/g"

type WithdrawSubmitInput struct {
	WithdrawId        int     `json:"withdrawId"`
	Amount            float64 `json:"amount"`
	WithdrawAccountId int     `json:"WithdrawAccount"`
}
type WithdrawRecordInput struct {
}

type WithdrawRecordOutput struct {
}

type GetWithdrawItemOutput struct {
	Tips string
	Item g.Map
}
