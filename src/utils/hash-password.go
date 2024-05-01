package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {

	hasher := sha256.New()

	hasher.Write([]byte(password))

	hashedBytes := hasher.Sum(nil)

	hashedPassword := hex.EncodeToString(hashedBytes)

	return hashedPassword
}
