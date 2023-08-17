package reportsvc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
	"testing"
)

func TestWalletLog_Report(t *testing.T) {

	type fields struct {
		StartTime string
		EndTime   string
		Account   string
		PageIndex int
		PageSize  int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *gdb.List
		wantErr bool
	}{
		{
			args: args{
				ctx: context.TODO(),
			},
			fields: fields{
				StartTime: "2023-08-10",
				EndTime:   "2023-08-20",
				PageSize:  10,
				PageIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &walletLog{
				Begin:   tt.fields.StartTime,
				End:     tt.fields.EndTime,
				Account: tt.fields.Account,
				Page:    tt.fields.PageIndex,
				Size:    tt.fields.PageSize,
			}
			gotRes, err := s.Report(tt.args.ctx)
			g.Dump(gotRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Report() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
