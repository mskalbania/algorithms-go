package week4

// https://www.hackerrank.com/challenges/three-month-preparation-kit-array-left-rotation/problem
func rotateLeft(d int32, arr []int32) []int32 {
	for i := int32(0); i < d; i++ {
		temp := make([]int32, len(arr))
		temp[len(arr)-1] = arr[0]
		for i := 0; i < len(arr)-1; i++ {
			temp[i] = arr[i+1]
		}
		arr = temp
	}
	return arr
}
