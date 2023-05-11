package main

import (
	"fmt"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey := "04653949f1185f5fc7cb441debe970c3049b7ee8e0a9d3d21d571f65d3c35fca8c3c678dfcf88afa23984596151c842f86c9856ae8c5cc6ff3f2812ce49888213a"
	publicKey := "74749212904585498199046374455021113593920534944642600239804505865032140419159"


	stakerPublicKeyHex := "04c4fec2cca2602f5d2359d52627b60f08a2368491310dc26c4fbca104a678f7fb6d353dfc5aa0d68359168551a4a0625d3788297c8b4742997d8a616cec0db338"
	bStakerPubKey, err := hex.DecodeString(stakerPublicKeyHex)
	if err != nil {
		panic(err)
	}

	receivedStakerPubKeyPoint, err := crypto.UnmarshalPubkey(bStakerPubKey)
	if err != nil {
		panic(err)
	}
	
	// fmt.Println(receivedStakerPubKeyPoint.X) [OK]

	bPrivateKey, err := hex.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}

	// for i, j := 0, len(bPrivateKey)-1; i < j; i, j = i+1, j-1 {
	// 	bPrivateKey[i], bPrivateKey[j] = bPrivateKey[j], bPrivateKey[i]
	// }

	fmt.Println("bPrivateKey:", bPrivateKey)

	nodeOperatorPrivKey := big.NewInt(0).SetBytes(bPrivateKey)

	fmt.Println("privKey:", nodeOperatorPrivKey.String())

	curve := crypto.S256()
	x, y := curve.ScalarMult(receivedStakerPubKeyPoint.X, receivedStakerPubKeyPoint.Y, nodeOperatorPrivKey.Bytes())

	fmt.Println(x, y)

	nodeOperatorSharedSecret := new(big.Int)
	nodeOperatorSharedSecret.Mul(receivedStakerPubKeyPoint.X, nodeOperatorPrivKey)

	fmt.Println(nodeOperatorSharedSecret)

	secretAsArray := nodeOperatorSharedSecret.Bytes()

	_ = publicKey

	fmt.Println(secretAsArray)

}