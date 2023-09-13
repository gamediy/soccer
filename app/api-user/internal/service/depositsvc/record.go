package depositsvc

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/utility/utils/xtrans"
)

var (
	Record = &record{}
)

type DepositRecordInput struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	OrderNo string `json:"orderNo"`
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
	if input.OrderNo != "" {
		dao.Deposit.Ctx(ctx).Where("order_no", input.OrderNo).Scan(&list)
		return list, nil
	}
	dao.Deposit.Ctx(ctx).Where("uid", userInfo.Uid).Page(input.Page, input.Size).Scan(&list)
	for index, output := range list {
		list[index].StatusRemark = xtrans.T(userInfo.Lang, output.StatusRemark)
	}
	return list, nil
}
