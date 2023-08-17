package spassport

import (
	"context"
	captcha "github.com/mojocn/base64Captcha"
	vpassport "star_net/app/api-user/api/passport"
	"star_net/utility/utils/xcaptcha"
)

func (s *passport) GetCaptcha(ctx context.Context, req *vpassport.CaptchaReq) (res *vpassport.CaptchaRes, err error) {
	var driver = xcaptcha.NewDriver().ConvertFonts()
	cc := captcha.NewCaptcha(driver, xcaptcha.Store)
	_, content, answer := cc.Driver.GenerateIdQuestionAnswer()
	item, _ := cc.Driver.DrawCaptcha(content)
	_ = cc.Store.Set(req.CaptchaId, answer)
	return &vpassport.CaptchaRes{Img: item.EncodeB64string()}, nil
}
