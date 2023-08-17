package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// sys_api

	ErrAuthNotEnough    = gerror.NewCode(gcode.New(-1, "暂无当前操作权限", ""))
	ErrAuth             = gerror.NewCode(gcode.New(-2, "未认证", ""))
	ErrDataNotFound     = gerror.NewCode(gcode.New(-3, "数据不存在", ""))
	ErrImgCannotBeEmpty = gerror.NewCode(gcode.New(-4, "图片不能为空", ""))
	ErrDataAlreadyExist = gerror.NewCode(gcode.New(-5, "数据已存在", ""))
	ErrParamsEmpty      = gerror.NewCode(gcode.New(-6, "参数不能为空", ""))
	ErrActionFast       = gerror.NewCode(gcode.New(-7, "you are too fast", ""))
	ErrKeyAlreadyExist  = gerror.NewCode(gcode.New(-9, "KEY 已存在", ""))
	ErrCaptcha          = gerror.NewCode(gcode.New(-8, "验证码错误", ""))
	ErrUploadFile       = gerror.NewCode(gcode.New(-104, "上传文件失败", ""))
)
