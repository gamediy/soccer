package walletsvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"testing"
)

func TestListBalanceCodes_Exec(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*entity.BalanceCode
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ListBalanceCodes{}
			got, err := s.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
