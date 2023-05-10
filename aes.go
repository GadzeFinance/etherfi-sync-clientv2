package main

import (
	"strings"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	//"bytes"
)

func Decrypt(encrypted_string string, ENCRYPTION_KEY string) ([]byte, error) {
	
	key := []byte(ENCRYPTION_KEY)
	parts := strings.Split(encrypted_string, ":")

	// the encrypted string should has the from [iv]:[ciphertext]
	if len(parts) != 2 {
		return nil, fmt.Errorf("ciphertext is not a multiple of 16")
	}

	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode iv:", err)
	}

	ciphertext, _ := hex.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext:", err)
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	if len(ciphertext) % aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of 16")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	// WARNING: didn't handle error from PKCS5UnPadding for now, maybe use same function from some packages
	ciphertext = PKCS5UnPadding(ciphertext)
	mode.CryptBlocks(ciphertext, ciphertext)
	
	return ciphertext, nil

}

// WARNING: NOT SURE IF THIS IS SECURE
// PKCS5UnPadding pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}