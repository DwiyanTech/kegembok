package utils

/*
	For Educational Purposes Only
	Script by Dwiyan with Luv luv <3
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func EncryptFile(file string, key []byte) ([]byte, bool) {
	errBool := false
	readFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error Skipping")
		errBool = true
	}
	block, err := aes.NewCipher(key)

	if err != nil {
		errBool = true
	}

	gcmAes, err := cipher.NewGCM(block)

	if err != nil {
		errBool = true
	}

	nonceSize := make([]byte, gcmAes.NonceSize())

	encryptedResult := gcmAes.Seal(nonceSize, nonceSize, readFile, nil)

	return encryptedResult, errBool

}
