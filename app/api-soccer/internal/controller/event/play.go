package event

import (
	"context"
	"star_net/app/api-soccer/api/event"
	"star_net/app/api-soccer/internal/service/eventsvc"
)

func (cEvent) Play(ctx context.Context, req *event.PlayReq) (res *event.PlayRes, err error) {
	res = &event.PlayRes{}
	eventsvc.PlayOdds.EventId = req.EventId
	exec, err := eventsvc.PlayOdds.Exec(ctx)
	res.List = exec
	return res, err
}
