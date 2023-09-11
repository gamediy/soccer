package xip

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestGetIpInfo(t *testing.T) {
	location := GetIpInfoIo(context.TODO(), "102.215.252.23")

	g.Dump(location)
}
