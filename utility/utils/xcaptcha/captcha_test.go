package xcaptcha

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestGet(t *testing.T) {
	id := "123"
	_, answer, err := Get(gctx.New(), id)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(Verify(id, answer))
}
