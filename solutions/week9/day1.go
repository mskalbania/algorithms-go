package week9

func nextPrime(start int) int {
	if start == 1 {
		return 2
	}

	var numberToTest = start + 1
loop:
	for {
		for i := 2; i < numberToTest; i++ {
			if numberToTest%i == 0 {
				numberToTest++
				continue loop
			}
		}
		return numberToTest
	}
}
