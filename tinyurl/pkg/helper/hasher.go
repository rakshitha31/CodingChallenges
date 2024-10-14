package helper

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateShortUrl(longUrl string) string {
	hasher := sha1.New()
	hasher.Write([]byte(longUrl))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha[:8]
}
