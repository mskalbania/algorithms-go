package week2

import "fmt"

func Day6() {
	fmt.Printf("%v", countingValleys(0, "UDDDUDUU"))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-counting-valleys/problem
*/
func countingValleys(steps int32, path string) int32 {
	var valleys int32
	var relativePosition int32
	for _, e := range path {
		step := string(e)
		if step == "U" {
			relativePosition++
			if relativePosition == 0 { //at relative 0 now, was below before - valley ended
				valleys++
			}
		} else {
			relativePosition--
		}
	}
	return valleys
}
