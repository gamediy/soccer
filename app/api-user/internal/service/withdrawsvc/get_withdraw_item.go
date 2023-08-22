package withdrawsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

var (
	WithdrawItem = &withdrawItem{}
)

type GetWithdrawItemOutput struct {
	Tips string
	Item g.Map
}

type withdrawItem struct {
}

func (s *withdrawItem) Exec(ctx context.Context) (*GetWithdrawItemOutput, error) {
	userInfo := service.GetUserInfo(ctx)
	wc := []entity.WithdrawAccount{}
	dao.WithdrawAccount.Ctx(ctx).Where("uid=? and status=1", userInfo.UidInt64).Scan(&wc)
	list := []entity.AmountItem{}
	res := GetWithdrawItemOutput{}
	dao.AmountItem.Ctx(ctx).Scan(&list, "status=? and type=?", 1, "Withdraw")
	res.Tips = "提现提示，取配置文件"
	for _, item := range list {

		bind := g.Map{}
		for _, account := range wc {
			if account.Protocol == item.Protocol {
				bind["address"] = account.Address
				bind["protocol"] = account.Protocol
			}
		}
		res.Item = g.Map{
			"payId":    item.Id,
			"title":    item.Title,
			"protocol": item.Protocol,
			"logo":     item.Logo,
			"currency": item.Currency,
			"address":  item.Address,
			"detail":   fmt.Sprintf("%s %d-%d", item.Detail, item.Min, item.Max),
			"bind":     bind,
		}
	}
	return &res, nil
}
