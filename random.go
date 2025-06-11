package utilities

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	"!@#$%^&*()-_=+[]{}|;:,.<>?/"

func RandomInt(min, max int64) int64 {
	return min + mrand.Int63n(max-min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[mrand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomString(6))
}

func GenerateRandomPassword() (string, error) {
	// Random length between 10 and 16
	lengthInt, err := rand.Int(rand.Reader, big.NewInt(7)) // 0 to 6
	if err != nil {
		return "", err
	}
	length := 10 + int(lengthInt.Int64()) // 10 to 16

	password := make([]byte, length)
	for i := range password {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randomInt.Int64()]
	}
	return string(password), nil
}
