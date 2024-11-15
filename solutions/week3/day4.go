package week3

// https://www.hackerrank.com/challenges/three-month-preparation-kit-sock-merchant/problem
func sockMerchant(n int32, ar []int32) int32 {
	byColour := make(map[int32]int32)
	for _, sock := range ar {
		_, contains := byColour[sock]
		if contains {
			byColour[sock] = byColour[sock] + 1
		} else {
			byColour[sock] = 1
		}
	}
	pairs := int32(0)
	for _, v := range byColour {
		pairs += v / 2
	}
	return pairs
}
