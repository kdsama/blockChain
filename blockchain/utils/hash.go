package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// NewSHA256 ...
func NewSHA256(timestamp int64, lastHash string, data string) string {
	hash := sha256.Sum256([]byte(data))
	// fmt.Println(hash)
	toR := string(hash[:])
	return hex.EncodeToString([]byte(toR))
}
