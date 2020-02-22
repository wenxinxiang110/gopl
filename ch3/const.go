package ch3

import (
	"fmt"
)

const (
	a = 1
	b
	c = 2
	d
)

// iota高级用法
const (
	flag1 = 1 << iota
	flag2 // 10
	flag3 // 100
	//	...

)

func TestConst() {

	fmt.Println(a, b, c, d) // "1 1 2 2"
}
