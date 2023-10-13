package helpers

import (
	"crypto/sha256"
	"fmt"
)

// Hashes password string using SHA256
func HashPwd(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	hex_pwd := h.Sum(nil)
	hashed_pwd := fmt.Sprintf("%x", hex_pwd)
	return hashed_pwd
}

// Compares SHA256 hashes
func CompareHashes(pwd string, hash string) bool {
	return HashPwd(pwd) == hash
}
