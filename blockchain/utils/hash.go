package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// NewSHA256 ...
func NewSHA256(timestamp int64, lastHash string, data string, nonce int64, difficulty int64) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d%s%s%d%d", timestamp, lastHash, data, nonce, difficulty)))
	// fmt.Println(hash)
	toR := string(hash[:])
	encodedString := hex.EncodeToString([]byte(toR))
	// fmt.Println(encodedString)
	return encodedString
}
