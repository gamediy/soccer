package usersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xpwd"
)

type PutLoginPass struct {
	Ctx     context.Context
	Uid     uint64
	OldPass string
	NewPass string
}

func (s *PutLoginPass) Exec() error {
	var u entity.User
	_ = dao.User.Ctx(s.Ctx).Scan(&u, "id", s.Uid)
	if u.Id == 0 {
		return fmt.Errorf("用户不存在")
	}
	if !xpwd.ComparePassword(u.Password, s.OldPass) {
		return fmt.Errorf("旧密码不匹配")
	}
	if _, err := dao.User.Ctx(s.Ctx).Update(g.Map{"password": xpwd.GenPwd(s.NewPass)}, "id", u.Id); err != nil {
		return err
	}
	return nil
}
