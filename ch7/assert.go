package ch7

import (
	"fmt"
	"go/types"
)

func ToString(x interface{}) string {
	switch x := x.(type) {
	case types.Nil:
		return "Null"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "True"
		}
		return "False"
	case string:
		return x
	}
	panic(fmt.Sprintf("unexpected type %T: %v", x, x))
}
