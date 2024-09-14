package utils

import "crypto/sha256"

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return string(hasher.Sum(nil))
}
