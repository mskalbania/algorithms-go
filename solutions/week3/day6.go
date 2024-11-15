package week3

import (
	"math"
	"slices"
)

// https://www.hackerrank.com/challenges/three-month-preparation-kit-maximum-perimeter-triangle/problem
func maximumPerimeterTriangle(sticks []int32) []int32 {
	var trianglesFound [][]int32
	slices.Sort(sticks)
	for i := 0; i < len(sticks)-2; i++ {
		threeSticks := sticks[i : i+3]
		if canFormTriangle(threeSticks) {
			trianglesFound = append(trianglesFound, threeSticks)
		}
	}
	if len(trianglesFound) == 0 {
		return []int32{-1}
	}
	var maxTriangle []int32
	maxPerimeter := int32(math.MinInt32)
	for _, triangle := range trianglesFound {
		perimeter := triangle[0] + triangle[1] + triangle[2]
		if perimeter > maxPerimeter {
			maxTriangle = triangle
			maxPerimeter = perimeter
		} else if perimeter == maxPerimeter { // from req: if there are several valid triangles having the maximum perimeter
			if triangle[2] > maxTriangle[2] { // select the one with higher maximum side
				maxTriangle = triangle
			} else if triangle[2] == maxTriangle[2] { //or when higher max sides also same
				if maxTriangle[0] > triangle[0] { //take the one with lower min side
					maxTriangle = triangle
				}
			}
		}
	}
	return maxTriangle
}

func canFormTriangle(sticks []int32) bool {
	return sticks[0]+sticks[1] > sticks[2]
}
