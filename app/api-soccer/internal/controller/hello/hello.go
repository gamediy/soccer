package hello

import (
	"context"
	"fmt"
	v1 "star_net/app/api-soccer/api/hello/v1"
)

var (
	Ctrl = &cHello{}
)

type cHello struct {
}

func (cHello) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {

	return nil, fmt.Errorf("用户名已存在")
}
