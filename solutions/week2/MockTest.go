package week2

import (
	"fmt"
	"math"
)

func MockTestTask1() {
	fmt.Println(flippingMatrix([][]int32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}))
}

/*
Sean invented a game involving a matrix where each cell of the matrix contains an integer.
He can reverse any of its rows or columns any number of times.
The goal of the game is to maximize the sum of the elements in the submatrix located in the upper-left quadrant of the matrix.
Given the initial configurations for matrices,
help Sean reverse the rows and columns of each matrix in the best possible way so that the sum of the elements in the matrix's upper-left quadrant is maximal.

Solution for 4x4 - find max 'a', find max 'b', find max 'c' and find max 'd'
abba
cddc
cddc
abba
*/
func flippingMatrix(matrix [][]int32) int32 {
	var sum int32
	for i := 0; i < len(matrix)/2; i++ { //since matrix is 2n x 2n look at half i and half j
		for j := 0; j < len(matrix)/2; j++ { //first iteration - look at a
			mx := math.Max(math.SmallestNonzeroFloat64, float64(matrix[i][j]))   //compare with 'a' left top
			mx = math.Max(mx, float64(matrix[i][len(matrix)-j-1]))               //compare with 'a' right top
			mx = math.Max(mx, float64(matrix[len(matrix)-i-1][j]))               //compare with 'a' left bottom
			mx = math.Max(mx, float64(matrix[len(matrix)-i-1][len(matrix)-j-1])) //compare with 'a' right bottom
			sum += int32(mx)
		}
	}
	return sum
}
