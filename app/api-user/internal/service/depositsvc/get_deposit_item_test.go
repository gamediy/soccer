package depositsvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"star_net/model"
	"testing"
)

func Test_getDepositItem_Exec(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDepositItemOutput
		wantErr bool
	}{
		{
			args: args{
				ctx: context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{
					UidInt64: 121,
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &getDepositItem{}
			got, err := s.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
