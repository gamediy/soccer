package usersvc

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"star_net/app/admin/api/user"
	adminModel "star_net/app/admin/internal/model"
	"star_net/db/dao"
)

type ReadList struct {
	Ctx context.Context
	Req *user.ReadListUserReq
}

func (s *ReadList) Exec() ([]*adminModel.User, int, error) {
	var (
		d = make([]*adminModel.User, 0)
	)
	db := dao.User.Ctx(s.Ctx)
	if err := db.Offset(s.Req.Page*s.Req.Size).Limit(s.Req.Size).OrderDesc("id").ScanList(&d, "User"); err != nil {
		return nil, 0, err
	}
	count, err := db.Count()
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return d, count, nil
	}
	if err = dao.Wallet.Ctx(s.Ctx).Where("uid", gdb.ListItemValuesUnique(&d, "User", "OrderNo")).
		ScanList(&d, "Wallet", "User", "uid:id"); err != nil {
		return nil, 0, err
	}
	return d, count, nil
}
