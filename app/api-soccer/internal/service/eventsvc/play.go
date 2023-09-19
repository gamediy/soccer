package eventsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/api-soccer/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"strings"
)

// 玩法列表
type Play struct {
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
	Type     string  `json:"type"`
}

func (this *Play) Exec(ctx context.Context) ([]*PlayOddsOutput, error) {

	userInfo := service.GetUserInfo(ctx)

	model := make([]*PlayOddsOutput, 0)
	playType := []entity.PlayType{}
	odds := []entity.EventsOdds{}
	dao.PlayType.Ctx(ctx).Where("status", 1).Scan(&playType)
	dao.EventsOdds.Ctx(ctx).Where("status=? and events_id=?", 1, this.EventId).Scan(&odds)

	all, _ := dao.Play.Ctx(ctx).All()
	boutStatus := [3]int{
		3, 1, 2,
	}
	for _, pt := range playType {

		for _, status := range boutStatus {

			e := PlayOddsOutput{}
			e.Title = fmt.Sprintf("%s_%d", pt.ZhName, status)
			e.Title = userInfo.I18n.T(ctx, e.Title)
			e.Code = pt.Code
			for _, item := range odds {
				if item.BoutStatus == status && item.PlayTypeCode == pt.Code && item.Status == 1 {
					e.Item = append(e.Item, &PlayOddsItem{
						Title:    this.getPlayName(all.List(), item.PlayCode, userInfo.Lang),
						Odds:     item.Odds,
						OddsId:   item.Id,
						PlayCode: item.PlayCode,
						Type:     this.getType(item),
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

func (this *Play) getPlayName(list gdb.List, playCode int, lang string) string {
	for _, item := range list {
		i := gconv.Int(item["code"])
		if i == playCode {
			s := item[fmt.Sprintf("%s_name", lang)].(string)
			return s
		}
	}
	return ""
}

func (this *Play) getType(item entity.EventsOdds) string {

	if item.PlayTypeCode == 400 {
		split := strings.Split(item.CalcRule, " ")
		return split[0]

	}
	return ""
}
