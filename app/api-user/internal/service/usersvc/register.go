package usersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xip"
	"star_net/utility/utils/xpwd"
)

type Register struct {
	Ctx      context.Context
	Account  string `dc:"账号" json:"account"`
	Password string `dc:"密码" json:"password"`
	Xid      string `dc:"邀请码" json:"xid" `
	Phone    string `json:"phone"`
	Email    string `json:"email"`

	Ip          string `json:"-"`
	RealName    string
	ClientAgent string
}

func (s *Register) Exec() (string, error) {

	if err := s.checkAccountExist(s.Ctx, s.Account); err != nil {
		return "", err
	}
	var d entity.User
	loc := xip.GetIpInfoIo(s.Ctx, s.Ip)
	d.Account = s.Account
	d.Status = 1
	d.Ip = s.Ip
	d.City = loc.City
	d.Country = loc.Country
	d.Password = xpwd.GenPwd(s.Password)
	d.Phone = s.Phone
	d.Email = s.Email
	d.RealName = s.RealName
	d.ClientAgent = s.ClientAgent
	parent, err := s.getParent()
	if err != nil {
		return "", err
	}
	if parent != nil {
		d.ParentPath = fmt.Sprintf("%s%d/", parent.ParentPath, parent.Id)
	}
	if err = g.DB().Transaction(s.Ctx, func(ctx context.Context, tx gdb.TX) error {
		uid, err := tx.Model(dao.User.Table()).InsertAndGetId(&d)
		if err != nil {
			return err
		}
		if _, err = tx.Model(dao.Wallet.Table()).Ctx(ctx).Insert(&entity.Wallet{
			Uid:     uid,
			Balance: 0,
			Account: s.Account,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}

	return "", nil
}

func (s *Register) getParent() (*entity.User, error) {
	var d entity.User
	if s.Xid == "" {
		return nil, nil
	}
	one, err := dao.User.Ctx(s.Ctx).One(&d, "xid", s.Xid)
	if one.IsEmpty() {
		return nil, nil
	}
	if err = one.Struct(&d); err != nil {
		return nil, err
	}
	return &d, err
}

func (s *Register) checkAccountExist(ctx context.Context, account string) error {
	count, err := dao.User.Ctx(ctx).Count("account", account)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	return nil
}
