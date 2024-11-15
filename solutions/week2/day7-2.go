package week2

import "fmt"

func Day7x2() {
	fmt.Printf("%v", marsExploration("SOSSOT"))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-mars-exploration/problem
*/
func marsExploration(s string) int32 {
	var altered int32
	for i := 0; i < len(s); i = i + 3 {
		if string(s[i]) != "S" {
			altered++
		}
		if string(s[i+1]) != "O" {
			altered++
		}
		if string(s[i+2]) != "S" {
			altered++
		}
	}
	return altered
}
