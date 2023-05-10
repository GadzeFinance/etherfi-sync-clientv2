package main

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