package week3

import "math"

// https://www.hackerrank.com/challenges/three-month-preparation-kit-drawing-book/problem
// NOTE: this is iteration/pragmatic possibly suboptimal approach, IDK what is mathematical relation to calculate it in 1 line
func pageCount(n int32, p int32) int32 {
	if p == 1 || p == n { //special case when no swaps required
		return 0
	}

	frontNum := int32(math.MaxInt32)
	backNum := int32(math.MaxInt32)

	//we will swap pages in loop those are used to track what pages are currently seen
	frontPage1 := int32(0)
	frontPage2 := int32(1)
	backPage1 := n
	var backPage2 int32
	if n%2 == 0 { //last page placement is different based on (un)even number of pages
		backPage2 = n + 1
	} else {
		backPage2 = n - 1
	}

	//first going from front
	for i := 0; i <= int(n)/2; i++ {
		if p == frontPage1 || p == frontPage2 {
			frontNum = int32(i)
			break
		} else {
			frontPage1 += 2
			frontPage2 += 2
		}
	}
	//first going from back
	for i := 0; i <= int(n)/2; i++ {
		if p == backPage1 || p == backPage2 {
			backNum = int32(i)
			break
		} else {
			backPage1 -= 2
			backPage2 -= 2
		}
	}
	return int32(math.Min(float64(frontNum), float64(backNum)))
}
