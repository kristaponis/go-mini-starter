package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
)

// HMACHashString is used to hash a string. This func is used
// on already existing string, ex. to create remember_hash token
// from remember token, to be stored in the database.
func HMACHashString(s string) string {
	h := hmac.New(sha256.New, []byte(os.Getenv("HMAC_KEY")))
	h.Write([]byte(s))
	b := h.Sum(nil)

	return base64.URLEncoding.EncodeToString(b)
}
