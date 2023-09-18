package soccer

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"reflect"
	"star_net/app/admin/api/soccer"
	"star_net/model"
	"testing"
)

func Test_cEvents_OpenResult(t *testing.T) {
	type args struct {
		ctx gctx.Ctx
		req *soccer.OpenResultReq
	}
	tests := []struct {
		name    string
		args    args
		wantRes *model.CommonRes
		wantErr bool
	}{
		{
			args: args{
				ctx: context.Background(),
				req: &soccer.OpenResultReq{
					Result:     "1-9",
					BoutStatus: 1,
					EventsId:   11,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cEvents{}
			gotRes, err := c.OpenResult(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("OpenResult() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
