package event

import (
	"context"
	"star_net/app/api-soccer/api/event"
	"star_net/app/api-soccer/internal/service/eventsvc"
)

func (cEvent) List(ctx context.Context, req *event.EventListReq) (res *event.EventListRes, err error) {

	list := eventsvc.EventsList{}
	list.Status = req.Status
	exec, err := list.Exec(ctx)
	res = &event.EventListRes{}
	res.List = exec
	return res, err
}
