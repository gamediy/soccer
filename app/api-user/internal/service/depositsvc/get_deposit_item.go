package depositsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/blockchain/tron/client"
)

var (
	GetDepositItem = &getDepositItem{}
)

type getDepositItem struct {
}
type GetDepositItemOutput struct {
	Tips string
	Item []g.Map
}

func (s *getDepositItem) Exec(ctx context.Context) (*GetDepositItemOutput, error) {
	userInfo := service.GetUserInfo(ctx)
	list := []entity.AmountItem{}
	res := GetDepositItemOutput{}
	dao.AmountItem.Ctx(ctx).Scan(&list, "status=? and type=?", 1, "Deposit")
	res.Tips = "支付提示，取配置文件"
	for _, item := range list {
		if item.Protocol == "Trc20" && item.Address == "" {
			account := entity.DigitalAccount{}
			dao.DigitalAccount.Ctx(ctx).Where("status=1 and uid=?", userInfo.Uid).Scan(&account)
			if account.Address == "" {
				key, addr := client.TronGrpcClient.GenerateKey()
				account.Address = addr
				account.PrivateKey = key
				account.Account = userInfo.Account
				account.Uid = userInfo.UidInt64
				account.Status = 1
				account.Net = "TRON"
			}
			item.Address = account.Address
		}
		res.Item = append(res.Item, g.Map{
			"payId":    item.Id,
			"title":    item.Title,
			"protocol": item.Protocol,
			"logo":     item.Logo,
			"currency": item.Currency,
			"address":  item.Address,
			"detail":   fmt.Sprintf("%s %d-%d", item.Detail, item.Min, item.Max),
		})
	}
	return &res, nil
}
