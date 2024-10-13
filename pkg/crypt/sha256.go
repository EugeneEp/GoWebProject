package crypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(pass string) string {
	hash := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(hash[:])
}
