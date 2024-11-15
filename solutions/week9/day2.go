package week9

import "fmt"

func Day2() {
	fmt.Printf("%v", stockmax([]int32{1, 2, 100}))
}

func stockmax(prices []int32) int64 {
	var lastSellPrice int32
	var profit int64
	var stockAmount int32
	for i := len(prices) - 1; i >= 0; i-- {
		//last sell was higher so profit reduced
		if lastSellPrice >= prices[i] {
			profit -= int64(prices[i])
			stockAmount++
		} else { //otherwise sell
			profit += int64(stockAmount * lastSellPrice)
			stockAmount = 0
			lastSellPrice = prices[i]
		}
	}
	profit += int64(stockAmount * lastSellPrice)
	return profit
}
