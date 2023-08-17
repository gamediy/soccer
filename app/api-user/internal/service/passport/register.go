package spassport

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	vpassport "star_net/app/api-user/api/passport"
	"star_net/app/api-user/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xjwt"
	"star_net/utility/utils/xpwd"
	"star_net/utility/utils/xuuid"
)

func (s *passport) Register(ctx context.Context, req *vpassport.RegisterReq) (res *vpassport.RegisterRes, err error) {
	fmt.Println("register")
	password := req.Password
	email := req.Email
	areaCode := req.AreaCode
	phone := req.Phone
	account := req.Account
	xid := req.Xid
	// 通过邀请码xid获取邀请人id
	var parentUser entity.User
	_ = dao.User.Ctx(ctx).Where(g.Map{
		"xid": xid,
	}).Scan(&parentUser)

	// 判断用户是否注册
	one, err := dao.User.Ctx(ctx).One("account", account)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if !one.IsEmpty() {
		return nil, consts.ErrUnameExist
	}
	// 根据ip获取用户所在国家
	byIP, err := getCountryByIP(ctx, req.Ip)
	if err != nil {
		return nil, err
	}

	var token string
	parentPath := ""
	if parentUser.Id > 0 && parentUser.ParentPath != "" {
		parentPath = fmt.Sprintf("%s%d/", parentUser.ParentPath, parentUser.Id)
	} else {
		parentPath = fmt.Sprintf("/%d/", parentUser.Id)
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Model(dao.User.Table()).Ctx(ctx).Data(g.Map{
			"password":     xpwd.GenPwd(password),
			"area_code":    areaCode,
			"phone":        phone,
			"email":        email,
			"account":      account,
			"xid":          "",
			"pid":          parentUser.Id,
			"parent_path":  parentPath,
			"status":       1,
			"ip":           req.Ip,
			"client_agent": req.UserAgent,
			"country":      byIP.Country,
			"city":         byIP.City,
		}).InsertAndGetId()
		if err != nil {
			return err
		}

		_, err = tx.Model(dao.User.Table()).Update(g.Map{
			"xid": xuuid.NumberToBase34(id),
		}, "uid", id)
		if err != nil {
			return err
		}

		initWallet := entity.Wallet{
			Uid:     id,
			Balance: 0,
			Account: account,
		}
		if _, err = tx.Model(dao.Wallet.Table()).Ctx(ctx).Insert(initWallet); err != nil {
			return nil
		}

		token, err = xjwt.GenToken(account, uint64(id), 0)
		if err != nil {
			return err
		}
		return nil // 返回 nil 表示事务正常提交
	})
	if err != nil {
		return nil, err
	}
	var data = &vpassport.RegisterRes{
		Token: token,
	}
	return data, nil
}

// GenUserRandomMemberCode 生成邀请码
func GenUserRandomMemberCode() string {
	return gstr.ToUpper(grand.Str("qwertyuiopasdfghjklzxcvbnm1234567890", 6))
}

type CountryCity struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func getCountryByIP(ctx context.Context, ip string) (*CountryCity, error) {
	var cc CountryCity
	if ip == "" {
		return &cc, nil
	}
	err := g.Client().GetVar(ctx, "https://ipapi.co/"+ip+"/json/").Scan(&cc)
	if err != nil {
		return nil, err
	}
	return &cc, nil
}
