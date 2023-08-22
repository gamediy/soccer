package common

import (
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yudeguang/ratelimit"
	"net/http"
	"star_net/utility/utils/xlimit"
	"star_net/utility/utils/xtrans"
	"time"
)

func MiddlewareDefaultCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		resp = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}
	lang := r.GetHeader("lang")
	if lang == "" {
		lang = "en"
	}
	fmt.Println(msg)

	m := xtrans.T(lang, msg)
	r.Response.WriteJson(g.Map{
		"code": code.Code(),
		"msg":  m,
		"data": resp,
	})
}

func MiddlewareRequestLimit(r *ghttp.Request) {
	ip := r.GetClientIp()
	limit := xlimit.CreateRateLimit(func(rule *ratelimit.Rule) {
		rule.AddRule(time.Hour, 10000)
		rule.AddRule(time.Minute, 600)
		rule.AddRule(time.Second, 10)
	})
	ok := limit.AllowVisit(ip)
	if !ok {
		g2 := gtoken.Resp{
			Code: -999,
			Msg:  "YOUR ACCESS IS ABNORMAL",
		}
		r.Response.WriteJsonExit(g2)

	}
	r.Middleware.Next()
}
