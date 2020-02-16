package chapter1

// switch特别用法demo
//不带操作对象（译注：switch不带操作对象时默认用true值代替，然后将每个case的表达式和true值进行比较）；
//可以直接罗列多种条件，像其它语言里面的多个if else一样，下面是一个例子：
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
