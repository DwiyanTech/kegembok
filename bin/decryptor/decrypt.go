package main

/*
	For Educational Purposes Only
	Script by Dwiyan with Luv luv <3
*/

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dwiyantech/kegembok/utils"
)

func usage() {
	fmt.Printf("Usage : \n--target-dir [target_path_directory] (Mandatory)\n--key [string_key_toencrypt] (default: kegembok-cuy)\n\n")
}
func main() {
	fmt.Printf("Kegembok Ransomware Decryptor v1.0 by dwiyantech \n" +
		"Note : This Ransomware is for educational Purposes only\n" +
		"Author not responsible for any consequences that may occur.\n")
	pathToEncrypt := flag.String("targetdir", "", "Target to Decrypt Directory")
	key := flag.String("key", "kegembok-cuy", "Key used for to Encrypt Files")
	flag.Parse()
	if *pathToEncrypt == "" {
		fmt.Println(*pathToEncrypt)
		usage()
		os.Exit(2)
	}

	exploreDecryptDirectory(*pathToEncrypt, *key)

}

func exploreDecryptDirectory(pathDirectory string, key string) {

	err := filepath.Walk(pathDirectory, func(path string, info os.FileInfo, err error) error {
		return recursiveDecryptPath(path, key, info, err)
	})
	if err != nil {
		panic(err)
	}

}

func recursiveDecryptPath(path string, key string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Error visiting path:", err)
		return err
	}

	stringForWindowsExtensions := ".kegembok$"

	resultBool, err := regexp.MatchString(stringForWindowsExtensions, path)
	if err != nil {
		fmt.Println("This not File : ", path)
	}

	if resultBool {
		keynya := utils.GenerateKeyFromStr(key)
		enc, errBool := utils.DecryptFile(path, keynya)
		if !errBool {
			fmt.Println("Decrytping Files:", path)
			matchedExtensions := regexp.MustCompile(".kegembok$")
			realPath := matchedExtensions.ReplaceAllString(path, "")
			os.WriteFile(realPath, enc, 0666)
			os.Remove(path)
		}
	}

	return nil
}
