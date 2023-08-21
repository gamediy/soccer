package withdrawsvc

import (
	"context"
	"star_net/app/api-user/internal/model"
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
	value := context.WithValue(background, "userInfo", info)
	type args struct {
		ctx   context.Context
		input model.WithdrawSubmitInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx: value,
				input: model.WithdrawSubmitInput{
					WithdrawId:        5,
					WithdrawAccountId: 1,
					Amount:            13,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &withdraw{}
			if err := s.Submit(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
