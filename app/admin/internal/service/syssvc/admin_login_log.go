package syssvc

import (
	"context"
	"star_net/app/admin/api/sys"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xstr"
)

type LoginLog struct {
	Ctx     context.Context
	Account string
	Uid     uint64
	Ip      string
}

func (s *LoginLog) Save() error {
	_, err := dao.AdminLoginLog.Ctx(s.Ctx).Insert(entity.AdminLoginLog{
		Uid:     int(s.Uid),
		Ip:      s.Ip,
		Account: s.Account,
	})
	return err
}

func ListAdminLoginLog(ctx context.Context, req *sys.ListAdminLoginLogReq) ([]*entity.AdminLoginLog, int, error) {
	var (
		d     = make([]*entity.AdminLoginLog, 0)
		total int
	)
	db := dao.AdminLoginLog.Ctx(ctx)
	if req.Ip != "" {
		db = db.WhereLike("ip", xstr.Like(req.Ip))
	}
	if req.Account != "" {
		db = db.WhereLike("account", xstr.Like(req.Account))
	}
	err := db.Page(req.Page, req.Size).OrderDesc("id").ScanAndCount(&d, &total, false)
	return d, total, err
}
