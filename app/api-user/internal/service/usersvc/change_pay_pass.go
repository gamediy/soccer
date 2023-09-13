package usersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xpwd"
)

type ChangePayPass struct {
	PayPass string `json:"newPass" v:"required|integer|size:6#交易密码必填|纯数字|6位"`
	Pass    string `json:"pass" v:"required#登录密码不能为空"`
	Address string `json:"address" v:"required#卡号不能为空"`
	Title   string `json:"title" v:"required#持卡人不能为空"`
}

func (s ChangePayPass) Exec(ctx context.Context) error {
	var withdrawAccount entity.WithdrawAccount
	_ = dao.WithdrawAccount.Ctx(ctx).Scan(&withdrawAccount, "title = ? and address = ?", s.Title, s.Address)
	if withdrawAccount.Id == 0 {
		return fmt.Errorf("银行卡不存在")
	}
	var u entity.User
	_ = dao.User.Ctx(ctx).Scan(&u, "id", withdrawAccount.Uid)
	if u.Id == 0 {
		return fmt.Errorf("用户不存在")
	}
	if u.PayPass == "" {
		return fmt.Errorf("请先设置交易密码")
	}
	if !xpwd.ComparePassword(u.Password, s.Pass) {
		return fmt.Errorf("密码错误")
	}
	if _, err := dao.User.Ctx(ctx).Update(g.Map{"pay_pass": xpwd.GenPwd(s.PayPass)}, "id", u.Id); err != nil {
		return err
	}
	return nil
}
