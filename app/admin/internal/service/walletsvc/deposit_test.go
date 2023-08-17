package walletsvc

import (
	"context"
	"star_net/app/admin/internal/model"
	"testing"
)

func Test_sDeposit_Update(t *testing.T) {
	background := context.Background()
	info := model.UserInfo{}
	info.Uid = 1
	info.Account = "account"
	info.ClientIp = "127.0.0.1"
	value := context.WithValue(background, "userInfo", info)
	type args struct {
		ctx          context.Context
		orderNo      int64
		status       int
		statusRemark string
		sysRemark    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx:          value,
				orderNo:      1689983101553348608,
				status:       1,
				statusRemark: "充值成功",
				sysRemark:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := Deposit{
				OrderNo:      tt.args.orderNo,
				Status:       tt.args.status,
				StatusRemark: tt.args.statusRemark,
				SysRemark:    tt.args.sysRemark,
			}
			if err := sd.Update(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
