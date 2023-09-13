package withdrawsvc

import (
	"context"
	"star_net/consts"
	"star_net/model"
	"testing"
)

func TestBindWithdrawAccount_Exec(t *testing.T) {
	type fields struct {
		BankId  uint64
		Address string
		Title   string
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
			fields: fields{
				BankId:  1,
				Address: "3123441333",
				Title:   "John",
			},
			args: args{ctx: context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{UidInt64: 121, Account: "john"})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BindWithdrawAccount{

				Address: tt.fields.Address,
				Title:   tt.fields.Title,
			}
			if err := s.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
