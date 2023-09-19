package soccer

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/core/soccer/play"
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
func Calc(ctx context.Context, eventId int64, result play.OpenResult) error {

	event := entity.Events{}
	dao.Events.Ctx(ctx).Scan(&event, eventId)
	if event.Id == 0 {
		return fmt.Errorf("没有赛事")
	}

	lostOrder := []entity.SoccerOrderSettle{}
	wonOrder := []entity.SoccerOrderSettle{}
	listOrder := []entity.SoccerOrderSettle{}

	dao.SoccerOrderSettle.Ctx(ctx).Where("events_id=? and status=1 and bout_status=?", eventId, result.BoutStatus).Scan(&listOrder)

	for _, order := range listOrder {
		order.Status = OrderStatusLost
		order.CalcAt = gtime.Now()
		order.EventsOpenResult = result.Result
		order.Profit = -order.Amount
		profit := play.List[order.PlayCode].WonCheck(result, play.CalcInfo{
			BetAmount: order.Amount,
			Odds:      order.Odds,
			Rule:      order.OddsCalcRule,
		})
		if profit > 0 {
			order.Status = OrderStatusWon
			order.Profit = profit - order.Amount
			wonOrder = append(wonOrder, order)

		} else {
			lostOrder = append(lostOrder, order)
		}

	}

	if len(lostOrder) > 0 {
		g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := tx.Model(dao.SoccerOrder.Table()).Ctx(ctx).Data(lostOrder).Batch(20).Insert()
			if err != nil {
				return err
			}
			lostOrderNo := make([]int64, 0)
			for index, _ := range lostOrder {
				lostOrderNo = append(lostOrderNo, lostOrder[index].OrderNo)
				tx.Model(dao.EventsOdds).Data("total_profit=total_profit+?", lostOrder[index].Amount).Where("id", lostOrder[index].OddsId).Update()
			}

			_, err = tx.Model(dao.SoccerOrderSettle.Table()).Ctx(ctx).Where("order_no in (?)", lostOrderNo).Batch(40).Delete()
			return err
		})
	}

	for i, item := range wonOrder {
		//go func(item entity.SoccerOrderSettle) {
		update := wallet.BalanceUpdate{}
		update.Uid = item.Uid

		update.Note = item.EventsTitle + " " + item.OddsTitle
		update.Amount = item.Profit + item.Amount
		update.BalanceCode = wallet.Profit
		update.OrderNoRelation = item.OrderNo
		err := update.Update(context.Background(), func(ctx context.Context, tx gdb.TX) error {
			_, err := tx.Model(dao.SoccerOrderSettle.Table()).Where("order_no", item.OrderNo).Delete()
			if err != nil {
				return err
			}
			_, err = tx.Model(dao.SoccerOrder.Table()).Data(wonOrder[i]).Insert()
			if err != nil {
				return err
			}
			_, err = tx.Model(dao.EventsOdds.Table()).Data("total_profit=total_profit-?", wonOrder[i].Profit).Where("id", wonOrder[i].OddsId).Update()
			return err
		})
		fmt.Println(err)
		//}(item)
	}

	return nil
}
