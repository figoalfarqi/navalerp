package helper

import (
	"crypto/rand"
	"math/big"
)

// GenerateNDigitString menghasilkan 6 angka random dalam bentuk string
func GenerateNDigitString(n int) (string, error) {
	const digits = "0123456789"
	result := make([]byte, n)

	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		result[i] = digits[n.Int64()]
	}

	return string(result), nil
}
