package common

import (
	"crypto/md5"
)

func StrTo16Upper(s string) string {
	var alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hash := md5.Sum([]byte(s))
	var result string
	for i := 0; i < 16; i++ {
		index := int(hash[i]) % len(alphabet)
		result += string(alphabet[index])
	}
	return result
}
