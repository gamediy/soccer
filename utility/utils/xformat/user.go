package xformat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func Account(ctx context.Context, account string) error {
	if err := g.Validator().Data(account).Rules("passport").Run(ctx); err != nil {
		return ErrPassport
	}
	return nil
}
