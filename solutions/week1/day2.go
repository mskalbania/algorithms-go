package week1

import "fmt"

func Day2() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}

/*
Given fives positive integers,
finds the minimum and maximum values that can be calculated by summing exactly four of the five integers.
https://www.hackerrank.com/challenges/three-month-preparation-kit-mini-max-sum/problem
*/
func miniMaxSum(arr []int32) {
	var sumOfAll int64
	for _, e := range arr {
		sumOfAll += int64(e)
	}
	var maxVal int64
	minVal := sumOfAll
	for _, e := range arr {
		sumWithout := sumOfAll - int64(e)
		if sumWithout > maxVal {
			maxVal = sumWithout
		}
		if sumWithout < minVal {
			minVal = sumWithout
		}
	}
	fmt.Printf("%d %d", minVal, maxVal)
}
