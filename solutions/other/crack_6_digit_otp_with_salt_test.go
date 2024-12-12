package other

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
	"time"
)

func TestCrackSixDigitSaltedCode(t *testing.T) {
	expectedCode := 259374

	exposedSalt := "hG4kP2tXhG4kP2tXhG4kP2tXhG4kP2tXhG4kP2tXhG4kP2tXhG4kP2tXhG4kP2tX"
	exposedHash := "dee3901b3b5483b99d99773ac785a1767551f0342fee6dc03ec3de879aa78da9"

	startTime := time.Now()

	for i := 100_000; i <= 999_999; i++ {
		hasher := sha256.New()
		hasher.Write([]byte(strconv.Itoa(i) + exposedSalt))
		hexS := hex.EncodeToString(hasher.Sum(nil))
		if hexS == exposedHash {
			t.Logf("Found the code: %d, hex: %s\n", i, hexS)
			t.Logf("Expected code: %d\n", expectedCode)
			t.Logf("Time taken: %v\n", time.Now().Sub(startTime))
		}
	}
}
