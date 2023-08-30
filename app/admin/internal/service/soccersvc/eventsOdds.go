package soccersvc

import (
	"context"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type BatchInsert struct {
	Data []*entity.EventsOdds
}

func (s BatchInsert) Exec(ctx context.Context) error {
	for _, i := range s.Data {
		count, err := dao.EventsOdds.Ctx(ctx).Count("bout_status = ? and play_code = ? and play_type_code = ?", i.BoutStatus, i.PlayCode, i.PlayTypeCode)
		if err != nil {
			return err
		}
		if count != 0 {
			continue
		}
		if _, err = dao.EventsOdds.Ctx(ctx).Insert(i); err != nil {
			return err
		}

	}
	return nil
}
