package chapter3

import (
	"bytes"
	"fmt"
)

// utf-8编码，不存在说某个字符的编码是其他字符的子串，因此可以这样直接查找
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 熟悉一下bytes包的api
func IntstoString(arr []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range arr {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteString("]")
	return buf.String()
}

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Copy() {
	a := make([]int, 0)
	b := []int{1, 2, 3}

	copy(a, b)
}

// 让数组s向右旋转n次
func Rotate(n int, s []string) []string {
	n %= len(s)

	front := s[len(s)-n:]
	back := s[:len(s)-n]

	return append(front, back...)
}
