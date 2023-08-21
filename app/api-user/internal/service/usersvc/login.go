package usersvc

//func login(ctx context.Context, req *vpassport.LoginReq) (res *vpassport.LoginRes, err error) {
//	if !xcaptcha.Store.Verify(req.CaptchaId, req.CaptchaCode, true) {
//		return nil, consts.ErrCaptcha
//	}
//	// 对比密码
//	one, err := dao.User.Ctx(ctx).One("account", req.Account)
//	if err != nil {
//		return nil, err
//	}
//	var user *entity.User
//	err = one.Struct(&user)
//	if err != nil {
//		return nil, err
//	}
//	flag := xpwd.ComparePassword(user.Password, req.Password)
//	if !flag {
//		return nil, consts.ErrLogin
//	}
//	token, err := xjwt.GenToken(req.Account, uint64(user.Id), 0)
//	if err != nil {
//		return nil, err
//	}
//	//var data = &user.LoginRes{
//	//	Token: token,
//	//}
//	return data, nil
//}
