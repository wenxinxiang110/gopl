package ch2

var pc [256]byte

func init() {
	// i&1: i为奇数则为1，偶数则为0
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 统计二进制中1bit的个数
// todo:不懂，日后再说
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
