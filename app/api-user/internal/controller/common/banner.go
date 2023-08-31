package common

import (
	"context"
	"star_net/app/api-user/api/common"
	"star_net/app/api-user/internal/service/commonsvc"
)

func (c cCommon) Banners(ctx context.Context, _ *common.BannersReq) (_ common.BannersRes, _ error) {
	x := commonsvc.ListBanners{}
	return x.Exec(ctx)
}
