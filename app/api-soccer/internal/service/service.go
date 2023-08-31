package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"star_net/model"
)

func GetUserInfo(ctx context.Context) *model.UserInfo {
	info, ok := ctx.Value(consts.UserInfo).(model.UserInfo)
	if !ok {

		g.Log().Error(ctx, "用户不存在")
		return &model.UserInfo{
			Lang: "zh",
		}
	}
	return &info

}
