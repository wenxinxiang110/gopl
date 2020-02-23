package ch6

import (
	"fmt"
	"image/color"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	defer cache.Unlock()

	v := cache.mapping[key]

	return v
}

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	*Point
	Color color.Color
}

func Test() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

}
