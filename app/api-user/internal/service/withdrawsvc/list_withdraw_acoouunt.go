package withdrawsvc

import (
	"context"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type WithdrawAccountList struct {
	Page, Size int
}

func (s WithdrawAccountList) Exec(ctx context.Context) (int, []*entity.WithdrawAccount, error) {
	var (
		u     = service.GetUserInfo(ctx)
		d     = make([]*entity.WithdrawAccount, 0)
		total int
	)
	if s.Size > 100 {
		s.Size = 10
	}
	if err := dao.WithdrawAccount.Ctx(ctx).Where("uid", u.Uid).OrderDesc("id").Page(s.Page, s.Size).ScanAndCount(&d, &total, false); err != nil {
		return 0, nil, err
	}
	return total, d, nil
}
