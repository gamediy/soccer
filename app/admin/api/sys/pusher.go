package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/model"
)

type SendAdminsReq struct {
	g.Meta `tags:"/sys/pusher" path:"/pusher/sendAdmins" method:"post" dc:"推送通知管理员"`
	Data   string `json:"data"`
}
type SendAdminsRes model.CommonRes
