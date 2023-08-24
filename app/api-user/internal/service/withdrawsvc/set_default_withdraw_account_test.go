package withdrawsvc

import (
	"context"
	"star_net/consts"
	"star_net/model"
	"testing"
)

func TestSetDefaultWithdrawAccount_Exec(t *testing.T) {
	type fields struct {
		Id int64
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
			fields: fields{Id: 1694634395836616704},
			args:   args{ctx: context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{Uid: 121})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SetDefaultWithdrawAccount{
				Id: tt.fields.Id,
			}
			if err := s.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
