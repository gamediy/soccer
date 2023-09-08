package event

import (
	"context"
	"star_net/app/api-soccer/api/event"
	"star_net/app/api-soccer/internal/service/eventsvc"
)

func (cEvent) Play(ctx context.Context, req *event.PlayReq) (res *event.PlayRes, err error) {
	res = &event.PlayRes{}
	play := eventsvc.Play{}

	play.EventId = req.EventId
	exec, err := play.Exec(ctx)
	res.List = exec
	return res, err
}
