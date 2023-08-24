package withdrawsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
)

type SetDefaultWithdrawAccount struct {
	Id int64
}

func (s SetDefaultWithdrawAccount) Exec(ctx context.Context) error {
	u := service.GetUserInfo(ctx)
	count, err := dao.WithdrawAccount.Ctx(ctx).Count("id = ? and uid = ?", s.Id, u.Uid)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("银行卡不存在")
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err = tx.Model(dao.WithdrawAccount.Table()).Ctx(ctx).Update(g.Map{"default": -1}, "uid", u.Uid); err != nil {
			return err
		}
		if _, err = tx.Model(dao.WithdrawAccount.Table()).Ctx(ctx).Update(g.Map{"default": 1}, "id", s.Id); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
