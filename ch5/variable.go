package ch5

import "math"

func Max(nums ...int) int {
	max := math.MinInt32

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
