package xredis

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

var (
	defaultTime = 1000
)

var (
	ErrActionFast = gerror.NewCode(gcode.New(-7, "You are too fast", ""))
)

func Lock(ctx context.Context, key string, lockExpire int, fun func(ctx context.Context) error) error {
	if lockExpire == 0 {
		lockExpire = defaultTime
	}
	ok, err := g.Redis().Do(context.TODO(), "set", key, time.Now(), "nx", "ex", lockExpire)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %s", err)
	}
	if ok.String() == "OK" {
		defer func() {
			g.Redis().Do(context.TODO(), "del", key)
		}()
		if err = fun(ctx); err != nil {
			return err
		}
	} else {
		return ErrActionFast
	}
	return nil
}
