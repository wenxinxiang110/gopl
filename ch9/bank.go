package ch9

import "sync"

var (
	// 通过一个buffer为1的channel达到sync.mutex的效果
	sema = make(chan struct{}, 1)

	mu sync.Mutex

	balance int
)

func Deposit(amount int) {
	//
	//sema <- struct{}{}
	//defer func() {
	//	<-sema
	//}()

	mu.Lock()
	defer mu.Unlock()
	balance += amount

}
