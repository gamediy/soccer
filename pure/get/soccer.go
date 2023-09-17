package get

import (
	"context"
	"fmt"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

func Team(ctx context.Context, id interface{}) (*entity.Team, error) {
	var d entity.Team
	_ = dao.Team.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, fmt.Errorf("队伍不存在")
	}
	return &d, nil
}

func League(ctx context.Context, id interface{}) (*entity.League, error) {
	var d entity.League
	_ = dao.League.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, fmt.Errorf("联盟不存在")
	}
	return &d, nil
}
