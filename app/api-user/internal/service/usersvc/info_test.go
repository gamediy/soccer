package usersvc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/app/api-user/api/user"
	"star_net/db/model/entity"
	"testing"
)

func TestGetInfo_Exec(t *testing.T) {
	type fields struct {
		Ctx context.Context
		Uid int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *entity.User
		want1   *entity.Wallet
		wantErr bool
	}{
		{
			fields: fields{Ctx: context.Background(), Uid: 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GetInfo{
				Ctx: tt.fields.Ctx,
				Uid: tt.fields.Uid,
			}
			got, got1, err := m.Exec()
			res := user.InfoRes{
				User:   got,
				Wallet: got1,
			}
			g.Dump(res)
			marshal, _ := json.Marshal(&res)
			fmt.Println(string(marshal))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got, got1)
		})
	}
}
