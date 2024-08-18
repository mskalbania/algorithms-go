package week2

import (
	"fmt"
	"strings"
)

func Day7() {
	fmt.Printf("%v", pangrams("We promptly judged antique ivory buckles for the next prize"))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-pangrams/problem
*/
func pangrams(s string) string {
	letterMap := make(map[string]int8)
	for i := 97; i <= 122; i++ {
		letterMap[string(rune(i))] = 0
	}
	for _, c := range s {
		letterMap[strings.ToLower(string(c))]++
	}
	for _, v := range letterMap {
		if v == 0 {
			return notPangram
		}
	}
	return pangram
}

const (
	pangram    = "pangram"
	notPangram = "not pangram"
)
