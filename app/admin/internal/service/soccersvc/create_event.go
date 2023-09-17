package soccersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/pure/get"
)

type CreateEvent struct {
	HomeTeamId int         `json:"homeTeamId"       description:"主队Id"`
	AwayTeamId int         `json:"awayTeamId"       description:"客队Id"`
	RestTime   string      `json:"restTime"         description:"进行时间"`
	StartTime  *gtime.Time `json:"startTime"        description:"开始时间" v:"required#开始时间不能为空"`
	LeagueId   int64       `json:"leagueId"         description:"联盟编号"`
	AddTime    *gtime.Time `json:"addTime"          description:"添加时间"`
	IsHot      int         `json:"isHot"            description:"热门"`
	Status     int         `json:"status"           description:"0：未开始 1接受下注，2：停止下注，3 已结算"`
	Handicap   int         `json:"handicap"         description:"让球"`
}

func (s CreateEvent) Exec(ctx context.Context) error {
	var d entity.Events
	if s.HomeTeamId == s.AwayTeamId {
		return fmt.Errorf("主客不能为同一战队")
	}
	homeTeam, err := get.Team(ctx, d.HomeTeamId)
	if err != nil {
		return err
	}
	awarTeam, err := get.Team(ctx, d.AwayTeamId)
	if err != nil {
		return err
	}
	d.HomeTeamName = homeTeam.ZhName
	d.HomeTeamId = fmt.Sprint(homeTeam.Id)
	d.AwayTeamId = awarTeam.Id
	d.AwayTeamName = awarTeam.ZhName
	d.StartTime = s.StartTime
	league, err := get.League(ctx, d.LeagueId)
	if err != nil {
		return err
	}
	d.LeagueName = league.ZhName
	d.LeagueId = league.Id
	d.FirstStatus = 0
	d.StartDate = s.StartTime
	d.SecondStatus = 0
	d.Status = 0
	d.AddTime = gtime.Now()
	d.IsHot = s.IsHot
	if _, err = dao.Events.Ctx(ctx).Insert(d); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}