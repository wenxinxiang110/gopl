package ch9

import (
	"fmt"
	"time"
)

func Cal() {
	//done := make(chan struct{})
	var done bool
	ch1 := make(chan string, 20)
	ch2 := make(chan string, 20)

	send := func(s string, ch chan<- string) {
		for !done {
			ch <- s
		}
		close(ch)
	}

	go send("ping", ch1)
	go send("pong", ch2)

	stick := time.NewTicker(time.Second)

	var n int
	for {
		select {
		case <-stick.C:
			done = true
			stick.Stop()
			fmt.Printf("1s内计算%v次", n)
			return
		case <-ch1:
			n++
		case <-ch2:
			n++
		}
	}

}
