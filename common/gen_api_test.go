package common

import (
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestHandleGenApi(t *testing.T) {
	menu := []string{
		"setting",
		"user",
		"wallet",
	}
	for _, s := range menu {
		if err := HandleGenApi(gctx.New(), "../app/admin/api/"+s, "/api/"+s); err != nil {
			t.Fatal(err)
		}
	}

}
