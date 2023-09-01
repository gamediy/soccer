package eventsvc

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/utility/utils/xtest"
	"testing"
)

func Test_list_Exec(t *testing.T) {
	PlayOdds.EventId = 1

	g.Dump(PlayOdds.Exec(xtest.GetContext()))

}
