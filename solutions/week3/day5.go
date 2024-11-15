package week3

import "math"

// https://www.hackerrank.com/challenges/three-month-preparation-kit-migratory-birds/problem
func migratoryBirds(arr []int32) int32 {
	birdsSpotted := make(map[int32]int32)
	for _, i := range arr {
		_, exists := birdsSpotted[i]
		if exists {
			birdsSpotted[i] = birdsSpotted[i] + 1
		} else {
			birdsSpotted[i] = 1
		}
	}

	var mostSpottedId int32
	mostSpottedTimes := int32(math.MinInt32)
	for k, v := range birdsSpotted {
		if v > mostSpottedTimes {
			mostSpottedTimes = v
			mostSpottedId = k
		}
		if v == mostSpottedTimes {
			mostSpottedId = int32(math.Min(float64(mostSpottedId), float64(k)))
		}
	}
	return mostSpottedId
}
