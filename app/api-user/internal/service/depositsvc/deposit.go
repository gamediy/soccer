package depositsvc

import (
	"context"
	"star_net/app/api-user/internal/model"
)

var (
	Service = &deposit{}
)

type deposit struct {
}
type RecordOutput deposit
type Deposit interface {
	Submit(ctx context.Context, input model.DepositSubmitInput) error //提交充值订单
	Record(ctx context.Context, input model.DepositRecordInput) ([]*model.DepositRecordOutput, error)
	GetDepositItem(ctx context.Context) (*model.GetDepositItemOutput, error) //获取充值列表
}
