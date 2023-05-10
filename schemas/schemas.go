package schemas

type Config struct {
	GRAPH_URL string `json:"GRAPH_URL"`
	BIDDER string `json:"BIDDER"`
	PRIVATE_KEYS_FILE_LOCATION string `json:"PRIVATE_KEYS_FILE_LOCATION"`
	OUTPUT_LOCATION string `json:"OUTPUT_LOCATION"`
	PASSWORD string `json:"PASSWORD"`
	IPFS_GATEWAY string `json:"IPFS_GATEWAY"`
}

type KeyStoreFile struct {
	Iv string `json:"iv"`
	Salt string `json:"salt"`
	Data string `json:"data"`
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
	PublicKeys []string `json:"public_keys"`
	PrivateKeys []string `json:"private_keys"`
}