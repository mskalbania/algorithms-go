package week4

import (
	"math"
	"slices"
)

// https://www.hackerrank.com/challenges/three-month-preparation-kit-picking-numbers/problem
func pickingNumbers(a []int32) int32 {
	slices.Sort(a)
	maxLen := int32(math.MinInt32)

	currentLen := 1
	var currentDiff int32
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a)-1; j++ {
			a1 := a[j]
			a2 := a[j+1]
			if a1 == a2 {
				currentLen++
				continue
			}
			if a2-a1 == 1 && currentDiff != 1 {
				currentDiff++
				currentLen++
				continue
			}
			break
		}
		maxLen = int32(math.Max(float64(maxLen), float64(currentLen)))
		currentLen = 1
		currentDiff = 0
	}
	return maxLen
}
