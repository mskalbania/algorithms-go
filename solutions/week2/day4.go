package week2

import (
	"fmt"
	"math"
)

func Day4() {
	fmt.Printf("%v", diagonalDifference([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-diagonal-difference/problem
*/
func diagonalDifference(arr [][]int32) int32 {
	var diag1 int32
	var diag2 int32
	for i := range arr {
		for j := range arr[i] {
			diag1 = diag1 + arr[i][j+i]
			diag2 = diag2 + arr[i][len(arr)-1-i]
			break
		}
	}
	return int32(math.Abs(float64(diag1 - diag2)))
}
