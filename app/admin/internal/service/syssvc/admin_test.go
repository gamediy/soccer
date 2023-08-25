package syssvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/internal/model"
	"star_net/consts"
	"testing"
)

func Test_sAdmin_UserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		args         args
		wantMenu     []*model.Menu
		wantButton   []*model.Menu
		wantUserInfo model.UserInfo
		wantErr      bool
	}{
		// TODO: Add test cases.
		{args: args{ctx: context.WithValue(context.Background(), consts.UserInfo, model.UserInfo{RuleName: "admin"})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := sAdmin{}
			gotMenu, _, gotUserInfo, err := sa.UserInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(gotMenu)
			g.Dump(gotUserInfo)

		})
	}
}
