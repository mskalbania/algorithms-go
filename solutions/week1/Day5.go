package week1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// Day5 /*
/*
Solution to:
https://www.hackerrank.com/challenges/three-month-preparation-kit-camel-case/problem
*/
func Day5() {
	lines := readAllLines()
	for _, line := range lines {
		converted := convert(line[4:], line[2:3], line[0:1])
		fmt.Println(strings.Join(converted, " "))
	}
}

func readAllLines() []string {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, strings.TrimSpace(string(line)))
	}
	return lines
}

func convert(s string, objectType string, operation string) []string {
	switch operation {
	case "S":
		return split(s, objectType)
	case "C":
		return combine(s, objectType)
	}
	return []string{}
}

/*
Takes:
1. method signature string eg. - "printArray()" objectType = M
2. variable eg. - "colourMap" objectType = V
3. class eg. - "MyClass" objectType = C
- and returns array of words {"print", "array"}/{"colour", "map"}/{"my", "class"}
*/
func split(s string, objectType string) []string {
	if objectType == "M" {
		s = s[0 : len(s)-2] //if method strip "()"
	}
	var words []string
	var word string
	for _, charRune := range s {
		if unicode.IsUpper(charRune) {
			if word != "" { //don't append empty string when class type passed
				words = append(words, word)
			}
			word = ""
			word += strings.ToLower(string(charRune))
		} else {
			word += string(charRune)
		}
	}
	words = append(words, word) //append last
	return words
}

/*
Takes space separated list of words and:
1. combines to method signature eg. - "printArray()" when objectType = M
2. combines to variable eg. - "colourMap" when objectType = V
3. class eg. - "MyClass" when objectType = C
*/
func combine(s string, objectType string) []string {
	elements := strings.Split(s, " ")
	combined := ""
	for i, element := range elements {
		if i == 0 {
			if objectType == "C" { //when class type capitalize first string
				combined += capitalFirstLetter(element)
			} else {
				combined += element
			}
		} else {
			combined += capitalFirstLetter(element)
		}
	}
	if objectType == "M" { //when method type append curly-braces
		combined += "()"
	}
	return []string{combined}
}

func capitalFirstLetter(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
