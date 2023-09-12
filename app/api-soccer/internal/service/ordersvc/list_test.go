package ordersvc

import (
	"context"
	"reflect"
	"star_net/utility/utils/xtest"
	"testing"
)

func TestOrderList_Exec(t *testing.T) {
	type fields struct {
		Type   int
		Status int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []OrderListOutput
	}{
		{
			args: args{
				ctx: xtest.GetContext(),
			},
			fields: fields{
				Status: 1,
				Type:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &OrderList{
				Type:   tt.fields.Type,
				Status: tt.fields.Status,
			}
			if got := order.Exec(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
