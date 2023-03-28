package helpers

import (
	"crypto/rand"
	"math/big"
)

func RandCharacter(number int) string {
	// String charset
	charset := "1234567890"

	// Getting random character
	var stringRand string

	for i := 1; i <= number; i++ {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		// number := big.NewInt(randNum.Int64())
		stringRand = stringRand + string(charset[randNum.Int64()])
	}

	return stringRand
}
