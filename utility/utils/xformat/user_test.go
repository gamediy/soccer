package xformat

import (
	"context"
	"testing"
)

func TestAccount(t *testing.T) {
	type args struct {
		ctx     context.Context
		account string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{args: args{ctx: context.Background(), account: "a33366633"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Account(tt.args.ctx, tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("Account() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
