package sys

import (
	"context"
	"star_net/app/api-user/api/sys"
	"star_net/app/api-user/internal/service/syssvc"
)

func (c sSys) GetDict(ctx context.Context, req *sys.GetDictReq) (_ *sys.GetDictRes, _ error) {
	x := syssvc.GetDict{
		K: req.K,
	}
	v, err := x.GetString(ctx)
	if err != nil {
		return nil, err
	}
	return &sys.GetDictRes{V: v}, nil
}
func (c sSys) GetDictAll(ctx context.Context, _ *sys.GetDictAllReq) (_ sys.GetDictAllRes, _ error) {
	x := syssvc.GetDict{}
	return x.GetAll(ctx)
}
