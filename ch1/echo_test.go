package ch1

import (
	"bytes"
	"testing"
)

func TestEchoOSArgs(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	if err := EchoOSArgs(buf, "\n"); err != nil {
		t.Error(err)
	}
	t.Logf("success,os.Args is\n%s", buf.String())
}
