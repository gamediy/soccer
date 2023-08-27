package play

import "testing"

func Test_getP400Profit(t *testing.T) {
	type args struct {
		result OpenResult
		info   CalcInfo
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{result: OpenResult{
				Result: "3-1",
			}, info: CalcInfo{
				BetAmount: 100,
				Rule:      "Home -1.5",
				Handicap:  1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getP400Profit(tt.args.result, tt.args.info); got != tt.want {
				t.Errorf("getP400Profit() = %v, want %v", got, tt.want)
			}
		})
	}
}
