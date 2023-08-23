package walletsvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type Logs struct {
	Page        int
	Size        int
	BalanceCode int
}

func (s Logs) Exec(ctx context.Context) (int, []*entity.WalletLog, error) {
	var (
		total int
		d     = make([]*entity.WalletLog, 0)
	)
	if s.Page < 0 {
		s.Page = 1
	}
	if s.Size > 100 || s.Size < 0 {
		s.Size = 10
	}

	db := dao.WalletLog.Ctx(ctx)

	if s.BalanceCode != 0 {
		db = db.Where("balance_code", s.BalanceCode)
	}
	if err := db.Where("uid", service.GetUserInfo(ctx).UidInt64).OrderDesc("id").Page(s.Page, s.Size).ScanAndCount(&d, &total, false); err != nil {
		g.Log().Error(ctx, err)
		return 0, nil, err
	}
	return total, d, nil
}
