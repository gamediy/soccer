package usersvc

import (
	"context"
	"testing"
)

func TestPutLoginPass_Exec(t *testing.T) {
	type fields struct {
		Ctx     context.Context
		Uid     uint64
		OldPass string
		NewPass string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{fields: fields{
			Ctx:     context.Background(),
			Uid:     121,
			OldPass: "a1231231",
			NewPass: "a1231231",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PutLoginPass{
				Ctx:     tt.fields.Ctx,
				Uid:     tt.fields.Uid,
				OldPass: tt.fields.OldPass,
				NewPass: tt.fields.NewPass,
			}
			if err := s.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
