package middleware

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"time"
)

// VerifyCode 为了考虑时间误差，判断前当前时间及前后30秒时间
func VerifyCode(secret string, code int32) bool {
	// 当前google值
	if getCode(secret, 0) == code {
		return true
	}
	// 前30秒google值
	if getCode(secret, -30) == code {
		return true
	}
	// 后30秒google值
	if getCode(secret, 30) == code {
		return true
	}
	return false
}

// 获取Google Code
func getCode(secret string, offset int64) int32 {
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	epochSeconds := time.Now().Unix() + offset
	return int32(oneTimePassword(key, toBytes(epochSeconds/30)))
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := toUint32(hashParts)
	pwd := number % 1000000
	return pwd
}
