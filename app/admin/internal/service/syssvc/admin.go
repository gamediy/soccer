package syssvc

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"star_net/app/admin/internal/logic/sysLogic"
	"star_net/app/admin/internal/model"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/pure/get"
	"star_net/utility/utils/xpwd"
	"time"
)

var (
	Admin = sAdmin{}
)

type sAdmin struct {
}

func (sAdmin) ListAdmin(ctx context.Context, page, size int) (int, []*entity.Admin, error) {
	return sysLogic.ListAdmin(ctx, page, size)
}

func (sAdmin) UserInfo(ctx context.Context) (menu []*model.Menu, button []*model.Menu, userInfo model.UserInfo, err error) {
	userInfo = ctx.Value("userInfo").(model.UserInfo)
	menu = []*model.Menu{}
	button = []*model.Menu{}
	if userInfo.RuleName != "admin" {
		menu, err = Menu.ListMenuByRoleId(ctx, userInfo.RuleId)
		button, err = Menu.ListButtonByRoleId(ctx, userInfo.RuleId)
	} else {
		menu, err = Menu.List(ctx)
		err := dao.Menu.Ctx(ctx).Where("type", 3).Scan(&button)
		if err != nil {
			return menu, button, userInfo, err
		}
	}

	return
}

func (sAdmin) EditPwd(ctx context.Context, uid int64, newPass string) error {

	pwd := xpwd.GenPwd(newPass)
	_, err := dao.Admin.Ctx(ctx).Data("pwd", pwd).Where("id", uid).Update()
	return err
}

func (sAdmin) EditAccount() {

}

func Login(r *ghttp.Request) (string, interface{}) {
	var (
		ctx   = r.Context()
		uname = r.Get("uname").String()
		pass  = r.Get("pass").String()
	)
	if uname == "" || pass == "" {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT OR PASSWORD CANNOT BE EMPTY."))
		r.ExitAll()
	}
	admin, err := get.AdminByUname(ctx, uname)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ERROR."))
		r.ExitAll()
	}

	if !xpwd.ComparePassword(admin.Pwd, pass) {
		r.Response.WriteJson(gtoken.Fail("WRONG PASSWORD."))
		r.ExitAll()
	}

	one, err := dao.Role.Ctx(ctx).Where("id", admin.Rid).One()
	if err != nil || one == nil {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ROLE ERROR."))
		r.ExitAll()
	}
	info := model.UserInfo{
		Account:   uname,
		Uid:       gconv.Float64(admin.Id),
		ClientIp:  r.GetClientIp(),
		RuleId:    admin.Rid,
		RuleName:  one.Map()["name"].(string),
		Email:     admin.Email,
		Phone:     admin.Phone,
		LoginTime: time.Now(),
	}
	// 唯一标识，扩展参数user data
	loginLog := sysLogic.LoginLogin{
		Uid: uint64(admin.Id),
		Ip:  r.GetClientIp(),
	}
	if err = loginLog.Save(ctx); err != nil {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ROLE ERROR."))
		r.ExitAll()
	}
	return uname, info
}
