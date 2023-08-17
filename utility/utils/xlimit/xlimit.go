package xlimit

import (
	"github.com/yudeguang/ratelimit"
)

func CreateRateLimit(fn func(rule *ratelimit.Rule)) *ratelimit.Rule {
	r := ratelimit.NewRule()
	fn(r)
	return r
}
