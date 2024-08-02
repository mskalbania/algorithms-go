package week1

import (
	"fmt"
)

func Day1() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}

/*
Calculates ratio of positive, negative and zero values. Prints to 6 dec places.
https://www.hackerrank.com/challenges/three-month-preparation-kit-plus-minus/problem
*/
func plusMinus(arr []int32) {
	var positive int
	var negative int
	var zeros int
	for _, number := range arr {
		if number > 0 {
			positive++
			continue
		}
		if number < 0 {
			negative++
			continue
		}
		zeros++
	}
	fmt.Printf("%.6f\n", float32(positive)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(negative)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(zeros)/float32(len(arr)))
}
