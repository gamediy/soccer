package syssvc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"reflect"
	"star_net/app/admin/internal/model"
	"testing"
)

func TestGetMenuAll(t *testing.T) {

}

func Test_sMenu_ListByRoleId(t *testing.T) {
	type args struct {
		ctx    context.Context
		roleId int
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Menu
		wantErr bool
	}{
		{
			args: args{
				roleId: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sMenu{}
			got, err := s.ListMenuByRoleId(tt.args.ctx, tt.args.roleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByRoleId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListByRoleId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
