package schemas

type Config struct {
	GRAPH_URL string `json:"GRAPH_URL"`
	BIDDER string `json:"BIDDER"`
	PRIVATE_KEYS_FILE_LOCATION string `json:"PRIVATE_KEYS_FILE_LOCATION"`
	OUTPUT_LOCATION string `json:"OUTPUT_LOCATION"`
	PASSWORD string `json:"PASSWORD"`
	IPFS_GATEWAY string `json:"IPFS_GATEWAY"`
	CONSENSUS_FOLDER_LOCATION string `json:"CONSENSUS_FOLDER_LOCATION"`
	ETHERFI_SC_CLIENT_LOCATION string `json:"ETHERFI_SC_CLIENT_LOCATION"`
	PATH_TO_PRYSYM_SH string `json:"PATH_TO_PRYSYM_SH"`
}

type ValidatorKeyInfo struct {
	ValidatorKeyFile []byte `json:"validatorKeyFile"`
	ValidatorKeyPassword []byte `json:"validatorKeyPassword"`
	KeystoreName []byte `json:"keystoreName"`
}

type KeyStoreFile struct {
	Iv string `json:"iv"`
	Salt string `json:"salt"`
	Data string `json:"data"`
	AuthTag string `json:"authTag`
	EtherfiDesktopAppVersion string `json:"etherfiDesktopAppVersion"`
}

type IPFSResponseType struct {
	EncryptedKeystoreName string `json:"encryptedKeystoreName"`
	EncryptedValidatorKey string `json:"encryptedValidatorKey"`
	EncryptedPassword string `json:"encryptedPassword"`
	StakerPublicKey string `json:"stakerPublicKey"`
	NodeOperatorPublicKey string `json:"nodeOperatorPublicKey"`
	EtherfiDesktopAppVersion string `json:"etherfiDesktopAppVersion"`
}

type GQLResponseType struct {
	Data struct {
		Bids []BidType `json:"bids"`
	} `json:"data"`
}

type BidType struct {
	Id string `json:"id"`
	BidderAddress string `json:"bidderAddress"`
  PubKeyIndex string `json:"pubKeyIndex"`
  Validator ValidatorType `json:"validator"`
}

type ValidatorType struct {
	Id string `json:"id"`
	Phase string `json:"phase"`
  IpfsHashForEncryptedValidatorKey string `json:"ipfsHashForEncryptedValidatorKey"`
  ValidatorPubKey string `json:"validatorPubKey"`        	
}

type DecryptedDataJSON struct {
	PublicKeys []string `json:"pubKeyArray"`
	PrivateKeys []string `json:"privKeyArray"`
}

type KeyPair struct {
	PublicKey string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}
