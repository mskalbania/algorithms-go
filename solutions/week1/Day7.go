package week1

import (
	"fmt"
	"runtime"
)

func Day7() {
	fmt.Printf("%v", matchingStrings([]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}))
}

/*
Takes a slice of strings and returns the occurrence of strings based on queries
strings - {"a","b"}, queries - {"a","c"}, returns - {1,0} 1 * 'a' | 0 * 'c'
https://www.hackerrank.com/challenges/three-month-preparation-kit-sparse-arrays/problem
*/
func matchingStrings(strings []string, queries []string) []int32 {
	stringOccurrence := make(map[string]int32)

	for _, s := range strings {
		occurrence, exists := stringOccurrence[s]
		if exists {
			stringOccurrence[s] = occurrence + 1
		} else {
			stringOccurrence[s] = 1
		}
	}
	result := make([]int32, 0, len(queries))
	for _, query := range queries {
		occurrence, exists := stringOccurrence[query]
		if exists {
			result = append(result, occurrence)
		} else {
			result = append(result, 0)
		}
	}
	runtime.Gosched()
	return result
}
