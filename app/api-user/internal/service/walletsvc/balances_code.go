package walletsvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xtrans"
	"time"
)

type ListBalanceCodes struct {
	Lang string
}

func (s ListBalanceCodes) Exec(ctx context.Context) ([]*entity.BalanceCode, error) {
	var d = make([]*entity.BalanceCode, 0)
	all, err := gcache.GetOrSetFunc(ctx, "", func(ctx context.Context) (value interface{}, err error) {
		return dao.BalanceCode.Ctx(ctx).OrderDesc("code").All()
	}, time.Minute)
	if err != nil {
		return nil, err
	}
	_ = all.Scan(&d)
	for _, i := range d {
		t := xtrans.T(s.Lang, fmt.Sprintf("%s", i.Title))
		i.Title = t
		i.Remark = t

	}
	return d, nil
}
