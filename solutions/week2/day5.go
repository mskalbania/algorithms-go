package week2

import "fmt"

func Day5() {
	fmt.Printf("%v", countingSort([]int32{1, 1, 3, 2, 1}))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-countingsort1/problem
*/
func countingSort(arr []int32) []int32 {
	frequencyArray := make([]int32, 100)
	for _, e := range arr {
		frequencyArray[e] = frequencyArray[e] + 1
	}
	return frequencyArray
}
