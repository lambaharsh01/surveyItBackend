package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/lambaharsh01/surveyItBackend/utils/constants"
)

func HashString(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CompareHashes(input string, hashed string) bool {
	return HashString(input) == hashed
}

func GenerateUniqueKey(keyLength int) string {
	const charset = constants.AlphanumericCharset
	var randomLength int = keyLength - 5

	result := make([]byte, randomLength)

	for i := 0; i < randomLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}

	timestamp := time.Now().UnixNano() % 100000 // 5 length

	var key string = fmt.Sprintf("%s%05d", string(result), timestamp) // 10 + 5

	return key
}
