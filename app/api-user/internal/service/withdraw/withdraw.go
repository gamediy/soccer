package withdraw

import (
	"context"
	"star_net/app/api-user/internal/model"
)

var (
	Service = &withdraw{}
)

type withdraw struct {
}

type Withdraw interface {
	Submit(ctx context.Context, input model.WithdrawSubmitInput) error
	Record(ctx context.Context, input model.WithdrawRecordInput) (*model.WithdrawRecordOutput, error)
	GetWithdrawItem(ctx context.Context) (*model.GetWithdrawItemOutput, error) //获取充值列表
}
