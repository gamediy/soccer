package spassport

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	vpassport "star_net/app/api-user/api/passport"
	"testing"
)

func Test_passport_Register(t *testing.T) {
	background := context.Background()
	ctx := context.WithValue(background, "User-Agent", "user-agent-33")
	type args struct {
		ctx context.Context
		req *vpassport.RegisterReq
	}
	tests := []struct {
		name    string
		args    args
		wantRes *vpassport.RegisterRes
		wantErr bool
	}{
		{
			args: args{
				ctx: ctx,
				req: &vpassport.RegisterReq{
					Account:  "user33",
					Xid:      "test",
					Email:    "test@gmail.com",
					Phone:    "98787775676",
					Password: "a1231231",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &passport{}
			gotRes, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(gotRes)
		})
	}
}
