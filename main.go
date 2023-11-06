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
	//key := "0123456789abcdef0123456789abcdef"
	key := "73646664666632337666333231326673"
	kind := "plain"
	data := `this is test payload with special characters - வணக்கம்`
	if len(os.Args) > 1 {
		kind = os.Args[1]
		if (kind != "plain" && kind != "encrypted") || len(os.Args) != 3 {
			fmt.Println("go run main.go [plain | encrypted] <content>")
			return
		}
		data = os.Args[2]
	}

	if kind == "plain" {
		encrypted := encrypt(data, []byte(key))
		fmt.Printf("Encrypted: %v\n", encrypted)

		decrypted := decrypt(encrypted, []byte(key))
		fmt.Printf("Decrypted: %v\n", decrypted)
		return
	}

	if kind == "encrypted" {
		decrypted := decrypt(data, []byte(key))
		fmt.Printf("Decrypted: %v\n", decrypted)
		return
	}
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

func encrypt(data string, key []byte) string {
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
	//return string(decrypted)

}
func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
