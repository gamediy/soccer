package usersvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/admin/api/user"
	adminModel "star_net/app/admin/internal/model"
	"star_net/model"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestReadList_Exec(t *testing.T) {
	type fields struct {
		Ctx context.Context
		Req *user.ReadListUserReq
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*adminModel.User
		want1   int
		wantErr bool
	}{
		{
			fields: fields{
				Ctx: context.Background(),
				Req: &user.ReadListUserReq{CommonPageReq: model.CommonPageReq{Size: 10}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ReadList{
				Ctx: tt.fields.Ctx,
				Req: tt.fields.Req,
			}
			got, got1, err := s.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got, got1)
		})
	}
}
