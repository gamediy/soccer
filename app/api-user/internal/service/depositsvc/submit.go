package depositsvc

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	consts2 "star_net/app/api-user/consts"
	"star_net/app/api-user/internal/service"
	"star_net/consts"
	"star_net/core/push"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xtime"
	"star_net/utility/utils/xuuid"
)

var (
	Submit = &submit{}
)

type DepositSubmitInput struct {
	PayId           int     `json:"payId"`           //充值Id
	Amount          float64 `json:"amount"`          //充值金额
	TransferOrderNo string  `json:"transferOrderNo"` //成功订单号
	TranserImg      string  `json:"transerImg"`      //成功图片

}
type submit struct {
	DepositSubmitInput
}

func (input *submit) Exec(ctx context.Context) error {

	userInfo := service.GetUserInfo(ctx)
	payInfo := entity.AmountItem{}
	dao.AmountItem.Ctx(ctx).Scan(&payInfo, input.PayId)
	if payInfo.Id == 0 || payInfo.Status == 0 {
		return consts2.ErrDepositClosed
	}

	if int64(input.Amount) < payInfo.Min || int64(input.Amount) > payInfo.Max {
		return consts2.ErrDepositIncorrect
	}

	order := entity.Deposit{}
	order.OrderNo = xuuid.GetsnowflakeUUID().Int64()
	order.Uid = int64(userInfo.Uid)
	order.Account = userInfo.Account
	order.Status = consts.DepositStatusPending
	order.Net = payInfo.Net
	order.Protocol = payInfo.Protocol
	order.Currency = payInfo.Currency
	order.Address = payInfo.Address
	order.StatusRemark = userInfo.I18n.T(ctx, "处理中")
	order.Amount = input.Amount
	order.Pid = int64(userInfo.Pid)
	order.ParentPath = userInfo.ParentPath
	order.FinishAt = xtime.Get1970Datetime()
	order.AmountItemId = payInfo.Id
	order.Detail = fmt.Sprintf("%s %s", payInfo.Title, payInfo.Detail)
	order.TransferOrderNo = input.TransferOrderNo
	order.TransferImg = input.TranserImg
	if order.Address == "" {
		if order.Protocol == "trc20" { //在获取充值列表的时候要给用户生成一个充值地址 银行卡 address如果为空可能需要调用支付三方API
			daccount := entity.DigitalAccount{}
			dao.DigitalAccount.Ctx(ctx).Where("uid", userInfo.Uid).Scan(&daccount)
			order.Address = daccount.Address
		}
	}

	insert, err := dao.Deposit.Ctx(ctx).Data(&order).Insert()
	if err != nil {
		return err
	}
	affected, err := insert.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return consts2.ErrTryAgain
	}

	message := push.Message{
		CH: push.MessageItem{
			Title: "用户充值",
			Body:  fmt.Sprintf("%s 充值:%.2f 用户：%s", payInfo.Title, input.Amount, userInfo.Account),
		},
		EN: push.MessageItem{
			Title: "order",
			Body:  fmt.Sprintf("%s order： %.2f account：%s", payInfo.Title, input.Amount, userInfo.Account),
		},
	}
	message.Trigger(push.ChannelAdmin, push.EventDeposit)

	return nil
}
