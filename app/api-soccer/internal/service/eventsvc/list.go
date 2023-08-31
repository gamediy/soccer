package eventsvc

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/app/api-soccer/internal/service"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

var (
	EventList = eventsList{}
)

type eventsList struct {
	Status int `json:"status"`
}
type EventsListOutput struct {
	EventsId  int64       `json:"eventsId"`
	Home      string      `json:"home"`
	Away      string      `json:"away"`
	RestTime  string      `json:"restTime"`
	Status    int         `json:"status"`
	League    string      `json:"league"`
	StartTime *gtime.Time `json:"startTime"`
	EndTime   *gtime.Time `json:"endTime"`
}

// 取进行中可下注的比赛
func (this *eventsList) Exec(ctx context.Context) ([]EventsListOutput, error) {
	list := []entity.Events{}
	model := dao.Events.Ctx(ctx)
	if this.Status >= 0 {
		model.Where("status", this.Status).Scan(&list)
	} else {
		model.Scan(&list)
	}
	userInfo := service.GetUserInfo(ctx)

	res := []EventsListOutput{}
	for _, item := range list {

		home, _ := dao.Team.Ctx(ctx).One("id", item.HomeTeamId)
		if home == nil {
			continue
		}
		away, _ := dao.Team.Ctx(ctx).One("id", item.AwayTeamId)
		if away == nil {
			continue
		}
		res = append(res, EventsListOutput{
			EventsId:  item.Id,
			Home:      home.Map()[userInfo.Lang+"_name"].(string),
			Away:      away.Map()[userInfo.Lang+"_name"].(string),
			League:    item.LeagueTitle,
			Status:    item.Status,
			RestTime:  item.RestTime,
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
		})
	}
	return res, nil

}
