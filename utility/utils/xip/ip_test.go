package xip

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestGetIpInfo(t *testing.T) {
	location, err := GetIpInfo("102.215.252.23")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(location)
}
