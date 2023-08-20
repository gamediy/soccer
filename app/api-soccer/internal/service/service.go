package service

import (
	"context"
	"star_net/model"
)

func GetUserInfo(ctx context.Context) *model.UserInfo {
	info := ctx.Value("userInfo").(model.UserInfo)
	return &info

}
