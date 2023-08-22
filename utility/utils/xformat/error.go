package xformat

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrPassport = gerror.NewCode(gcode.New(-200, "账户字母开头，只能包含字母、数字和下划线，长度在6~18之间", ""))
)
