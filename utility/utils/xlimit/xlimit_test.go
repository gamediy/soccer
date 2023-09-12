package xlimit

import (
	"github.com/yudeguang/ratelimit"
	"testing"
	"time"
)

func TestCreateRateLimit(t *testing.T) {
	limit := CreateRateLimit(func(rule *ratelimit.Rule) {
		rule.AddRule(time.Minute*2, 2)
		rule.AddRule(time.Second*10, 1)
	})
	for i := 0; i < 3333; i++ {
		go func() {
			println(i, limit.AllowVisit("123"))
		}()
		time.Sleep(time.Second)
	}
	select {}
}
