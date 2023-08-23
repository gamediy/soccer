package walletsvc

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/db/model/entity"
	"testing"
)

func TestBanks_Exec(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want []*entity.Bank
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Bank{}
			got := s.GetBanks(tt.args.ctx)
			g.Dump(got)
		})
	}
}
