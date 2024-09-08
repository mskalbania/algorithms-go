package week3

import (
	"bufio"
	"fmt"
	"os"
)

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-strings-xor/problem
*/
func Day3() {
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	var xor string
	for i := 0; i < len(input[0]); i++ {
		c1 := input[0][i]
		c2 := input[1][i]
		if c1 == c2 {
			xor += `0`
		} else {
			xor += `1`
		}
	}
	fmt.Println(xor)
}
