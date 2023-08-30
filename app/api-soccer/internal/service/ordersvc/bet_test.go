package ordersvc

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/model"
	"star_net/utility/utils/xtest"
	"testing"
)

func Test_betAll(t *testing.T) {
	list := []entity.EventsOdds{}
	dao.EventsOdds.Ctx(context.TODO()).Scan(&list)
	for _, odds := range list {
		bet := bet{}
		bet.BetInput.Amount = 1
		bet.BetInput.OddsId = odds.Id
		bet.Exec(xtest.GetContext())
	}
}

func Test_bet_Exec(t *testing.T) {
	fmt.Println(g.Config().MustGet(context.TODO(), "database.default.link"))
	type fields struct {
		BetInput BetInput
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{ctx: context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{UidInt64: 121, Account: "join"})},
			fields: fields{
				BetInput: BetInput{
					OddsId: 3,
					Amount: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := &bet{
				BetInput: tt.fields.BetInput,
			}
			if err := input.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
