package walletsvc

import (
	"context"
	"star_net/app/admin/internal/model"
	"star_net/consts"
	"testing"
)

func Test_sWithdraw_Update(t *testing.T) {
	background := context.Background()
	info := model.UserInfo{}
	info.Uid = 1
	info.Account = "account"
	info.ClientIp = "127.0.0.1"
	value := context.WithValue(background, consts.UserInfo, info)
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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := Withdraw{
				OrderNo:      1690614696714964992,
				Status:       consts.WithdrawStatusSuccess,
				StatusRemark: "success",
				SysRemark:    "",
			}
			if err := sw.Update(value); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
