package xpusher

import (
	"testing"
)

func init() {
	Init("1644349", "3400fc3f83796ca48785", "2e84608a0d38f20e852e", "ap1")
}
func TestTrigger(t *testing.T) {
	data := map[string]string{"message": "hello world"}
	err := Trigger("admins", "share", data)
	if err != nil {
		t.Fatal(err)
	}
}
