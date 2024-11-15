package week2

import "fmt"

func Day1() {
	fmt.Printf("%v", lonelyinteger([]int32{2, 2, 1, 3, 3, 4, 5, 4, 5}))
}

/*
Given an array of integers, where all elements but one occur twice, find the unique element.
https://www.hackerrank.com/challenges/three-month-preparation-kit-lonely-integer/problem
*/
func lonelyinteger(a []int32) int32 {
	numberToOccurrence := make(map[int32]int32)
	for _, e := range a {
		occurrence, exists := numberToOccurrence[e]
		if exists {
			numberToOccurrence[e] = occurrence + 1
		} else {
			numberToOccurrence[e] = 1
		}
	}
	for k, v := range numberToOccurrence {
		if v == 1 {
			return k
		}
	}
	return 0 //illegal state
}
