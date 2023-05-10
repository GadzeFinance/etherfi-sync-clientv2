package main

import (
	"strings"
	"crypto/aes"
	"crypto/cipher"
	//"crypto/rand"
	"encoding/hex"
	"fmt"
	"bytes"
	//"io"
)

func encrypt(plaintext string, ENCRYPTION_KEY string) string {

	// convert to bytes
	key := []byte(ENCRYPTION_KEY)
	text := []byte(plaintext)
	bPlaintext := PKCS5Padding(text, aes.BlockSize, len(plaintext))
	

	// Create a new aes cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, len(bPlaintext))
	iv, _ := hex.DecodeString("4fa9077c8f4ae788988bb8cae303bf53")
	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	panic(err)
	// }

	//test_iv, _ := hex.DecodeString("4fa9077c8f4ae788988bb8cae303bf53")

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)


	hex_string := hex.EncodeToString(iv) + ":" + hex.EncodeToString(ciphertext[:aes.BlockSize]) + hex.EncodeToString(ciphertext[aes.BlockSize:])
	
	fmt.Println(hex_string)

	return hex_string

}

// WARNING: NOT SURE IF THIS IS SECURE
func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func decrypt(encrypted_string string, ENCRYPTION_KEY string) (string, error) {
	
	key := []byte(ENCRYPTION_KEY)

	parts := strings.Split(encrypted_string, ":")

	// fmt.Println("iv:", parts[0])
	// TODO: check if parts has two
	iv, _ := hex.DecodeString(parts[0])
	ciphertext, _ := hex.DecodeString(parts[1])

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
		return "", err
	}

	if len(ciphertext) % aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of 16")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	ciphertext = PKCS5UnPadding(ciphertext)
	mode.CryptBlocks(ciphertext, ciphertext)
	

	fmt.Println(string(ciphertext))
	
	return hex.EncodeToString(ciphertext), nil

}

// WARNING: NOT SURE IF THIS IS SECURE
// PKCS5UnPadding pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func main() {

	key := "6368616e676520746869732070617373"
	text := "exampleplaintext"

	en_string := encrypt(text, key)

	decrypt(en_string, key)

}