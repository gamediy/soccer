package sys

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/service/syssvc"
)

var (
	AdminLoginLog = cAdminLoginLog{}
)

type cAdminLoginLog struct{}

func (c cAdminLoginLog) List(ctx context.Context, req *sys.ListAdminLoginLogReq) (res *sys.ListAdminLoginLogRes, err error) {
	list, total, err := syssvc.ListAdminLoginLog(ctx, req)
	if err != nil {
		return nil, err
	}
	return &sys.ListAdminLoginLogRes{
		Total: total,
		List:  list,
	}, nil
}
