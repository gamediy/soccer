package ordersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/api-soccer/internal/service"
	"star_net/core/soccer"
	"star_net/core/wallet"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xtime"
	"star_net/utility/utils/xuuid"
)

type BetInput struct {
	OddsId int64   `json:"oddsId"`
	Amount float64 `json:"amount"`
}
type BetOutput struct {
}
type Bet struct {
	BetInput
}

func (input *Bet) Exec(ctx context.Context) error {
	if input.Amount <= 0 {
		return fmt.Errorf("金额错误")
	}
	odds := entity.EventsOdds{}
	dao.EventsOdds.Ctx(ctx).Scan(&odds, "id", input.OddsId)
	if odds.Status != 1 {
		return fmt.Errorf("没有此玩法")
	}
	event := entity.Events{}
	dao.Events.Ctx(ctx).Scan(&event, odds.EventsId)
	if event.Status != 1 {
		return fmt.Errorf("比赛已关闭 ", gconv.String(odds.EventsId))
	}
	userInfo := service.GetUserInfo(ctx)
	update := wallet.BalanceUpdate{}
	update.Uid = userInfo.UidInt64

	update.BalanceCode = wallet.Bet
	update.Amount = input.Amount
	update.OrderNoRelation = xuuid.GetsnowflakeUUID().Int64()
	return update.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
		order := entity.SoccerOrderSettle{}
		order.OrderNo = update.OrderNoRelation
		order.Uid = userInfo.UidInt64
		order.Account = userInfo.Account
		order.Pid = int64(userInfo.Pid)
		order.ParentPath = userInfo.ParentPath
		order.OddsId = odds.Id
		order.OddsCalcRule = odds.CalcRule
		order.Odds = odds.Odds
		order.OddsTitle = odds.Title
		order.BoutStatus = odds.BoutStatus
		order.EventsId = event.Id
		order.Status = soccer.OrderStatusBetSuccess
		order.EventsStartTime = event.StartTime
		order.CalcAt = xtime.Get1970Datetime()
		order.EventsTitle = fmt.Sprintf("%s vs %s", event.HomeTeamName, event.AwayTeamName)
		order.LeagueTitle = event.LeagueName
		order.LeagueId = event.LeagueId
		order.PlayCode = odds.PlayCode
		order.Amount = input.Amount
		_, err := tx.Model(dao.SoccerOrderSettle.Table()).Data(&order).Insert()
		if err != nil {
			return err
		}
		odds.TotalAmount += input.Amount
		_, err = tx.Model(dao.EventsOdds.Table()).Update(&odds, "id", odds.Id)
		return err

	})

}
