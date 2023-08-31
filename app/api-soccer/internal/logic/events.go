package logic

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type eventsOutput struct {
	Home   string
	Away   string
	League string
}

func GetTeamByLang(ctx context.Context, lang string, eventsId int64) eventsOutput {

	item := entity.Events{}
	dao.Events.Ctx(ctx).Scan(&item, "id", eventsId)
	out := eventsOutput{}
	if item.Id == 0 {
		return out
	}
	home, _ := dao.Team.Ctx(ctx).One("id", item.HomeTeamId)
	if home != nil {
		out.Home = home.Map()[lang+"_name"].(string)
	}
	away, _ := dao.Team.Ctx(ctx).One("id", item.AwayTeamId)
	if away != nil {
		out.Away = away.Map()[lang+"_name"].(string)
	}
	league, _ := dao.League.Ctx(ctx).One("id", item.LeagueId)
	if league != nil {
		out.League = league.Map()[lang+"_name"].(string)
	}
	return out
}
