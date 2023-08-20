package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrTryAgain       = gerror.NewCode(gcode.New(-4, "失败请重试", ""))
	ErrRepeatedSubmit = gerror.NewCode(gcode.New(-5, "请不要重复提交", ""))
)
