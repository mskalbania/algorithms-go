package week2

import "fmt"

func Day2() {
	fmt.Printf("%v", gradingStudents([]int32{73, 67, 38, 33}))
}

/*
https://www.hackerrank.com/challenges/three-month-preparation-kit-grading/problem
*/
func gradingStudents(grades []int32) []int32 {
	gradesAfterRounding := make([]int32, 0, len(grades))
outerLoop:
	for _, grade := range grades {
		//1st rule - if already dividable by 5 OR if it doesn't meet 38>= requirement add
		if grade%5 == 0 || grade < 38 {
			gradesAfterRounding = append(gradesAfterRounding, grade)
			continue
		}
		for i := 1; i <= 3; i++ { //2nd rule - if close to div by 5 in at most 3 - round to next div by 5
			increasedGrade := grade + int32(i)
			if increasedGrade%5 == 0 {
				gradesAfterRounding = append(gradesAfterRounding, increasedGrade)
				continue outerLoop
			}
		}
		gradesAfterRounding = append(gradesAfterRounding, grade)
	}
	return gradesAfterRounding
}
