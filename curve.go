package main

import (
	"fmt"
	"encoding/hex"
	"math/big"
	// "crypto/elliptic"
	//"github.com/btcsuite/btcd/btcec"
	// "github.com/coinbase/kryptology/pkg/core/curves"

	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

func main() {
	privateKey := "04653949f1185f5fc7cb441debe970c3049b7ee8e0a9d3d21d571f65d3c35fca8c3c678dfcf88afa23984596151c842f86c9856ae8c5cc6ff3f2812ce49888213a"
	publicKey := "74749212904585498199046374455021113593920534944642600239804505865032140419159"
	// encryptedValidatorKey := "417f4f4f430a0c67c2886afc20e11238:7acdc9a354518e5aa191803435f96dbad0177cb939e9290f34c5b4d1087bb2cc1c5b323ba7083c4875684fc375e183e3b42202830c6020cc257b91b6d8bcebb0c09bc1ffab982e77ea67449452c96af2dcf3cfc3909c66a0de62d463ab056f5c65e7dc22259e689357b4ace94098f3c38aa6bd3410c7bcbd12b94a5cca2aa578f83bd11cf79b3e8538c016fe335bf676a13f15a2f046d29c83a2e93e34fb49f21d6c625748f510301e1a930defc010ca369d6089639aae7328ba43134bf18d8d8517d1268d386098be71125e9fd0752f86cc03bbe67729e1008e87591c1eeee8d5b1209e4a1c77039ec0586a77be1ad8ea92173bedd66c0a15d84c6235b56fa1fb2f9c02dfb69e472f276dd957fca7fbe4ef519290ca48bf3c750707672998b11fe4c7de0ee9b7180c0a466a26f0b8ca9cf0625ea1154096501a8db1c65850c4e72d43c74e49d8a6f29d595fd5b685538664f3dcf7d938665cc2e12665c5934101e3a6ac85b8a72e12ef58343efc22b81492bba4f1964119879c9dad7904b37aef73c6c5cdce0000ffc45c70cb98364d130da7048c6092d5d62c08a14d86753223e42ea6969dbda17389c928b94ccbd0476092e1bdd3b0b3e4f8adbe33080421882bb3be876163f9987b3808d37c7bf3f4ba19081ff2b0bbc4557163dc81aae4790ee0df8d06a7a4b58a36b7200a61a73b9b75b75a3a901a875f534a52703ce2dd3247993931f234eec52d94fffc0d9da710baab4c35ff3de36f5300891a4cde9dd180cbfa90c6489c5480e89573489948cd780f64c2900e3c32586ee0088e14a0fcbbbb5d4ce99c3eae8f54bbfbe9d131390836df97985b018f0241274961893e6637671d6b348414391bb3b2f93bb42c59eafb36b8d3dde992acdd792c40544cf1cf9172503e869337d7afa7421a7b2b2ee38fd62adbbdd0c3ebf792a9891c"
	// // encryptedKeystoreName := "3753d1206611e6ab3ef686a3e2ca1c71:83a2f9316bafde517ccba982a8979baa75211fdbec7340536c057f411eb0cd34552098ccb46c452118a8f51195ebf474"
	// encryptedPassword := "94fa0973a57bfc09a676a4f2ab69b6d7:e8fa115d66beebf5e9df1cc781e357f6325c12737a8035444498e89e2dac9af4"

	stakerPublicKeyHex := "04c4fec2cca2602f5d2359d52627b60f08a2368491310dc26c4fbca104a678f7fb6d353dfc5aa0d68359168551a4a0625d3788297c8b4742997d8a616cec0db338"
	bStakerPubKey, err := hex.DecodeString(stakerPublicKeyHex)
	if err != nil {
		panic(err)
	}

	receivedStakerPubKeyPoint, err := crypto.UnmarshalPubkey(bStakerPubKey)
	if err != nil {
		panic(err)
	}

	_ = receivedStakerPubKeyPoint

	// fmt.Println(receivedStakerPubKeyPoint.X) [OK]

	nodeOperatorPrivKey := fromString(privateKey)

	fmt.Println("privKey:", nodeOperatorPrivKey.String())

	// Is this mod generic to use? because I didn't realy understand the math
	beMod, _ := big.NewInt(0).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	nodeOperatorPrivKey.Mod(nodeOperatorPrivKey, beMod)

	fmt.Println(nodeOperatorPrivKey)

	curve := crypto.S256()
	nodeOperatorSharedSecret, _ := curve.ScalarMult(receivedStakerPubKeyPoint.X, receivedStakerPubKeyPoint.Y, nodeOperatorPrivKey.Bytes())

	//pubx, puby := elliptic.Unmarshal(curve, bStakerPubKey)
	//fmt.Println(pubx, puby)
	//x, y := curve.ScalarMult(pubx, puby, nodeOperatorPrivKey.Bytes())

	fmt.Println("shared secret:", nodeOperatorSharedSecret)

	secretAsArray := nodeOperatorSharedSecret.Bytes()
	// secretAsArray := nodeOperatorSharedSecret.String()

	fmt.Println("secretAsArray:", len(secretAsArray), secretAsArray)

	// fmt.Println("e:", encryptedValidatorKey)
	// // fmt.Println("s:", hex.EncodeToString(secretAsArray))

	// validatorKeyString, _ := Decrypt(encryptedValidatorKey, secretAsArray)
	
	// fmt.Println(string(validatorKeyString))
	// validatorKeyPassword, _ := Decrypt(encryptedPassword, secretAsArray)
	// fmt.Println(string(validatorKeyPassword))
	// keystoreName, _ := Decrypt(encryptedKeystoreName, string(secretAsArray))

	// fmt.Println(validatorKeyString, validatorKeyPassword, keystoreName)

	// const validatorKeyString = decrypt(file["encryptedValidatorKey"], nodeOperatorSharedSecret.toArrayLike(Buffer, "be", 32));
  // const validatorKeyPassword = decrypt(file["encryptedPassword"],secretAsArray);
  // const keystoreName = decrypt(file["encryptedKeystoreName"],secretAsArray);

	_ = publicKey

	// fmt.Println(secretAsArray)

}