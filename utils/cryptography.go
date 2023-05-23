package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/pbkdf2"
)

func fromString(str string) *big.Int {
	// Parse the input string as a decimal string
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
	privateKey := keypairForIndex.PrivateKey
	encryptedValidatorKey := file.EncryptedValidatorKey
	encryptedKeystoreName := file.EncryptedKeystoreName
	encryptedPassword := file.EncryptedPassword

	stakerPublicKeyHex := file.StakerPublicKey

	bStakerPubKey, err := hex.DecodeString(stakerPublicKeyHex)
	if err != nil {
		panic(err)
	}

	receivedStakerPubKeyPoint, err := crypto.UnmarshalPubkey(bStakerPubKey)
	if err != nil {
		panic(err)
	}

	nodeOperatorPrivKey := fromString(privateKey)

	// Is this mod generic to use? because I didn't realy understand the math
	beMod, _ := big.NewInt(0).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	nodeOperatorPrivKey.Mod(nodeOperatorPrivKey, beMod)

	// fmt.Println(nodeOperatorPrivKey)

	curve := crypto.S256()
	nodeOperatorSharedSecret, _ := curve.ScalarMult(receivedStakerPubKeyPoint.X, receivedStakerPubKeyPoint.Y, nodeOperatorPrivKey.Bytes())


	secretAsArray := nodeOperatorSharedSecret.Bytes()

	// fmt.Println("secretAsArray:", len(secretAsArray), secretAsArray)

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

	if len(ciphertext)%aes.BlockSize != 0 {
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
	ciphertext, err := hex.DecodeString(privateKeys.Data)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	// TODO: we need to figure out how the last big of the cryptography works
	key := pbkdf2.Key([]byte(privKeyPassword), salt, 100000, 32, sha256.New)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// TODO: didn't handle error from PKCS5UnPadding for now, maybe use same function from some packages
	decryptedData := PKCS5UnPadding(ciphertext)

	var decryptedDataJSON schemas.DecryptedDataJSON
	err = json.Unmarshal(decryptedData, &decryptedDataJSON)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	// fmt.Println(PrettyPrint(decryptedDataJSON))

	return decryptedDataJSON, nil

}
