package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HMAC(value string) string {
	key := "BOOKOTO"
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
