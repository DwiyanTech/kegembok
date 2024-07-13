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

func DecryptFile(file string, key []byte) ([]byte, bool) {
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

	nonceSize := readFile[:gcmAes.NonceSize()]
	encryptedFiles := readFile[gcmAes.NonceSize():]

	encryptedResult, err := gcmAes.Open(nil, nonceSize, encryptedFiles, nil)
	if err != nil {
		errBool = true
	}
	return encryptedResult, errBool

}
