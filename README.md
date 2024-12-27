# Kegembok Ransomware

<p align="center">

<img src="/assets/img/kegembok_logo.png">

</p>

Kegembok is a Ransomware tools, a cross-platform (Linux, Mac, and Windows) ransomware made from the Golang programming language, encryption using AES-256-GCM, in this program you can use your own key.
the program looklike :
```
Kegembok Ransomware Encryptor v1.0 by dwiyantech 
Note : This Ransomware is for educational Purposes only
Author not responsible for any consequences that may occur.

Usage : 
--target-dir [target_path_directory] (Mandatory)
--key [string_key_toencrypt] (default: kegembok-cuy)
```
## Package Used
All package used go core package, so no need additionals download package for this program, for further more visit golang docs : https://go.dev/doc/
## How to Use
This Compiled from Linux for further informations in golang documantation  
#### 1. Clone the project First

```
  git clone https://github.com/dwiyantech/kegembok
```
#### 2. Map into cloned Directory Compile using go to build for
Command above is for windows
```bash 
  # For Encryptor 
  GOOS=windows GOARCH=amd64 go build ./bin/encryptor/encrypt.go
  # For Decryptor 
  GOOS=windows GOARCH=amd64 go build ./bin/decryptor/decrypt.go
```
Command above is for Linux
```
# For Encryptor 
go build ./bin/encryptor/encrypt.go
# For Decryptor 
go build ./bin/decryptor/decrypt.go
```
Command above is for Darwin based (Mac OS)

```
# For Encryptor 
env GOOS=darwin GOARCH=amd64 go build ./bin/encryptor/encrypt.go
# For Decryptor 
env GOOS=darwin GOARCH=amd64 go build ./bin/decryptor/decrypt.go
```
#### 3. Run The progam
Linux and Darwin (Mac OS)
```
encrypt -targetdir="Path Files" -key="Your Key String"
decrypt -targetdir="Path Files" -key="Your Key String"
```
Windows
```
encrypt.exe -targetdir="Path Files" -key="Your Key String"
decrypt.exe -targetdir="Path Files" -key="Your Key String"
```

## Contributing

Contributions are always welcome!, you can submit PR for bug fix or additionals feature

## Feedback

If you have any feedback, please make issues on this project


## Disclaimer
Kegembok is developed for research purposes only. The creators of Kegembok are not responsible for any misuse of this program. This program should not be used in any unauthorized or illegal manner. Always ensure ethical and legal use of this program.

