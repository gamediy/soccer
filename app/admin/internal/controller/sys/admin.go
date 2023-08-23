package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/admin/api/sys"
	model2 "star_net/app/admin/internal/model"
	"star_net/app/admin/internal/service/syssvc"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/pure/add"
	"star_net/pure/del"
	"star_net/pure/get"
	"star_net/pure/put"
	"star_net/utility/utils/xotp"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct {
}

func (c *cAdmin) AddAdmin(ctx context.Context, req *sys.AddAdminReq) (res *sys.AddAdminRes, err error) {
	if err = add.Admin(ctx, req.Admin); err != nil {
		return nil, err
	}

	return
}

func (c *cAdmin) DelAdmin(ctx context.Context, req *sys.DelAdminReq) (res *sys.DelAdminRes, err error) {
	if err = del.Admin(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c *cAdmin) GetAdmin(ctx context.Context, req *sys.GetAdminReq) (res *sys.GetAdminRes, err error) {
	d, err := get.AdminById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return (*sys.GetAdminRes)(d), nil
}

func (c *cAdmin) ListAdmin(ctx context.Context, req *sys.ListAdminReq) (res *sys.ListAdminRes, err error) {
	count, admins, err := syssvc.Admin.ListAdmin(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &sys.ListAdminRes{Total: count, List: admins}, nil
}
func (c *cAdmin) UpdateAdmin(ctx context.Context, req *sys.UpdateAdminReq) (res *sys.UpdateAdminRes, err error) {
	if err = put.Admin(ctx, req.Admin); err != nil {
		return nil, err
	}
	return
}

func (c *cAdmin) UserInfo(ctx context.Context, req *sys.UserInfoReq) (res *sys.UserInfoRes, err error) {
	menu, button, userInfo, err := syssvc.Admin.UserInfo(ctx)
	res = &sys.UserInfoRes{}
	res.Menus = menu
	res.Buttons = button
	res.Info = userInfo
	return
}

func (c *cAdmin) EditPwd(ctx context.Context, req *sys.EditPwdReq) (res *model.CommonRes, err error) {
	info, ok := ctx.Value(consts.UserInfo).(model2.UserInfo)
	if !ok {
		return nil, fmt.Errorf("user info error")
	}
	err = syssvc.Admin.EditPwd(ctx, int64(info.Uid), req.NewPwd)
	return
}

func (c *cAdmin) UserEditPwd(ctx context.Context, req *sys.UserEditPwdReq) (res *model.CommonRes, err error) {
	err = syssvc.Admin.EditPwd(ctx, req.Uid, req.NewPwd)
	return
}

// ValidateOtp 账号登录后二次校验
func (c *cAdmin) ValidateOtp(ctx context.Context, req *sys.ValidateOtpReq) (*sys.ValidateOtpRes, error) {
	uname := req.Uname
	code := req.TotpCode
	//g.RequestFromCtx(ctx).GetCtxVar("uid")
	var userInfo entity.Admin
	err := dao.Admin.Ctx(ctx).Where("uname", uname).Scan(&userInfo)
	if err != nil {
		return &sys.ValidateOtpRes{
			Ok: false,
		}, err
	}
	return &sys.ValidateOtpRes{
		Ok: xotp.ValidateOtp(code, userInfo.KeySecret),
	}, nil
}

func (c *cAdmin) GetOtp(ctx context.Context, req *sys.GetOtpReq) (res *sys.GetOptRes, err error) {
	account := gconv.String(ctx.Value("account"))
	otpKey, err := xotp.GenerateOtp(account, "Disillusioned")
	if err != nil {
		return nil, err
	}
	return &sys.GetOptRes{
		KeySecret: otpKey.Secret(),
		KeyString: otpKey.String(),
	}, nil
}

// SetOtp 首次验证otp，并把keySecret存入表
func (c *cAdmin) SetOtp(ctx context.Context, req *sys.SetOtpReq) (*sys.SetOtpRes, error) {
	code := req.TotpCode
	secret := req.KeySecret
	var flag = xotp.ValidateOtp(code, secret)
	uid := gconv.Int(ctx.Value("uid"))

	//校验通过写入表
	if flag {
		_, err := dao.Admin.Ctx(ctx).Data("key_secret", secret).WherePri(uid).Update()
		if err != nil {
			return &sys.SetOtpRes{
				Ok: false,
			}, err
		}
		return &sys.SetOtpRes{
			Ok: true,
		}, nil
	} else {
		return &sys.SetOtpRes{
			Ok: false,
		}, gerror.New("验证出错")
	}
}

// CheckOtp 账号登录后判断是否需要otp，根据表中是否存在keySecret值来判断
func (c *cAdmin) CheckOtp(ctx context.Context, req *sys.CheckOtpReq) (*sys.CheckOtpRes, error) {
	uname := req.Uname
	var userInfo entity.Admin
	err := dao.Admin.Ctx(ctx).Where("uname", uname).Scan(&userInfo)
	if err != nil {
		return &sys.CheckOtpRes{
			Need2Step: false,
		}, err
	}
	if gconv.Bool(userInfo.KeySecret) {
		return &sys.CheckOtpRes{
			Need2Step: true,
		}, nil
	} else {
		return &sys.CheckOtpRes{
			Need2Step: false,
		}, nil
	}
}
