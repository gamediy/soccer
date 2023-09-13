package usersvc

import (
	"context"
	"testing"
)

func TestChangePayPass_Exec(t *testing.T) {
	type fields struct {
		PayPass string
		Pass    string
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
				PayPass: "123456",
				Pass:    "a1231231",
				Address: "3123441333",
				Title:   "John",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ChangePayPass{
				PayPass: tt.fields.PayPass,
				Pass:    tt.fields.Pass,
				Address: tt.fields.Address,
				Title:   tt.fields.Title,
			}
			if err := s.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
