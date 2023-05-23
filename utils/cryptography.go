package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
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

func DecryptValidatorKeyInfo(file schemas.IPFSResponseType, keypairForIndex schemas.KeyPair) schemas.ValidatorKeyInfo {
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
		return schemas.ValidatorKeyInfo{}
	}
	
	// Get the staker's pubkey point from the public key
	receivedStakerPubKeyPoint, err := crypto.UnmarshalPubkey(bStakerPubKey)
	if err != nil {
		panic(err)
		return schemas.ValidatorKeyInfo{}
	}
	
	// Get the NO's private key
	nodeOperatorPrivKey := fromString(privateKey)
	// It seems that we need to mod this value to get the private key fit in to the curve library functions
	beMod, _ := big.NewInt(0).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	nodeOperatorPrivKey.Mod(nodeOperatorPrivKey, beMod)

	// Multiply the staker's pubKey point and NO's private key to generate shared secret
	curve := crypto.S256()
	nodeOperatorSharedSecret, _ := curve.ScalarMult(receivedStakerPubKeyPoint.X, receivedStakerPubKeyPoint.Y, nodeOperatorPrivKey.Bytes())
	secretAsArray := nodeOperatorSharedSecret.Bytes()

	// Use the shared secret to decrypt encrypted data
	bValidatorKey, _ := Decrypt(encryptedValidatorKey, hex.EncodeToString(secretAsArray))
	bValidatorKeyPassword, _ := Decrypt(encryptedPassword, hex.EncodeToString(secretAsArray))
	bKeystoreName, _ := Decrypt(encryptedKeystoreName, hex.EncodeToString(secretAsArray))

	return schemas.ValidatorKeyInfo{
		ValidatorKeyFile:     bValidatorKey,
		ValidatorKeyPassword: bValidatorKeyPassword,
		KeystoreName:         bKeystoreName,
	}
}


func Decrypt(encrypted_string string, ENCRYPTION_KEY string) ([]byte, error) {
	// Decode the encryption key
	key, err := hex.DecodeString(ENCRYPTION_KEY)

	// There're three parts in the encrypted string: [iv]:[data]:[authTag]
	parts := strings.Split(encrypted_string, ":")
	iv, err := hex.DecodeString(parts[0])
	ciphertext, err := hex.DecodeString(parts[1] + parts[2])

	// Create an AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
		return []byte{}, err
	}

	// Create a new GCM mode cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
		return []byte{}, err
	}

	// Decrypt using the cipher
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		panic(err.Error())
		return []byte{}, err
	}

	return plaintext, nil
}

func DecryptPrivateKeys(privateKeys schemas.KeyStoreFile, privKeyPassword string) (schemas.DecryptedDataJSON, error) {
	iv, err := hex.DecodeString(privateKeys.Iv)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}
	salt, err := hex.DecodeString(privateKeys.Salt)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}
	ciphertext, err := hex.DecodeString(privateKeys.Data + privateKeys.AuthTag)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	key := pbkdf2.Key([]byte(privKeyPassword), salt, 100000, 32, sha256.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
		return schemas.DecryptedDataJSON{}, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
		return schemas.DecryptedDataJSON{}, err
	}
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	var decryptedDataJSON schemas.DecryptedDataJSON
	err = json.Unmarshal(plaintext, &decryptedDataJSON)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	return decryptedDataJSON, nil
}
