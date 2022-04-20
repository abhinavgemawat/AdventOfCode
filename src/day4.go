package src

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func Day4(secretKey string) (int, error) {
	result := 0
	for {
		r := strconv.FormatInt(int64(result), 10)
		str := secretKey + r
		md5Result := GetMD5Hash(str)
		if md5Result[0:6] == "000000" {
			break
		}
		result++
	}

	return result, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
