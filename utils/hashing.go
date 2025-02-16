package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(input string) string {
    hasher := sha256.New()
    hasher.Write([]byte(input))
    return hex.EncodeToString(hasher.Sum(nil))
}

func CompareHashes(input string, hashed string) bool {
    return HashString(input) == hashed
}
