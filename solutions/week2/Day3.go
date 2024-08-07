package week2

import (
	"fmt"
	"strconv"
)

func Day3() {
	fmt.Printf("%v", flippingBits(1))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-flipping-bits/problem
*/
func flippingBits(n int64) int64 {
	binary := fmt.Sprintf("%032b", n)
	convertedBinary := ""
	for _, c := range binary {
		char := string(c)
		if char == "0" {
			convertedBinary = convertedBinary + "1"
		} else {
			convertedBinary = convertedBinary + "0"
		}
	}
	parsed, _ := strconv.ParseUint(convertedBinary, 2, 32)
	return int64(parsed)
}
