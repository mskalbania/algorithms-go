package week1

import (
	"fmt"
	"sort"
)

func MockTestTask1() {
	fmt.Printf("%v", findMedian([]int32{1, 8, 3}))
}

func findMedian(arr []int32) int32 {
	converted := make([]int, 0, len(arr))
	for _, e := range arr {
		converted = append(converted, int(e))
	}
	sort.Ints(converted)
	return int32(converted[len(converted)/2])
}
