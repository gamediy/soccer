package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	g.Meta   `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
	UserName string `v:"required#用户名必填"`
}

type Res struct {
	g.Meta `mime:"text/html" example:"string"`
}
