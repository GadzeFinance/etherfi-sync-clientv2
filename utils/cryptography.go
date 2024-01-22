package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/pbkdf2"
)

func fromString(str string) *big.Int {
	// Parse the input string as a decimal string
	// This imitates the wierd behavior of BN.js
	res := big.NewInt(0)
	for _, ch := range str {
		res.Mul(res, big.NewInt(10))
		val := new(big.Int)
		val.SetString(string(ch), 16)
		res.Add(res, val)
	}
	return res
}

func DecryptValidatorKeyInfo(file *schemas.IPFSResponseType, keypairForIndex schemas.KeyPair) schemas.ValidatorKeyInfo {
	// Fetch necessary data
	privateKey := keypairForIndex.PrivateKey
	encryptedValidatorKey := file.EncryptedValidatorKey
	encryptedKeystoreName := file.EncryptedKeystoreName
	encryptedPassword := file.EncryptedPassword
	stakerPublicKeyHex := file.StakerPublicKey

	// Get the staker's public key from its hex string
	bStakerPubKey, err := hex.DecodeString(stakerPublicKeyHex)
	if err != nil {
		panic(err)
	}

	// Get the staker's pubkey point from the public key
	receivedStakerPubKeyPoint, err := crypto.UnmarshalPubkey(bStakerPubKey)
	if err != nil {
		panic(err)
	}

	// Get the NO's private key
	nodeOperatorPrivKey := fromString(privateKey)
	// It seems that we need to mod this value to get the private key fit in to the curve library functions
	beMod, _ := big.NewInt(0).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	nodeOperatorPrivKey.Mod(nodeOperatorPrivKey, beMod)

	// Multiply the staker's pubKey point and NO's private key to generate shared secret
	curve := crypto.S256()
	nodeOperatorSharedSecret, _ := curve.ScalarMult(receivedStakerPubKeyPoint.X, receivedStakerPubKeyPoint.Y, nodeOperatorPrivKey.Bytes())

	// zero padded secret
	nodeOperatorSecret := make([]byte, 32)
	secretLen := len(nodeOperatorSharedSecret.Bytes())
	copy(nodeOperatorSecret[32-secretLen:], nodeOperatorSharedSecret.Bytes())

	// For compatibility, if all three encrypted fields are in the form [iv]:[data], we decrypt them using CBC mode
	isUsingCBC := false
	if len(strings.Split(encryptedKeystoreName, ":")) == 2 && len(strings.Split(encryptedValidatorKey, ":")) == 2 && len(strings.Split(encryptedPassword, ":")) == 2 {
		isUsingCBC = true
	}

	var bValidatorKey []byte
	var bValidatorKeyPassword []byte
	var bKeystoreName []byte

	// Use the shared secret to decrypt encrypted data
	if isUsingCBC {
		bValidatorKey, _ = DecryptCBC(encryptedValidatorKey, nodeOperatorSecret)
		bValidatorKeyPassword, _ = DecryptCBC(encryptedPassword, nodeOperatorSecret)
		bKeystoreName, _ = DecryptCBC(encryptedKeystoreName, nodeOperatorSecret)
	} else {
		bValidatorKey, _ = DecryptGCM(encryptedValidatorKey, nodeOperatorSecret)
		bValidatorKeyPassword, _ = DecryptGCM(encryptedPassword, nodeOperatorSecret)
		bKeystoreName, _ = DecryptGCM(encryptedKeystoreName, nodeOperatorSecret)
	}

	return schemas.ValidatorKeyInfo{
		ValidatorKeyFile:     bValidatorKey,
		ValidatorKeyPassword: bValidatorKeyPassword,
		KeystoreName:         bKeystoreName,
	}
}

func DecryptGCM(encrypted_string string, key []byte) ([]byte, error) {

	// There're three parts in the encrypted string: [iv]:[data]:[authTag]
	parts := strings.Split(encrypted_string, ":")
	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		panic(err.Error())
	}
	ciphertext, err := hex.DecodeString(parts[1] + parts[2])
	if err != nil {
		panic(err.Error())
	}
	// Create an AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Create a new GCM mode cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Decrypt using the cipher
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext, nil
}

func DecryptCBC(encrypted_string string, key []byte) ([]byte, error) {
	parts := strings.Split(encrypted_string, ":")
	// the encrypted string should has the from [iv]:[ciphertext]
	if len(parts) != 2 {
		return nil, fmt.Errorf("ciphertext is not a multiple of 16")
	}
	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode iv: %v", err)
	}
	ciphertext, _ := hex.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %v", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of 16")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)
	return ciphertext, nil
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func DecryptPrivateKeysGCM(privateKeys *schemas.KeyStoreFile, privKeyPassword string) (schemas.DecryptedDataJSON, error) {
	iv, err := hex.DecodeString(privateKeys.Iv)
	if err != nil {
		panic(err)
	}
	salt, err := hex.DecodeString(privateKeys.Salt)
	if err != nil {
		panic(err)
	}
	ciphertext, err := hex.DecodeString(privateKeys.Data + privateKeys.AuthTag)
	if err != nil {
		panic(err)
	}

	key := pbkdf2.Key([]byte(privKeyPassword), salt, 100000, 32, sha256.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	var decryptedDataJSON schemas.DecryptedDataJSON
	err = json.Unmarshal(plaintext, &decryptedDataJSON)
	if err != nil {
		panic(err)
	}

	return decryptedDataJSON, nil
}

func DecryptPrivateKeysCBC(privateKeys *schemas.KeyStoreFile, privKeyPassword string) (schemas.DecryptedDataJSON, error) {
	iv, err := hex.DecodeString(privateKeys.Iv)
	if err != nil {
		panic(err)
	}
	salt, err := hex.DecodeString(privateKeys.Salt)
	if err != nil {
		panic(err)
	}
	ciphertext, err := hex.DecodeString(privateKeys.Data)
	if err != nil {
		panic(err)
	}

	key := pbkdf2.Key([]byte(privKeyPassword), salt, 100000, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	decryptedData := PKCS5UnPadding(ciphertext)

	// fmt.Println("decrypted:", PrettyPrint(decryptedData))

	var decryptedDataJSON schemas.DecryptedDataJSON
	err = json.Unmarshal(decryptedData, &decryptedDataJSON)
	if err != nil {
		panic(err)
	}

	// fmt.Println("json:", decryptedDataJSON)

	return decryptedDataJSON, nil
}

func GetKeyPairByPubKeyIndex(pubkeyIndexString string, privateKeys []string, publicKeys []string) (schemas.KeyPair, error) {
	index, err := strconv.ParseInt(pubkeyIndexString, 10, 0)
	if err != nil {
		return schemas.KeyPair{}, err
	}
	return schemas.KeyPair{
		PrivateKey: privateKeys[index],
		PublicKey:  publicKeys[index],
	}, nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
