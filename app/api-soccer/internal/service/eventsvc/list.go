package eventsvc

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

// 玩法列表
type list struct {
}

type ListOutput struct {
	entity.PlayType
	Item []*entity.EventsOdds
}

func (*list) Exec(ctx context.Context) ([]*ListOutput, error) {
	model := make([]*ListOutput, 0)
	playType := []entity.PlayType{}
	odds := []entity.EventsOdds{}
	dao.PlayType.Ctx(ctx).Where("status", 1).Scan(&playType)
	dao.EventsOdds.Ctx(ctx).Where("status", 1).Scan(&odds)
	for _, pt := range playType {
		e := ListOutput{}
		e.Id = pt.Id
		e.Name = pt.Name
		e.Code = pt.Code
		e.Item = make([]*entity.EventsOdds, 0)
		for _, odd := range odds {
			if odd.PlayTypeCode == pt.Code {
				e.Item = append(e.Item, &odd)
			}
		}
		model = append(model, &e)
	}

	return model, nil
}
