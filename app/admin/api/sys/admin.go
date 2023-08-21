package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	model2 "star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"star_net/model"
)

type ListAdminReq struct {
	g.Meta `tags:"/sys/admin" method:"get" path:"/admin/list" dc:"管理员列表"`
	model.CommonPageReq
}
type ListAdminRes struct {
	Total int             `json:"total"`
	List  []*entity.Admin `json:"list"`
}

type AddAdminReq struct {
	g.Meta `tags:"/sys/admin" method:"post" path:"/admin" dc:"管理员添加 (waiting)"`
	*entity.Admin
}
type AddAdminRes model.CommonRes

type GetAdminReq struct {
	g.Meta `tags:"/sys/admin" method:"get" path:"/admin" dc:"管理员查询单个"`
	Id     uint64 `json:"id"`
}
type GetAdminRes entity.Admin

type UpdateAdminReq struct {
	g.Meta `tags:"/sys/admin" method:"put" path:"/admin" dc:"管理员修改单个数据"`
	*entity.Admin
}
type UpdateAdminRes model.CommonRes

type DelAdminReq struct {
	g.Meta `tags:"/sys/admin" method:"delete" path:"/admin" dc:"管理员删除单个数据"`
	Id     uint64 `json:"id"`
}

type DelAdminRes model.CommonRes

type UserInfoReq struct {
	g.Meta `tags:"/sys/admin" method:"get" path:"/admin/user_info" dc:"用户信息"`
}
type UserInfoRes struct {
	Menus   []*model2.Menu  `json:"menus"`
	Buttons []*model2.Menu  `json:"buttons"`
	Info    model2.UserInfo `json:"info"`
}

type EditPwdReq struct {
	g.Meta `tags:"/sys/admin" method:"put" path:"/admin/edit_pwd" dc:"管理员修该自己密码"`
	NewPwd string `v:"required" json:"newPwd"`
}

type UserEditPwdReq struct {
	g.Meta `tags:"/sys/admin" method:"put" path:"/admin/edit_pwd_user" dc:"管理员修该用户密码"`
	NewPwd string `v:"required" json:"newPwd"`
	Uid    int64
}

type ValidateOtpReq struct {
	g.Meta   `tags:"/sys/admin" method:"post" path:"/admin/otp/validate" dc:"校验谷歌验证码"`
	TotpCode string `json:"totpCode"`
	Uname    string `json:"uname"`
}

type ValidateOtpRes struct {
	Ok bool `json:"ok"`
}

type SetOtpReq struct {
	g.Meta    `tags:"/sys/admin" method:"post" path:"/admin/otp/set" dc:"设置谷歌验证码"`
	TotpCode  string `json:"totpCode"`
	KeySecret string `json:"keySecret"`
}

type SetOtpRes struct {
	Ok bool `json:"ok"`
}

type CheckOtpReq struct {
	g.Meta `tags:"/sys/admin" method:"get" path:"/admin/otp/check" dc:"判断是否需要otp"`
	Uname  string `json:"uname"`
}

type CheckOtpRes struct {
	Need2Step bool `json:"need2Step"`
}

type GetOtpReq struct {
	g.Meta `tags:"/sys/admin" method:"get" path:"/admin/otp" dc:"获取otp设置密钥"`
}

type GetOptRes struct {
	KeySecret string `json:"keySecret"`
	KeyString string `json:"keyString"`
}
