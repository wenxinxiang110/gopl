package ch7

import (
	"bytes"
	"testing"
)

func TestCountingWriterFactoryImpl_CountingWriter(t *testing.T) {

	var (
		buff       = new(bytes.Buffer)
		ft         = new(CountingWriterFactoryImpl)
		w, counter = ft.CountingWriter(buff)
	)

	if *counter != 0 {
		t.Errorf("counter is %d", *counter)
	}

	_, _ = w.Write([]byte{'1'})
	if *counter != 1 {
		t.Errorf("counter is %d", *counter)
	}
	s := buff.String()
	if s != "1" {
		t.Errorf("writer is %s", s)
	}
}
