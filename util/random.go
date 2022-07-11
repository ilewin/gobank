package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabetic = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetic)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabetic[rand.Intn(k)])
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return RandomInt(1, 100000)
}

func RandomCurrency() string {
	currenies := []string{"USD", "EUR", "PLN"}
	return currenies[rand.Intn(3)]
}
