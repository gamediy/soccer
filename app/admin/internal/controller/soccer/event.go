package soccer

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"regexp"
	"star_net/app/admin/api/soccer"
	"star_net/app/admin/internal/service/soccersvc"
	soccer2 "star_net/core/soccer"
	"star_net/core/soccer/play"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xcrud"
	"star_net/utility/utils/xplay"
)

var (
	Events = cEvents{}
)

type cEvents struct{}

func (c cEvents) Create(ctx context.Context, req *soccer.CreateEventsReq) (_ *soccer.CreateEventsRes, _ error) {
	x := soccersvc.CreateEvent{}
	x.HomeTeamId = req.HomeTeamId
	x.AwayTeamId = req.AwayTeamId
	x.RestTime = ""
	x.LeagueId = req.LeagueId
	x.IsHot = req.IsHot
	x.StartTime = req.StartTime
	x.EndTime = req.EndTime
	if err := x.Exec(ctx); err != nil {
		return nil, err
	}
	return
}

func (c cEvents) Update(ctx context.Context, req *soccer.UpdateEventsReq) (_ *soccer.UpdateEventsRes, _ error) {
	x := xcrud.Update{Ctx: ctx, Table: "p_events", Data: req}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}
func (c cEvents) Read(ctx context.Context, req *soccer.ReadEventsReq) (_ *soccer.ReadEventsRes, _ error) {
	var d soccer.ReadEventsRes
	x := xcrud.Read{Ctx: ctx, Table: "p_events", V: req.Id}
	if err := x.Exec(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
func (c cEvents) ReadList(ctx context.Context, req *soccer.ReadListEventsReq) (_ *soccer.ReadListEventsRes, _ error) {
	var (
		d     = make([]*entity.Events, 0)
		total int
	)
	x := xcrud.ReadList{Ctx: ctx, Table: "p_events", Page: req.Page, Size: req.Size}
	if err := x.Exec(&d, &total); err != nil {
		return nil, err
	}
	return &soccer.ReadListEventsRes{List: d, Total: total}, nil
}
func (c cEvents) Del(ctx context.Context, req *soccer.DelEventsReq) (_ *soccer.DelEventsRes, _ error) {
	x := xcrud.Del{Ctx: ctx, Table: "p_events", V: req.Id}
	if err := x.Exec(); err != nil {
		return nil, err
	}
	return
}

func (c cEvents) OpenResult(ctx gctx.Ctx, req *soccer.OpenResultReq) (res *model.CommonRes, err error) {
	event := entity.Events{}
	dao.Events.Ctx(ctx).Scan(&event, "id", req.EventsId)
	if event.Id == 0 {
		return res, fmt.Errorf("没有赛事")
	}
	re := regexp.MustCompile(`(\d+)-(\d+)`)
	matchString := re.MatchString(req.Result)
	if !matchString {
		return nil, fmt.Errorf("开奖格式错误，正确格式：1-2")
	}
	if req.BoutStatus == 1 {
		event.FirstOpenTime = gtime.Now()
		event.FirstOpenResult = req.Result
		event.FirstStatus = 2
		event.SecondStatus = 1
		soccer2.Calc(ctx, req.EventsId, play.OpenResult{
			Result:     req.Result,
			BoutStatus: req.BoutStatus,
		})
		_, err := dao.Events.Ctx(ctx).Data(&event).Where("id", req.EventsId).Update()
		if err != nil {
			return nil, err
		}

	} else if req.BoutStatus == 2 {

		event.SecondOpenTime = gtime.Now()
		event.SecondOpenResult = req.Result
		event.SecondStatus = 2
		event.Status = 3

		soccer2.Calc(ctx, req.EventsId, play.OpenResult{
			Result:     req.Result,
			BoutStatus: req.BoutStatus,
		})

		toTwo, f2, err := xplay.OpenResutToTwo(event.FirstOpenResult)
		if err != nil {
			return nil, err
		}
		two, f, err := xplay.OpenResutToTwo(req.Result)
		if err != nil {
			return nil, err
		}
		soccer2.Calc(ctx, req.EventsId, play.OpenResult{
			Result:     fmt.Sprintf("%d-%d", gconv.Int(toTwo+two), gconv.Int(f2+f)),
			BoutStatus: 3,
		})
		dao.Events.Ctx(ctx).Data(&event).Where("id", req.EventsId).Update()
	}
	return res, err
}
