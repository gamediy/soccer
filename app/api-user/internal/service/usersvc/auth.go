package usersvc

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/model"
)

func Login(r *ghttp.Request) (string, interface{}) {
	var (
		//ctx = r.Context()
		req struct {
			Uname string `json:"uname" v:"required"`
			Pass  string `json:"pass" v:"required"`
		}
	)
	if err := r.Parse(&req); err != nil {
		r.ExitAll()
	}
	//
	//if uname == "" || pass == "" {
	//	r.Response.WriteJson(gtoken.Fail("ACCOUNT OR PASSWORD CANNOT BE EMPTY."))
	//	r.ExitAll()
	//}

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
