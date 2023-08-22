package usersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xpwd"
)

type SetPayPass struct {
	Pass string
	Uid  int64
}

func (s *SetPayPass) Exec(ctx context.Context) error {
	var u entity.User
	_ = dao.User.Ctx(ctx).Scan(&u, "id", s.Uid)
	if u.Id == 0 {
		return fmt.Errorf("用户不存在")
	}
	if u.PayPass != "" {
		return fmt.Errorf("交易密码已被设置")
	}
	if _, err := dao.User.Ctx(ctx).Update(g.Map{"pay_pass": xpwd.GenPwd(s.Pass)}, "id", s.Uid); err != nil {
		return err
	}
	return nil
}
