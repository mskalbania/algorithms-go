package week1

import "fmt"

func Day4() {
	fmt.Printf("%v", breakingRecords([]int32{12, 24, 10, 24}))
}

/*
Solution to:
https://www.hackerrank.com/challenges/three-month-preparation-kit-breaking-best-and-worst-records/problem
*/
func breakingRecords(scores []int32) []int32 {
	minScore, maxScore, timesMin, timesMax := scores[0], scores[0], int32(0), int32(0)
	for i := 1; i < len(scores); i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
			timesMax++
		}
		if scores[i] < minScore {
			minScore = scores[i]
			timesMin++
		}
	}
	return []int32{timesMax, timesMin}
}
