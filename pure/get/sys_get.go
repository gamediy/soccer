package get

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"star_net/consts"
)

func AdminIdFromCtx(ctx context.Context) uint64 {
	return ghttp.RequestFromCtx(ctx).Get(consts.TokenAdminIdKey).Uint64()
}

func AdminUnameFromCtx(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Get(consts.TokenAdminUname).String()
}
