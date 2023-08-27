package eventsvc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/model"
	"star_net/utility/utils/xtrans"
	"testing"
)

func Test_list_Exec(t *testing.T) {
	var u model.UserInfo
	u.Lang = "en"
	u.I18n = xtrans.New(u.Lang)
	value := context.WithValue(context.Background(), "userInfo", u)

	g.Dump(value)

}
