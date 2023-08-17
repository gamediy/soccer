package syssvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/internal/model"
	"star_net/db/model/entity"
	"testing"
)

func TestAddRole(t *testing.T) {
	var model model.Role
	model.Name = "test"
	model.Menu = []*entity.RoleMenu{
		{MenuId: 1},
		{MenuId: 2},
	}
	g.Dump(Role.Add(context.Background(), model))
}

func TestDelRole(t *testing.T) {
	type args struct {
		ctx    context.Context
		roleId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete role",
			args: args{
				ctx:    context.Background(),
				roleId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Role.Del(tt.args.ctx, tt.args.roleId); (err != nil) != tt.wantErr {
				t.Errorf("DelRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditRole(t *testing.T) {
	var m model.Role
	m.Id = 2
	m.Name = "test223"
	m.Menu = []*entity.RoleMenu{
		{
			MenuId: 1,
		},
		{
			MenuId: 161,
		},
	}
	type args struct {
		ctx  context.Context
		role model.Role
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "edit role",
			args: args{
				role: m,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Role.Edit(tt.args.ctx, tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("EditRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
