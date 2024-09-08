package week3

import (
	"fmt"
	"sort"
)

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-two-arrays/problem
*/
func Day1() {
	fmt.Println(twoArrays(10, []int32{2, 1, 3}, []int32{7, 8, 9}))
}

func twoArrays(k int32, A []int32, B []int32) string {
	//Hackerrank doesn't support below go features
	//slices.Sort(A)                       //'A' in natural order
	//slices.SortFunc(B, intReversedOrder) //sort 'B' in reverse order

	//Using old API
	AInt := make([]int, len(A))
	BInt := make([]int, len(A)) //same len
	for i := 0; i < len(A); i++ {
		AInt[i] = int(A[i])
		BInt[i] = int(B[i])
	}
	sort.Ints(AInt)
	sort.Sort(sort.Reverse(sort.IntSlice(BInt)))

	for i, e := range AInt {
		if e+BInt[i] < int(k) {
			return "NO"
		}
	}
	return "YES"
}

func intReversedOrder(i, j int32) int {
	if i < j {
		return 1
	} else if i > j {
		return -1
	}
	return 0
}
