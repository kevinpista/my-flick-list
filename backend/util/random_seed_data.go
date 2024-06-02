package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet ="abcdefghijklmnopqrstuvwxyz"

// Ensures everytime we run function, the generated values are different
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random integer between min & max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min +1)
}

// Generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates random name between length 4 and 24 characters
func RandomName() string {
	length := rand.Intn(24 - 4) + 4
	return RandomString(length)
}

/*
func RandomInt64() int64 {
	return RandomInt(0, 1000)
}
*/

