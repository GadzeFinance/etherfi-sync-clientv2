
const EC = require('elliptic')
const BN = require('bn.js')


const decryptValidatorKeyInfo = (keypairForIndex) => {
  const curve = new EC.ec("secp256k1");
  const { privateKey, publicKey } = keypairForIndex
  const stakerPublicKeyHex = "04c4fec2cca2602f5d2359d52627b60f08a2368491310dc26c4fbca104a678f7fb6d353dfc5aa0d68359168551a4a0625d3788297c8b4742997d8a616cec0db338"
  const receivedStakerPubKeyPoint = curve.keyFromPublic(stakerPublicKeyHex, "hex").getPublic();
  console.log("receivedStakerPubKeyPoint", receivedStakerPubKeyPoint.getX().toString())
  const nodeOperatorPrivKey = new BN(privateKey);

  const test1 = new BN("04653949f11");
  console.log("test1:", test1.toString())

  console.log("nodeOperatorPrivKey", nodeOperatorPrivKey.toString())
  const nodeOperatorSharedSecret = receivedStakerPubKeyPoint.mul(nodeOperatorPrivKey).getX();

  console.log("nodeOperatorSharedSecret", nodeOperatorSharedSecret.toString())
  const secretAsArray = nodeOperatorSharedSecret.toArrayLike(Buffer, "be", 32)
  console.log(secretAsArray)
  //const validatorKeyString = decrypt(file["encryptedValidatorKey"], nodeOperatorSharedSecret.toArrayLike(Buffer, "be", 32));
  //const validatorKeyPassword = decrypt(file["encryptedPassword"],secretAsArray);
  const keystoreName = decrypt(file["encryptedKeystoreName"],secretAsArray);
  console.log("keystoreName:", keystoreName)
  // return { validatorKeyFile: JSON.parse(validatorKeyString), validatorKeyPassword, keystoreName }
}

const keypairForIndex = {
  privateKey: "04653949f1185f5fc7cb441debe970c3049b7ee8e0a9d3d21d571f65d3c35fca8c3c678dfcf88afa23984596151c842f86c9856ae8c5cc6ff3f2812ce49888213a",
	publicKey: "74749212904585498199046374455021113593920534944642600239804505865032140419159"
}

decryptValidatorKeyInfo(keypairForIndex)