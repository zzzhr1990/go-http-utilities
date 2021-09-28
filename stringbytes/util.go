package stringbytes

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

var letterRunes = []rune("abcdef01234567890")

func ComputeMd5(data []byte) string {
	h := md5.Sum(data)
	return hex.EncodeToString(h[:])
}
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandStringRunesWithLetters(n int, letters []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func ComputeSha1(data []byte) string {
	h := sha1.Sum(data)
	return hex.EncodeToString(h[:])
}
