package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/internal/model"
)

type InfoReq struct {
	g.Meta `tags:"用户" method:"get" path:"/info" dc:"用户信息"`
}
type InfoRes struct {
	*model.User
	*model.Wallet
}
