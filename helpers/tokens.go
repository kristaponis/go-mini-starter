package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

// RandomBytes generates crypto random slice of bytes to be used
// in the generation of remember tokens or other strings.
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// RememberToken returns random string, randomized with provided
// int as number of bytes. Recommendation is from 32 bytes.
func RememberToken(n int) (string, error) {
	b, err := RandomBytes(n)
	if err != nil {
		return "", nil
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
