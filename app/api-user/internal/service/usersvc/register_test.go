package usersvc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestRegister_Exec(t *testing.T) {
	type fields struct {
		Ctx      context.Context
		RealName string
		Account  string
		Password string
		Xid      string
		Phone    string
		Email    string
		Country  string
		City     string
		Ip       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			fields: fields{
				Ctx:      gctx.New(),
				RealName: "Tom",
				Account:  "Tom",
				Password: "a1231231",
				Country:  "Japan",
				City:     "Tokyo",
				Ip:       "localhost",
				Phone:    "38112093123",
				Email:    "test@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Register{
				Ctx:      tt.fields.Ctx,
				RealName: tt.fields.RealName,
				Account:  tt.fields.Account,
				Password: tt.fields.Password,
				Xid:      tt.fields.Xid,
				Phone:    tt.fields.Phone,
				Email:    tt.fields.Email,
				Ip:       tt.fields.Ip,
			}
			got, err := s.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
