package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

func GetCryptedBytes(hashType, salt string, value []byte) string {
	hash := sha1.New()
	hash.Write([]byte(salt))
	hash.Write(value)
	hashedBytes := hash.Sum(nil)
	return strings.TrimRight(base64.URLEncoding.EncodeToString(hashedBytes), "=")
}
