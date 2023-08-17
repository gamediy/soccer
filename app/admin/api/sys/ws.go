package sys

import "github.com/gogf/gf/v2/frame/g"

type AdminWsReq struct {
	g.Meta `tags:"/sys/ws" method:"get" path:"/ws/connect" dc:"长链接"`
}
type AdminWsRes struct {
	g.Meta `mine:"text/html" type:"string"  dc:""`
}
type SendMsgReq struct {
	g.Meta    `tags:"/sys/ws" method:"post" path:"/ws/send" dc:"通知单个管理员"`
	FromUname string `json:"fromUname" dc:"发送用户名 (不用填写)"`
	ToUname   string `json:"toUname" dc:"接受用户名 必填"`
	ToUid     uint64 `json:"toUid" dc:"接口用户ID (不用填写)"`
	Position  string `json:"position" dc:"位置:top-center"`
	Msg       string `v:"required" json:"msg"`
	Type      string `v:"required" d:"info" json:"type" dc:"类型：info,primary,success,error"` // 用于通知类型
}
type SendMsgRes struct{}

type SendAllReq struct {
	g.Meta    `tags:"/sys/ws" method:"post" path:"/ws/noticeAdmins" dc:"通知所有管理员"`
	FromUname string `json:"fromUname" dc:"发送用户名 (不用填写)"`
	ToUname   string `json:"toUname" dc:"接受用户名 必填"`
	ToUid     uint64 `json:"toUid" dc:"接口用户ID (不用填写)"`
	Position  string `json:"position" dc:"位置:top-center"`
	Msg       string `v:"required" json:"msg"`
	Type      string `v:"required" d:"info" json:"type" dc:"类型：info,primary,success,error"` // 用于通知类型
}
type SendAllRes struct{}
