package soccer

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/api/soccer"
	"testing"
)

func Test_cPlay_GetPalyAll(t *testing.T) {

	c := cPlay{}
	g.Dump(c.GetPlayAll(context.Background(), &soccer.GetPlayAllReq{}))

}
