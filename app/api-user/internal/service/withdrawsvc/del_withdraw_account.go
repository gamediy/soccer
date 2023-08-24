package withdrawsvc

import (
	"context"
	"fmt"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xpwd"
)

type DelWithdrawAccount struct {
	Id      int64
	PayPass string
}

func (d DelWithdrawAccount) Exec(ctx context.Context) error {
	var u entity.User
	userInfo := service.GetUserInfo(ctx)
	_ = dao.User.Ctx(ctx).Scan(&u, "id", userInfo.Uid)
	if u.Id == 0 {
		return fmt.Errorf("用户数据不存在")
	}
	if u.PayPass == "" {
		return fmt.Errorf("请先设置交易密码")
	}
	if !xpwd.ComparePassword(u.PayPass, d.PayPass) {
		return fmt.Errorf("交易密码错误")
	}
	count, _ := dao.WithdrawAccount.Ctx(ctx).Count("id = ? and uid = ?", d.Id, userInfo.UidInt64)
	if count == 0 {
		return fmt.Errorf("已删除")
	}
	_, _ = dao.WithdrawAccount.Ctx(ctx).Delete("id", d.Id)
	return nil
}
