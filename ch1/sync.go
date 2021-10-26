package ch1

import (
	"log"
	"time"
)

func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		out <- i * factor
	}
}

func consumer(in <-chan int) {
	for data := range in {
		log.Println(data)
	}
}

func consumerAndProducer(bufferSize int, closeSign <-chan struct{}) {
	ch := make(chan int, bufferSize)
	go producer(3, ch)
	go consumer(ch)
	//go producer(5, ch)
	select {
	case <-closeSign:
		close(ch)
		return
	}
}
