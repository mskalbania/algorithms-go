package week3

// https://www.hackerrank.com/challenges/between-two-sets/problem
func getTotalX(a []int32, b []int32) int32 {
	start := a[len(a)-1]
	end := b[len(b)-1]
	var counter = 0
outter:
	for i := start; i <= end; i++ {
		for _, ai := range a {
			if i%ai != 0 {
				continue outter
			}
		}
		for _, bi := range b {
			if bi%i != 0 {
				continue outter
			}
		}
		counter++
	}
	return int32(counter)
}
