package order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/api-soccer/internal/model"
	"star_net/app/api-soccer/internal/service"
	"star_net/core/wallet"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xuuid"
)

func (order) Bet(ctx context.Context, input model.BetInput) error {
	if input.Amount <= 0 {
		return fmt.Errorf("金额错误")
	}
	odds := entity.EventsOdds{}
	dao.EventsOdds.Ctx(ctx).Scan(&odds, input.OddsId)
	if odds.Status != 1 {
		return fmt.Errorf("没有此玩法")
	}
	event := entity.Events{}
	dao.Events.Ctx(ctx).Scan(&event, odds.EventsId)
	if event.Status != 1 {
		return fmt.Errorf("比赛已关闭")
	}
	userInfo := service.GetUserInfo(ctx)
	update := wallet.BalanceUpdate{}
	update.Uid = userInfo.UidInt64
	update.Title = fmt.Sprintf("%s", odds.Title)
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
		order.OddsProfitRate = odds.ProfitRate
		order.OddsTitle = odds.Title
		order.BoutStatus = odds.BoutStatus
		order.EventsId = event.Id
		order.EventsStartTime = event.StartTime
		order.EventsTitle = fmt.Sprintf("%s vs %s", event.EnHomeTeam, event.EnAwayTeam)
		order.LeagueTitle = event.EnLeagueTitle
		order.LeagueId = event.LeagueId
		order.PlayCode = odds.PlayCode
		_, err := tx.Model(dao.SoccerOrderSettle.Table()).Data(&order).Insert()
		return err

	})

}
