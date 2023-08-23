package walletsvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"star_net/db/model/entity"
	"star_net/model"
	"testing"
)

func TestLogs_Exec(t *testing.T) {
	type fields struct {
		Page int
		Size int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   []*entity.WalletLog
		wantErr bool
	}{
		{
			args: args{ctx: context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{
				UidInt64: 121,
			})},
			fields: fields{
				Page: 1,
				Size: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Logs{
				Page: tt.fields.Page,
				Size: tt.fields.Size,
			}
			got, got1, err := s.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got, got1)
		})
	}
}
