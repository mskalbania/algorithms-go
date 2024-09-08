package week3

import (
	"fmt"
)

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-the-birthday-bar/problem
*/
func Day2() {
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
}

func birthday(s []int32, d int32, m int32) int32 {
	var ways int32
	for i := 0; i < len(s); i++ {
		var segLen int32
		var sum int32
		for j := i; j < len(s); j++ {
			segLen++
			if segLen > m {
				break
			}
			sum += s[j]
			if sum == d && segLen == m {
				ways++
				break
			}
		}
	}
	return ways
}
