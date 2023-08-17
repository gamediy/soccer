package xjwt

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestGenToken(t *testing.T) {
	g.Dump(GenToken("0745944159", 10060, 0))
}
func TestParseToken(t *testing.T) {
	g.Dump(parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA3MDgzODcyODYiLCJ1aWQiOjE1MCwiUmlkIjowLCJleHAiOjE2ODI5NDE4ODIsImlzcyI6Im15LXByb2plY3QifQ.QRG7cRoTNlUp4g8beX2MccI85SZbEHBcP89ytdl8ADE"))
}
