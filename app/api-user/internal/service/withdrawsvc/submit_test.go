package withdrawsvc

import (
	"context"
	"star_net/consts"
	"star_net/model"

	"star_net/utility/utils/xtrans"
	"testing"
)

func Test_withdraw_Submit(t *testing.T) {

	background := context.Background()
	info := model.UserInfo{}
	info.Uid = 1
	info.Pid = 0
	info.I18n = xtrans.New(xtrans.EN)
	info.Account = "account"
	value := context.WithValue(background, consts.UserInfo, info)
	type args struct {
		ctx   context.Context
		input WithdrawSubmitInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx: value,
				input: WithdrawSubmitInput{
					WithdrawId:        5,
					WithdrawAccountId: 1,
					Amount:            13,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := Submit.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
