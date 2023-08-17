package report

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"star_net/db/model/entity"
	"testing"
)

func TestReport_PutReport(t *testing.T) {
	type fields struct {
		User        entity.User
		Amount      float64
		BalanceCode int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			fields: fields{
				User: entity.User{
					Id: 1,
				},
				Amount:      1,
				BalanceCode: 501,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Report{
				User:        tt.fields.User,
				Amount:      tt.fields.Amount,
				BalanceCode: tt.fields.BalanceCode,
			}
			if err := this.PutReport(); (err != nil) != tt.wantErr {
				t.Errorf("PutReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
