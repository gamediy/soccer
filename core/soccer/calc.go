package soccer

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/core/wallet"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

const (
	OrderStatusCancellation = 0
	OrderStatusBetSuccess   = 1
	OrderStatusWon          = 2
	OrderStatusLost         = 3
)

// 结算
func Calc(ctx context.Context, eventId int64, result OpenResult) error {

	event := entity.Events{}
	dao.Events.Ctx(ctx).Scan(&event, eventId)
	if event.Id == 0 {
		return fmt.Errorf("没有赛事")
	}
	odds := []*entity.EventsOdds{}
	dao.EventsOdds.Ctx(ctx).Where("events_id", eventId).Scan(&odds)
	if len(odds) == 0 {
		return fmt.Errorf("没有赛事赔率")
	}
	wonPlayCode := make([]int, 0)
	lostPlayCode := make([]int, 0)
	for _, odd := range odds {
		err := PlayList[odd.PlayCode].WonCheck(result, odd.CalcRule)
		if err != nil {
			lostPlayCode = append(lostPlayCode, odd.PlayCode)
		} else {
			wonPlayCode = append(wonPlayCode, odd.PlayCode)
		}
	}
	lostOrder := []entity.SoccerOrderSettle{}
	wonOrder := []entity.SoccerOrderSettle{}
	dao.SoccerOrderSettle.Ctx(ctx).Where("events_id=? and play_code in (?) and status=1", eventId, lostPlayCode).Scan(&lostOrder)
	for index, _ := range lostOrder {
		lostOrder[index].Status = OrderStatusLost
		lostOrder[index].CalcAt = gtime.Now()
		lostOrder[index].Profit = -lostOrder[index].Amount

	}
	dao.SoccerOrderSettle.Ctx(ctx).Data(lostOrder).Batch(20).Update()
	for i, item := range wonOrder {
		go func(item entity.SoccerOrderSettle) {
			wonOrder[i].Profit = item.OddsProfitRate*item.Amount - item.Amount
			wonOrder[i].CalcAt = gtime.Now()
			wonOrder[i].Status = OrderStatusWon
			update := wallet.BalanceUpdate{}
			update.Uid = item.Uid
			update.Title = item.EventsTitle
			update.Note = item.OddsTitle
			update.BalanceCode = wallet.Profit
			update.Update(context.Background(), func(ctx context.Context, tx gdb.TX) error {
				_, err := tx.Model(dao.SoccerOrderSettle.Table()).Where("order_no", item.OrderNo).Update(wonOrder[i])
				if err != nil {
					return err
				}
				return err
			})
		}(item)
	}
	return nil
}
