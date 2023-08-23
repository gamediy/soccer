package walletsvc

import (
	"context"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type Bank struct{}

func (s Bank) GetBanks(ctx context.Context) []*entity.Bank {
	var d = make([]*entity.Bank, 0)
	_ = dao.Bank.Ctx(ctx).Order("id").Scan(&d, "status = 1")
	for _, i := range d {
		i.Icon = consts.ImgPrefix + i.Icon
	}
	return d
}
