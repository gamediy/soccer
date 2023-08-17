package spassport

import (
	"context"
	vpassport "star_net/app/api-user/api/passport"
	"star_net/app/api-user/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xcaptcha"
	"star_net/utility/utils/xjwt"
	"star_net/utility/utils/xpwd"
)

func (s *passport) Login(ctx context.Context, req *vpassport.LoginReq) (res *vpassport.LoginRes, err error) {
	if !xcaptcha.Store.Verify(req.CaptchaId, req.CaptchaCode, true) {
		return nil, consts.ErrCaptcha
	}
	// 对比密码
	one, err := dao.User.Ctx(ctx).One("account", req.Account)
	if err != nil {
		return nil, err
	}
	var user *entity.User
	err = one.Struct(&user)
	if err != nil {
		return nil, err
	}
	flag := xpwd.ComparePassword(user.Password, req.Password)
	if !flag {
		return nil, consts.ErrLogin
	}
	token, err := xjwt.GenToken(req.Account, uint64(user.Id), 0)
	if err != nil {
		return nil, err
	}
	var data = &vpassport.LoginRes{
		Token: token,
	}
	return data, nil
}
