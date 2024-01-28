package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

func CryptBytes(hashType, salt string, value []byte) string {
	if hashType == "" {
		hashType = "SHA1"
	}
	if salt == "" {
		saltBytes := make([]byte, 16)
		rand.Read(saltBytes)
		salt = base64.URLEncoding.EncodeToString(saltBytes)
	}
	hash := sha1.New()
	hash.Write([]byte(salt))
	hash.Write(value)
	hashedBytes := hash.Sum(nil)
	result := fmt.Sprintf("$%s$%s$%s", hashType, salt, strings.TrimRight(base64.URLEncoding.EncodeToString(hashedBytes), "="))
	return result
}
