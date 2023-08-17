package xuuid

import (
	"fmt"
	"testing"
)

func TestConvertToBase30(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			args: args{input: 10020},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberToBase34(tt.args.input)
			fmt.Println(got)
		})
	}
}
