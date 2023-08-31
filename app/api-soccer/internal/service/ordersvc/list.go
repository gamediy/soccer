package ordersvc

import (
	"context"
	"fmt"
	"star_net/app/api-soccer/internal/logic"
	"star_net/app/api-soccer/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type OrderList struct {
	Type   int `json:"type" dc:"1:未结 2：已结"`
	Status int `json:"status" dc:"0：撤单，1:投注成功，2：中奖，3：未中奖 -1 所有"`
}

type OrderListOutput struct {
	Events       string  `json:"eventsTitle"`
	League       string  `json:"league"`
	Profit       float64 `json:"profit"`
	StatusRemark string  `json:"statusRemark"`
	Odds         float64 `json:"odds" dc:"赔率"`
	Amount       float64 `json:"amount" dc:"下注金额"`
	BoutRemark   string  `json:"boutRemark" dc:"场次"`
}

func (order *OrderList) Exec(ctx context.Context) []OrderListOutput {
	userInfo := service.GetUserInfo(ctx)
	list := []OrderListOutput{}
	if order.Type == 1 {
		orderList := []entity.SoccerOrderSettle{}
		db := dao.SoccerOrderSettle.Ctx(ctx).Where("uid", userInfo.UidInt64)
		if order.Status >= 0 {
			db.Where("status", order.Status).Scan(&orderList)
		} else {
			db.Scan(&orderList)
		}
		for _, settle := range orderList {
			lang := logic.GetTeamByLang(ctx, userInfo.Lang, settle.EventsId)
			list = append(list, OrderListOutput{
				Events:       fmt.Sprintf("%s vs %s", lang.Home, lang.Away),
				League:       lang.League,
				StatusRemark: userInfo.I18n.T(ctx, fmt.Sprintf("订单状态_%d", settle.Status)),
				Profit:       settle.Profit,
				Amount:       settle.Amount,
				Odds:         settle.Odds,
				BoutRemark:   userInfo.I18n.T(ctx, fmt.Sprintf("场次_%d", settle.BoutStatus)),
			})
		}

	} else {

		orderList := []entity.SoccerOrder{}
		db := dao.SoccerOrder.Ctx(ctx).Where("uid", userInfo.UidInt64)
		if order.Status >= 0 {
			db.Where("status", order.Status).Scan(&orderList)
		} else {
			db.Scan(&orderList)
		}
		for _, settle := range orderList {
			lang := logic.GetTeamByLang(ctx, userInfo.Lang, settle.EventsId)
			list = append(list, OrderListOutput{
				Events:       fmt.Sprintf("%s vs %s", lang.Home, lang.Away),
				League:       lang.League,
				StatusRemark: userInfo.I18n.T(ctx, fmt.Sprintf("订单状态_%d", settle.Status)),
				Profit:       settle.Profit,
				Amount:       settle.Amount,
				Odds:         settle.Odds,
				BoutRemark:   userInfo.I18n.T(ctx, fmt.Sprintf("场次_%d", settle.BoutStatus)),
			})
		}

	}
	return list
}
