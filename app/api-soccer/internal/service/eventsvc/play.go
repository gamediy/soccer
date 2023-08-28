package eventsvc

import (
	"context"
	"fmt"
	"star_net/app/api-soccer/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

var (
	PlayOdds = play{}
)

// 玩法列表
type play struct {
	EventId int
}

type PlayOddsOutput struct {
	Title string          `json:"title"`
	Code  int             `json:"code"`
	Item  []*PlayOddsItem `json:"item"`
}
type PlayOddsItem struct {
	Title    string  `json:"title"`
	OddsId   int64   `json:"oddsId"`
	Odds     float64 `json:"odds"`
	PlayCode int     `json:"playCode"`
}

func (this *play) Exec(ctx context.Context) ([]*PlayOddsOutput, error) {

	userInfo := service.GetUserInfo(ctx)
	model := make([]*PlayOddsOutput, 0)
	playType := []entity.PlayType{}
	odds := []entity.EventsOdds{}
	dao.PlayType.Ctx(ctx).Where("status", 1).Scan(&playType)
	dao.EventsOdds.Ctx(ctx).Where("status=? and events_id=?", 1, this.EventId).Scan(&odds)
	boutStatus := [3]int{
		3, 1, 2,
	}
	for _, pt := range playType {
		for _, status := range boutStatus {
			e := PlayOddsOutput{}
			e.Title = fmt.Sprintf("%s_%d", pt.Name, status)
			e.Title = userInfo.I18n.T(ctx, e.Title)
			e.Code = pt.Code
			for _, item := range odds {
				if item.BoutStatus == status && item.PlayTypeCode == pt.Code {

					e.Item = append(e.Item, &PlayOddsItem{
						Title:    item.Title,
						Odds:     item.Odds,
						OddsId:   item.Id,
						PlayCode: item.PlayCode,
					})

				}
			}
			if len(e.Item) > 0 {
				model = append(model, &e)
			}

		}

	}

	return model, nil
}
