package usersvc

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/model"
)

func Login(r *ghttp.Request) (string, interface{}) {
	return "", nil
}
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Code != 0 {
		switch r.Method {
		case "PUT", "DELETE", "GET", "POST":
			r.Response.WriteJsonExit(respData)
		}
		return
	}
	var (
		s2       = respData.Data.(map[string]interface{})
		userInfo = s2["data"].(map[string]interface{})
		u        = model.UserInfo{
			Uid:     userInfo["uid"].(float64),
			Account: userInfo["account"].(string),
		}
	)
	r.SetCtxVar("userName", s2["userKey"])
	r.SetCtxVar("roleName", userInfo["ruleName"])
	r.SetCtxVar("uid", userInfo["uid"])
	r.SetCtxVar("account", userInfo["account"])
	r.SetCtxVar("userInfo", u)
}
