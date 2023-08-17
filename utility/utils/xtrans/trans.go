package xtrans

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

const (
	ZH = "zh"
	EN = "en"
)

func T(language interface{}, word string) string {
	if language == nil {
		language = "en"
	}
	var (
		ctx  = context.TODO()
		i18n = g.I18n()
	)
	if err := i18n.SetPath(fmt.Sprint("/resource/i18n")); err != nil {
		panic(err)
	}
	i18n.SetLanguage(fmt.Sprint(language))
	return i18n.Translate(ctx, word)
}

func New(language interface{}) *gi18n.Manager {
	if language == nil {
		language = "en"
	}
	var (
		i18n = gi18n.New()
	)
	if err := i18n.SetPath(fmt.Sprint("/resource/i18n")); err != nil {
		panic(err)
	}
	i18n.SetLanguage(fmt.Sprint(language))
	return i18n
}
