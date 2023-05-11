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

// func main() {

// 	encrypted_string := "417f4f4f430a0c67c2886afc20e11238:7acdc9a354518e5aa191803435f96dbad0177cb939e9290f34c5b4d1087bb2cc1c5b323ba7083c4875684fc375e183e3b42202830c6020cc257b91b6d8bcebb0c09bc1ffab982e77ea67449452c96af2dcf3cfc3909c66a0de62d463ab056f5c65e7dc22259e689357b4ace94098f3c38aa6bd3410c7bcbd12b94a5cca2aa578f83bd11cf79b3e8538c016fe335bf676a13f15a2f046d29c83a2e93e34fb49f21d6c625748f510301e1a930defc010ca369d6089639aae7328ba43134bf18d8d8517d1268d386098be71125e9fd0752f86cc03bbe67729e1008e87591c1eeee8d5b1209e4a1c77039ec0586a77be1ad8ea92173bedd66c0a15d84c6235b56fa1fb2f9c02dfb69e472f276dd957fca7fbe4ef519290ca48bf3c750707672998b11fe4c7de0ee9b7180c0a466a26f0b8ca9cf0625ea1154096501a8db1c65850c4e72d43c74e49d8a6f29d595fd5b685538664f3dcf7d938665cc2e12665c5934101e3a6ac85b8a72e12ef58343efc22b81492bba4f1964119879c9dad7904b37aef73c6c5cdce0000ffc45c70cb98364d130da7048c6092d5d62c08a14d86753223e42ea6969dbda17389c928b94ccbd0476092e1bdd3b0b3e4f8adbe33080421882bb3be876163f9987b3808d37c7bf3f4ba19081ff2b0bbc4557163dc81aae4790ee0df8d06a7a4b58a36b7200a61a73b9b75b75a3a901a875f534a52703ce2dd3247993931f234eec52d94fffc0d9da710baab4c35ff3de36f5300891a4cde9dd180cbfa90c6489c5480e89573489948cd780f64c2900e3c32586ee0088e14a0fcbbbb5d4ce99c3eae8f54bbfbe9d131390836df97985b018f0241274961893e6637671d6b348414391bb3b2f93bb42c59eafb36b8d3dde992acdd792c40544cf1cf9172503e869337d7afa7421a7b2b2ee38fd62adbbdd0c3ebf792a9891c"
// 	ENCRYPTION_KEY := ""
// }