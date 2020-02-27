package ch8

import (
	"fmt"
	"time"
)

func Cal() {
	go Spinner(100 * time.Millisecond)
	const n = 45
	fibN := Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}
