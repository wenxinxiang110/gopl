package ch8

import (
	"fmt"
	"os"
	"time"
)

// 模拟火箭倒计时
func Down() {
	fmt.Println("Commencing countdown.  Press return to abort.")

	// 读取打断的信号
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// time.Tick（）无法关闭,除非整个程序的生命周期都要用
	//tick := time.Tick(time.Second)
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for cd := 10; cd > 0; cd-- {
		println(cd)

		// 多个case同时就绪，select会随机选择一个！
		select {
		case <-tick.C:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		default:
			//	do something
		}
	}

	// 发射
	//launch()
}
