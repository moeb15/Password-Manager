package helpers

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
)

// Hashes password string using SHA256
func HashPwd(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	hashed_pwd := hex.EncodeToString(h.Sum(nil))
	return hashed_pwd
}

// Compares SHA256 hashes
func CompareHashes(pwd string, hash string) bool {
	return HashPwd(pwd) == hash
}

// Encrypts plaintext message using AES
func EncryptAES(key []byte, msg string) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	out := make([]byte, len(msg))
	c.Encrypt(out, []byte(msg))
	return hex.EncodeToString(out), nil
}

// Decrypts ciphertexts generated using AES
func DecryptAES(key []byte, cphr_txt string) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	out := make([]byte, len(cphr_txt))
	c.Decrypt(out, []byte(cphr_txt))
	msg := string(out[:])
	return msg, nil
}
