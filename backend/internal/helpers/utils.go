package helpers

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
)

const key_size = 32

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
	new_key := PadOrTrim(key, key_size)
	padded_msg := PadOrTrim([]byte(msg), key_size)
	c, err := aes.NewCipher(new_key)
	if err != nil {
		return "", err
	}

	out := make([]byte, key_size)
	c.Encrypt(out, padded_msg)
	return hex.EncodeToString(out), nil
}

// Decrypts ciphertexts generated using AES
func DecryptAES(key []byte, cphr_txt string) (string, error) {
	hex_cpher, _ := hex.DecodeString(cphr_txt)
	new_key := PadOrTrim(key, key_size)
	c, err := aes.NewCipher(new_key)
	if err != nil {
		return "", err
	}
	out := make([]byte, key_size)
	c.Decrypt(out, hex_cpher)
	msg := string(out[:])
	return msg, nil
}

// Pad or trim key to the required key size
func PadOrTrim(key []byte, size int) []byte {
	if len(key) == size {
		return key
	} else if len(key) > size {
		return key[:size+1]
	} else {
		zero_padding := make([]byte, size-len(key))
		for k := range zero_padding {
			zero_padding[k] = 0
		}
		key = append(key, zero_padding...)
		return key
	}
}
