package xcaptcha

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mojocn/base64Captcha"
	captcha "github.com/mojocn/base64Captcha"
	"star_net/consts"
)

var (
	Store = base64Captcha.DefaultMemStore
)

func NewDriver() *base64Captcha.DriverString {
	driver := new(base64Captcha.DriverString)
	driver.Height = 60
	driver.Width = 187
	driver.NoiseCount = 1
	//driver.ShowLineOptions = base64Captcha.OptionShowSlimeLine
	driver.Length = 4
	driver.Source = "1234567890"
	driver.Fonts = []string{"wqy-microhei.ttc", "3Dumb.ttf"}
	return driver
}

func Get(ctx context.Context, id string) (string, string, error) {
	var driver = NewDriver().ConvertFonts()
	cc := captcha.NewCaptcha(driver, Store)
	_, content, answer := cc.Driver.GenerateIdQuestionAnswer()
	item, err := cc.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}
	_ = cc.Store.Set(id, answer)
	g.Log().Infof(ctx, "captcha info ID:%s, answer:%s", id, answer)
	return item.EncodeB64string(), answer, nil
}
func Verify(id, answer string) error {
	if !Store.Verify(id, answer, true) || id == "" || answer == "" {
		return consts.ErrCaptcha
	}
	return nil
}
