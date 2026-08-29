package main

import (
	"fmt"
	"math"
)

func main() {
	s := make([]int, 10)

	for i := 0; i < 10; i++ {
		s[i] = i * i
	}

	// target value
	t := 17

	var foundIndex int = -1

	// left and right boundaries of search
	var l, r int = 0, len(s) - 1

	for l <= r {
		m := int(math.Floor(float64(l)+float64(r)) / 2)

		if t == s[m] {
			foundIndex = m
			break
		}

		if t > s[m] {
			l = m + 1
			continue
		}

		if t < s[m] {
			r = m - 1
			continue
		}
	}

	fmt.Printf("value found at index %v\n", foundIndex)

}
