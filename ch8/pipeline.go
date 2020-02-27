package ch8

//1. 读取一个关闭的channel时，channel返回对应的0值；
//2. range会一直循环到channel关闭
//3. 可以用cap()方法获取channel的容量;len()方法返回的是元素的个数

func PipDemo() {

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go squarer(naturals, squares)

	printer(squares)
}

func counter(writer chan<- int) {
	for x := 0; x < 100; x++ {
		writer <- x
	}
	close(writer)
}
func squarer(reader <-chan int, writer chan<- int) {
	for x := range reader {
		writer <- x
	}
	close(writer)
}
func printer(reader <-chan int) {
	for x := range reader {
		println(x)
	}
}
