package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrTryAgain         = gerror.NewCode(gcode.New(-4, "失败请重试", ""))
	ErrRepeatedSubmit   = gerror.NewCode(gcode.New(-5, "请不要重复提交", ""))
	ErrUnameExist       = gerror.NewCode(gcode.New(-103, "用户名已存在", ""))
	ErrDepositClosed    = gerror.NewCode(gcode.New(-1000, "充值通道已关闭", ""))
	ErrDepositIncorrect = gerror.NewCode(gcode.New(-1001, "充值金额不正确", ""))
)
