package week4

import (
	"math"
	"slices"
)

// https://www.hackerrank.com/challenges/three-month-preparation-kit-closest-numbers/problem
func closestNumbers(arr []int32) []int32 {
	slices.Sort(arr)
	var result []int32
	currLowest := int32(math.MaxInt32)
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff <= currLowest {
			if diff < currLowest {
				currLowest = diff
				result = nil
			}
			result = append(result, arr[i], arr[i+1])
		}
	}
	return result
}
