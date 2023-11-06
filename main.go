package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	keyStr := "0123456789abcdef0123456789abcdef"

	if len(os.Args) == 1 {
		plain := "hello"
		encrypted := encrypt(plain, keyStr)
		decrypted := decrypt(encrypted, keyStr)
		fmt.Printf("Plain: %s\nEncrypted: %s\nDecrypted: %s\n", plain, encrypted, decrypted)
		return
	}

	if len(os.Args) != 3 {
		fmt.Println(`go run main.go <encrypt | decrypt> "<content>"`)
		return	
	}

	kind := os.Args[1]
	if kind == "encrypt" {
		fmt.Println(encrypt(os.Args[2], keyStr))
		return
	}

	if kind == "decrypt" {
		fmt.Println(decrypt(os.Args[2], keyStr))
		return
	}
	
	fmt.Println(`go run main.go <encrypt | decrypt> "<content>"`)
}


func encrypt(data string, keyStr string) string {
	key := []byte(keyStr)
	initialVector := generateIV()
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	content := []byte(data)
	content = pkcs5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb := cipher.NewCBCEncrypter(block, initialVector)
	ecb.CryptBlocks(crypted, content)

	finalArray := make([]byte, len(initialVector)+len(crypted))
	copy(finalArray[:len(initialVector)], initialVector)
	copy(finalArray[len(initialVector):], crypted)

	return base64.StdEncoding.EncodeToString(finalArray)
}

func decrypt(encryptedBase64 string, keyStr string) string {
	key := []byte(keyStr)
	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		panic(err)
	}
	
	keySpec, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ivBytes := make([]byte, 16)
	content := make([]byte, len(encrypted)-len(ivBytes))
	copy(ivBytes, encrypted[:len(ivBytes)])
	copy(content, encrypted[len(ivBytes):])
	ecb := cipher.NewCBCDecrypter(keySpec, ivBytes)
	decrypted := make([]byte, len(content))
	ecb.CryptBlocks(decrypted, content)

	return string(pkcs5Trimming(decrypted))
}

func generateIV() []byte {
	ivLength := 16
	lowAsciiLimit := 47
	highAsciiLimit := 126

	ivBuffer := make([]byte, ivLength)
	for i := 0; i < ivLength; i++ {
		randomNumber := lowAsciiLimit + rand.Intn(highAsciiLimit-lowAsciiLimit+1)
		ivBuffer[i] = byte(randomNumber)
	}
	return ivBuffer
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
