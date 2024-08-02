package week1

import (
	"fmt"
	"strconv"
)

func Day3() {
	result := timeConversion("13:00:01PM")
	fmt.Printf("%s", result)
}

/*
Converts 12h to 24h format
https://www.hackerrank.com/challenges/three-month-preparation-kit-time-conversion/problem
*/
func timeConversion(s string) string {
	suffix := s[len(s)-2:]
	time := s[0 : len(s)-2]
	hour, _ := strconv.Atoi(time[0:2])
	if suffix == "AM" {
		if hour == 12 {
			return "00" + time[2:]
		}
		return time
	} else {
		if hour == 12 {
			return time
		}
		return strconv.Itoa(hour+12) + time[2:]
	}
}
