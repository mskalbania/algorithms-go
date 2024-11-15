package week1

import "fmt"

func Day6() {
	fmt.Printf("%v", divisibleSumPairs(0, 3, []int32{1, 3, 2, 6, 1, 2}))
}

/*
Calculates how many sums of pair elements are dividable by k.
k = 3; arr = {1, 2, 3}
output = 1 (1, 2) pair (1+2)/3
https://www.hackerrank.com/challenges/three-month-preparation-kit-divisible-sum-pairs/problem
*/
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	divisable := int32(0)
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			pairSum := ar[i] + ar[j]
			if pairSum%k == 0 {
				divisable++
			}
		}
	}
	return divisable
}
