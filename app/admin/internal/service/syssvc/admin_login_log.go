package syssvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/model"
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
		Uid: int(s.Uid),
		Ip:  s.Ip,
	})
	return err
}

func ListAdminLoginLog(ctx context.Context, req *sys.ListAdminLoginLogReq) ([]*model.AdminLoginLog, int, error) {
	var (
		d     = make([]*model.AdminLoginLog, 0)
		total int
	)
	db := g.DB().Ctx(ctx).Model(dao.AdminLoginLog.Table() + " t1").LeftJoin("s_admin t2 on t1.uid = t2.id")
	if req.Ip != "" {
		db = db.WhereLike("t1.ip", xstr.Like(req.Ip))
	}
	if req.Uname != "" {
		db = db.WhereLike("t2.uname", xstr.Like(req.Uname))
	}
	err := db.Page(req.Page, req.Size).OrderDesc("t1.id").Fields("t1.*,t2.uname").ScanAndCount(&d, &total, false)
	return d, total, err
}
