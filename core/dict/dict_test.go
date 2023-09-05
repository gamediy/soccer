package dict

import (
	"context"
	"testing"
)

func TestGetVAsString(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				ctx: context.Background(),
				key: ImageDomain,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVAsString(tt.args.ctx, tt.args.key); got != tt.want {
				t.Errorf("GetVAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
