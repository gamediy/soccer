package soccer

import (
	"context"
	"star_net/core/soccer/play"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		ctx     context.Context
		eventId int64
		result  play.OpenResult
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx:     context.TODO(),
				eventId: 1,
				result: play.OpenResult{
					Result:     "1-4",
					BoutStatus: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Calc(tt.args.ctx, tt.args.eventId, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
