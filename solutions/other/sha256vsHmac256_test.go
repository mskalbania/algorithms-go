package other

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

var Result string

func BenchmarkRegularSHA256(b *testing.B) {
	var r string
	testOtp := []byte("123456")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hasher := sha256.New()
		hasher.Write(testOtp)
		r = base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	}
	Result = r // prevent compiler optimization
}

func BenchmarkHMACSHA256(b *testing.B) {
	var r string
	testOtp := []byte("123456")
	secret := []byte("some-random-secret-key-for-testing-purposes")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h := hmac.New(sha256.New, secret)
		h.Write(testOtp)
		r = base64.StdEncoding.EncodeToString(h.Sum(nil))
	}
	Result = r // prevent compiler optimization
}
