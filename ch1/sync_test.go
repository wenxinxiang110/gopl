package ch1

import (
	"testing"
)

func Test_consumerAndProducer(t *testing.T) {
	cls := make(chan struct{}, 0)
	consumerAndProducer(10, cls)
	close(cls)
}
