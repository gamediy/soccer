package user

import (
	"context"
	"star_net/app/api-user/api/user"
	"star_net/utility/utils/xcaptcha"
)

func (c *cUser) GetCaptcha(ctx context.Context, req *user.CaptchaReq) (res *user.CaptchaRes, err error) {
	img, _, err := xcaptcha.Get(ctx, req.CaptchaId)
	if err != nil {
		return nil, err
	}
	return &user.CaptchaRes{Img: img}, nil
}
