package depositsvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"star_net/model"
	"star_net/utility/utils/xpusher"
	"testing"
)

func init() {
	xpusher.InitFromCfg(context.Background())
	//xtrans.Init()
}

func TestSubmit_Exec(t *testing.T) {
	type fields struct {
		PayId           int
		Amount          float64
		TransferOrderNo string
		TranserImg      string
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
				PayId:           85,
				Amount:          200,
				TranserImg:      "Uvteyd.png",
				TransferOrderNo: "1231ff",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := &Submit{
				PayId:           tt.fields.PayId,
				Amount:          tt.fields.Amount,
				TransferOrderNo: tt.fields.TransferOrderNo,
				TranserImg:      tt.fields.TranserImg,
			}
			g.Dump(input)
		})
	}
}
