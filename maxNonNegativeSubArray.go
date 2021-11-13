package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, -3, 9, 2, 4, -8, 8, 5}
	fmt.Println(maxSub(array))
}

func maxSub(a []int) int {
	len := len(a)
	sum := 0
	i := 0
	max := -1
	for i < len {
		for a[i] < 0 && i < len {
			i++
		}
		for i < len && a[i] >= 0 {
			sum += a[i]
			i++
			if sum > max {
				max = sum
			}
		}
		sum = 0
	}
	return max
}
