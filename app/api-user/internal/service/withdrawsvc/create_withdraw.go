package withdrawsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	consts2 "star_net/app/api-user/consts"
	"star_net/app/api-user/internal/service"
	"star_net/consts"
	"star_net/core/push"
	"star_net/core/wallet"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xtime"
	"star_net/utility/utils/xtrans"
	"star_net/utility/utils/xuuid"
)

type CreateWithdraw struct {
	AmountItemId      int     `json:"amountItemId" v:"required:"`
	Amount            float64 `json:"amount" v:"required"`
	WithdrawAccountId int     `json:"withdrawAccountId" v:"required"`
	Lang              string
}

func (input *CreateWithdraw) Exec(ctx context.Context) error {
	userInfo := service.GetUserInfo(ctx)
	withdrawInfo := entity.AmountItem{}
	_ = dao.AmountItem.Ctx(ctx).Scan(&withdrawInfo, "id", input.AmountItemId)
	if withdrawInfo.Id == 0 || withdrawInfo.Status == -1 {
		return fmt.Errorf("提现通道已关闭")
	}

	if int64(input.Amount) < withdrawInfo.Min || int64(input.Amount) > withdrawInfo.Max {
		return fmt.Errorf("提现金额不正确")
	}
	withdrawAccount := entity.WithdrawAccount{}
	_ = dao.WithdrawAccount.Ctx(ctx).Scan(&withdrawAccount, "id", input.WithdrawAccountId)
	if withdrawAccount.Id == 0 {
		return fmt.Errorf("请先绑定提现账号")
	}
	order := entity.Withdraw{}
	count, err := dao.Withdraw.Ctx(ctx).Count("uid=? and status=0", userInfo.UidInt64)
	if err != nil {
		return err
	}
	if count > 0 {
		return consts2.ErrRepeatedSubmit
	}
	order.OrderNo = xuuid.GetsnowflakeUUID().Int64()
	order.Uid = int64(userInfo.Uid)
	order.Account = userInfo.Account
	order.Status = consts.DepositStatusPending
	order.Net = withdrawAccount.Net
	order.Protocol = withdrawAccount.Protocol
	order.Currency = withdrawAccount.Currency
	order.Address = withdrawAccount.Address
	order.StatusRemark = xtrans.T(input.Lang, "处理中")
	order.Amount = input.Amount
	order.Pid = int64(userInfo.Pid)
	order.ParentPath = userInfo.ParentPath
	order.FinishAt = xtime.Get1970Datetime()
	order.AmountItemId = withdrawInfo.Id
	order.Detail = fmt.Sprintf("%s %s", withdrawInfo.Title, withdrawInfo.Detail)

	update := wallet.BalanceUpdate{}
	update.Uid = int64(userInfo.Uid)
	update.Amount = order.Amount
	update.Title = "Withdraw"
	update.OrderNoRelation = order.OrderNo
	update.Note = fmt.Sprintf("%s_%s", order.Protocol, order.Address)
	update.BalanceCode = wallet.Withdraw
	return update.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err = tx.Model(dao.Withdraw.Table()).Ctx(ctx).Insert(&order); err != nil {
			return err
		}
		message := push.Message{
			CH: push.MessageItem{
				Title: "用户提现",
				Body:  fmt.Sprintf("%s 提现:%.2f 用户：%s", withdrawInfo.Title, input.Amount, userInfo.Account),
			},
			EN: push.MessageItem{
				Title: "withdraw",
				Body:  fmt.Sprintf("%s ordersvc： %.2f account：%s", withdrawInfo.Title, input.Amount, userInfo.Account),
			},
		}
		if err = message.Trigger(push.ChannelAdmin, push.EventDeposit); err != nil {
			return err
		}
		return err
	})
}
