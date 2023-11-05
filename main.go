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
	key := "0123456789abcdef"
	data := `this is test payload with special characters - வணக்கம்`
	if len(os.Args) > 1 {
		data = os.Args[1]
	}

	encrypted := encrypt(data, []byte(key))
	fmt.Printf("Encrypted: %v\n", encrypted)

	decrypted := decrypt(encrypted, []byte(key))
	fmt.Printf("Decrypted: %v\n", decrypted)
}

func generateIV() string {
	ivLength := 16
	lowAsciiLimit := 47
	highAsciiLimit := 126

	finalIvBuffer := make([]byte, ivLength)
	for i := 0; i < ivLength; i++ {
		randomNumber := lowAsciiLimit + rand.Intn(highAsciiLimit-lowAsciiLimit+1)
		finalIvBuffer[i] = byte(randomNumber)
	}
	return string(finalIvBuffer)
}

func encrypt(src string, key []byte) string {
	initialVector := generateIV()
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	content := []byte(src)
	content = pkcs5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	ecb.CryptBlocks(crypted, content)

	initialVectorBytes := []byte(initialVector)
	finalArray := make([]byte, len(initialVectorBytes) + len(crypted))
	copy(finalArray[:len(initialVectorBytes)], initialVectorBytes)
	copy(finalArray[len(initialVectorBytes):], crypted)

	return base64.StdEncoding.EncodeToString(finalArray)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func decrypt(encryptedBase64 string, key []byte) string {
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
func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
