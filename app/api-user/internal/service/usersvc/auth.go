package usersvc

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/consts"
	"star_net/model"
	"star_net/utility/utils/xtrans"
)

func UserLogin(r *ghttp.Request) (string, interface{}) {
	var (
		ctx = r.Context()
		req struct {
			Account  string `json:"account" v:"required#账户必填"`
			Password string `json:"password" v:"required#密码必填"`
		}
		lang = r.GetHeader("lang")
	)
	if lang == "" {
		lang = "en"
	}
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(gtoken.Fail(xtrans.T(lang, err.Error())))
	}
	login := Login{
		Ctx:      ctx,
		Account:  req.Account,
		Password: req.Password,
		Ip:       r.GetClientIp(),
	}
	userInfo, err := login.Exec()
	if err != nil {
		r.Response.WriteJsonExit(gtoken.Fail(xtrans.T(lang, err.Error())))
	}
	return userInfo.Account, userInfo
}
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Code != 0 {
		switch r.Method {
		case "PUT", "DELETE", "GET", "POST":
			r.Response.WriteJsonExit(respData)
		}
		return
	}
	var u model.UserInfo
	if err := gjson.New(respData.Data).Get("data").Struct(&u); err != nil {
		r.Response.WriteJsonExit(err)
	}
	r.SetCtxVar("account", u.Account)
	r.SetCtxVar("uid", u.UidInt64)
	r.SetCtxVar(consts.UserInfo, u)
	r.Middleware.Next()
}
