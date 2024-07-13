package utils

/*
	For Educational Purposes Only
	Script by Dwiyan with Luv luv <3
*/

import (
	"crypto/aes"
	"crypto/sha256"
)

func GenerateKeyFromStr(keyString string) []byte {
	hash := sha256.Sum256([]byte(keyString))
	key := hash[:aes.BlockSize]
	return key
}
