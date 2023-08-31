package commonsvc

import (
	"context"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
)

type ListBanners struct {
}

func (s ListBanners) Exec(ctx context.Context) ([]*entity.Banner, error) {
	var d = make([]*entity.Banner, 0)
	err := dao.Banner.Ctx(ctx).Scan(&d)
	for _, i := range d {
		i.Image = consts.ImgPrefix + i.Image
	}
	return d, err
}
