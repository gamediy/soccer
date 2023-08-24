package withdrawsvc

import (
	"context"
	"fmt"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xuuid"
)

type BindWithdrawAccount struct {
	BankId  uint64 `v:"required#银行ID不能为空" dc:"银行ID"`
	Address string `v:"required#收款地址不能为空"`
	Title   string `json:"title" dc:"持卡人" v:"required#持卡人不能为空"`
}

func (s BindWithdrawAccount) Exec(ctx context.Context) error {
	var bank entity.Bank
	_ = dao.Bank.Ctx(ctx).Scan(&bank, "id", s.BankId)
	if bank.Id == 0 {
		return fmt.Errorf("银行不存在")
	}
	user := service.GetUserInfo(ctx)
	d := entity.WithdrawAccount{}
	d.Id = xuuid.GetsnowflakeUUID().Int64()
	d.Net = bank.Net
	d.Protocol = bank.Protocol
	d.Uid = int(user.UidInt64)
	d.Account = user.Account
	d.Currency = bank.Currency
	d.Address = s.Address
	d.Status = 1
	d.Default = -1
	d.Title = s.Title
	if _, err := dao.WithdrawAccount.Ctx(ctx).Insert(d); err != nil {
		return err
	}
	return nil
}
