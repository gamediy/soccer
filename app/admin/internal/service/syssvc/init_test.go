package syssvc

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"star_net/consts"
	"testing"
)

func TestInitD_SetImgPrefix(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InitD{}
			s.SetImgPrefix(tt.args.ctx)
		})
	}
	g.Dump(consts.ImgPrefix)
}

func TestInitD_SetWhiteIps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InitD{}
			if err := s.SetWhiteIps(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("SetWhiteIps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	g.Dump(consts.WhiteIps)
}
