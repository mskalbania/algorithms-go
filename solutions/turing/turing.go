package main

import (
	"math"
	"math/rand/v2"
	"strings"
)

//https://www.turing.com/interview-questions/golang#advanced-golang-interview-questions-and-answers

func sumEvenFrom1To100() int {
	var sum int
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func smallestAndLargest(arr []int) (int, int) {
	smallest := math.MaxInt32
	largest := math.MinInt32
	for _, i := range arr {
		if i < smallest {
			smallest = i
		}
		if i > largest {
			largest = i
		}
	}
	return smallest, largest
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func factorial(num uint64) uint64 {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func numWords(s string) int {
	return len(strings.Split(s, " "))
}

func random(a, b int) int {
	return rand.IntN(b-a+1) + a
}
