package depositsvc

import (
	"context"
	md "star_net/model"
	"star_net/utility/utils/xpusher"

	"star_net/utility/utils/xtrans"
	"testing"
)

func Test_deposit_Submit(t *testing.T) {
	background := context.Background()
	info := md.UserInfo{}
	info.Uid = 1
	info.Pid = 0
	info.Account = "account"
	info.I18n = xtrans.New(xtrans.EN)
	value := context.WithValue(background, consts.UserInfo, info)
	xpusher.InitFromCfg(value)
	type args struct {
		ctx   context.Context
		input DepositSubmitInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx: value,
				input: DepositSubmitInput{
					PayId:           1,
					Amount:          3,
					TransferOrderNo: "1234567892",
					TranserImg:      "https://5b0988e595225.cdn.sohucs.com/images/20190216/35011958e979447ba9134de2a10dc657.jpeg",
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
