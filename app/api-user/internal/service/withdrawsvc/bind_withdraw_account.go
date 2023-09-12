package withdrawsvc

import (
	"context"
	"fmt"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xpwd"
	"star_net/utility/utils/xuuid"
)

type BindWithdrawAccount struct {
	BankId  uint64 `v:"required#银行ID不能为空" dc:"银行ID"`
	Address string `v:"required#收款地址不能为空"`
	Title   string `json:"title" dc:"持卡人" v:"required#持卡人不能为空"`
	Pass    string `json:"pass"`
}

func (s BindWithdrawAccount) Exec(ctx context.Context) error {
	var bank entity.Bank
	_ = dao.Bank.Ctx(ctx).Scan(&bank, "id", s.BankId)
	if bank.Id == 0 {
		return fmt.Errorf("银行不存在")
	}
	user := service.GetUserInfo(ctx)
	if err := s.CheckPayPass(ctx, s.Pass, user.UidInt64); err != nil {
		return err
	}
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

func (s BindWithdrawAccount) CheckPayPass(ctx context.Context, pass string, uid int64) error {
	var u entity.User
	_ = dao.User.Ctx(ctx).Scan(&u, "id", uid)
	if u.Id == 0 {
		return fmt.Errorf("用户不存在")
	}
	if u.PayPass == "" {
		return fmt.Errorf("请先设置交易密码")
	}
	if !xpwd.ComparePassword(u.PayPass, pass) {
		return fmt.Errorf("交易密码错误")
	}
	return nil
}
