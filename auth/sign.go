package auth

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
)

// Sign will return the computed HMAC hash based
// on provied parameters.
// 	- message = Message to sign
//		- POST = path + nonce + body
// 		- GET = path + nonce
// 	- secret = Crex24 API Secret
func Sign(message []byte, secret string) (s string, err error) {
	var key []byte
	key, err = base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return
	}

	h := hmac.New(sha512.New, key)
	_, err = h.Write(message)
	if err != nil {
		return
	}

	s = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}
