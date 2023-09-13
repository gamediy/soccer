package walletsvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/app/api-user/internal/service"
	"star_net/db/dao"
	"star_net/utility/utils/xtrans"
)

type Logs struct {
	Page        int
	Size        int
	BalanceCode int
}

type LogsOutput struct {
	OrderNo string `json:"orderNo"         description:"订单号，（有可能是充值提现的订单号）"`

	Title         string      `json:"title"     description:"余额编号"`
	BalanceBefore float64     `json:"balanceBefore"   description:"之前余额"`
	BalanceAfter  float64     `json:"balanceAfter"    description:"之后余额"`
	Note          string      `json:"note"            description:"说明"`
	Amount        float64     `json:"amount"          description:"金额"`
	CreatedAt     *gtime.Time `json:"createdAt"       description:"创建时间"`
}

func (s Logs) Exec(ctx context.Context) (int, []*LogsOutput, error) {
	userInfo := service.GetUserInfo(ctx)
	var (
		total int
		d     = make([]*LogsOutput, 0)
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
	for i, output := range d {
		d[i].Title = xtrans.T(userInfo.Lang, output.Title)
	}
	return total, d, nil
}
