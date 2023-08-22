package usersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xpwd"
)

type Login struct {
	Ctx      context.Context
	Account  string
	Password string
	Ip       string
}

func (s *Login) Exec() (*model.UserInfo, error) {
	var user entity.User
	_ = dao.User.Ctx(s.Ctx).Scan(&user, "account", s.Account)
	if user.Id == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	if !xpwd.ComparePassword(user.Password, s.Password) {
		return nil, fmt.Errorf("密码错误")
	}
	var info model.UserInfo
	info.Uid = gconv.Float64(user.Id)
	info.UidInt64 = user.Id
	info.Account = user.Account
	info.ClientIP = s.Ip
	info.Pid = gconv.Float64(user.Pid)
	info.ParentPath = user.ParentPath
	info.Lang = user.Lang

	// insert login log
	insertLoginLog := InertLoginLog{
		Ctx:     s.Ctx,
		Account: s.Account,
		Uid:     uint64(user.Id),
		Ip:      s.Ip,
	}
	if err := insertLoginLog.Exec(); err != nil {
		return nil, err
	}
	return &info, nil
}
