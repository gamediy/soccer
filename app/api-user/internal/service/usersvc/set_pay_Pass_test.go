package usersvc

import (
	"context"
	"testing"
)

func TestSetPayPass_Exec(t *testing.T) {
	type fields struct {
		Pass string
		Uid  int64
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
		{fields: fields{
			Uid:  121,
			Pass: "123456",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SetPayPass{
				Pass: tt.fields.Pass,
				Uid:  tt.fields.Uid,
			}
			if err := s.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
