package syssvc

import (
	"context"
	"star_net/app/admin/internal/logic/sysLogic"
	"star_net/app/admin/internal/model"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/utility/utils/xpwd"
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
