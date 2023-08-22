package depositsvc

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
)

var (
	Record = &record{}
)

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

type record struct {
	DepositRecordInput
}

func (input *record) Exec(ctx context.Context) ([]*DepositRecordOutput, error) {

	userInfo := service.GetUserInfo(ctx)
	list := []*DepositRecordOutput{}
	dao.Deposit.Ctx(ctx).Where("uid", userInfo.Uid).Offset(input.Size * input.Page).Limit(input.Page).Scan(&list)
	return list, nil
}
