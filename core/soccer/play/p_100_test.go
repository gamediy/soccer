package play

import "testing"

func Test_getProfit(t *testing.T) {
	type args struct {
		openResult OpenResult
		calcInfo   CalcInfo
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				openResult: OpenResult{
					Result: "1-1",
				},
				calcInfo: CalcInfo{
					BetAmount: 100,
					Odds:      1.8,
					Rule:      "Over 2/2.5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getP400Profit(tt.args.openResult, tt.args.calcInfo); got != tt.want {
				t.Errorf("getProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
