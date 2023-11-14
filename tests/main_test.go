package testing

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
)

func TestFetchFromIPFSCBC (t *testing.T) {

	config, _ := utils.GetAndCheckConfig()

	for _, hash := range GetIPFSHashes() {
		IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, hash)
		if err != nil {
			t.Error("Expected no errors in IPFS response, received error: ", err)
		}

		fieldNames := []string {
			"EncryptedKeystoreName",
			"EncryptedValidatorKey",
			"EncryptedPassword",
			"StakerPublicKey",
			"NodeOperatorPublicKey",
			"EtherfiDesktopAppVersion",
		}

		for _, fieldName := range fieldNames {
			value := reflect.ValueOf(IPFSResponse)
			field := value.FieldByName(fieldName)
			if !field.IsValid() {
				t.Error("Invalid field ", field, ". With value ", value)
			}
		}

	}
}

func TestDecryptCBCPrivatekey (t *testing.T) {
	privateKey, err := utils.ExtractPrivateKeysFromFS(filepath.Join("testFiles", "cbc_privateEtherfiKeystore.json"))
	if err != nil {
		t.Fail()
	}

	validatorKey, err := utils.DecryptPrivateKeysCBC(privateKey, GetCBCOperatorPassword())
	if err != nil {
		t.Fail()
	}

	if reflect.TypeOf(validatorKey.PrivateKeys).Kind() != reflect.Slice {
		t.Fail()
	}

	if reflect.TypeOf(validatorKey.PublicKeys).Kind() != reflect.Slice {
		fmt.Println(reflect.TypeOf(validatorKey.PublicKeys).Kind())
		t.Fail()
	}

	t.Run("Get Index CBC", func(t *testing.T) {

		expectedKeypair := GetCBCKeypair()
		Keypair, err := utils.GetKeyPairByPubKeyIndex("11", validatorKey.PrivateKeys, validatorKey.PublicKeys)
		if err != nil {
			t.Fail()
		}

		if expectedKeypair.PrivateKey != Keypair.PrivateKey {
			fmt.Println(Keypair.PrivateKey)
			t.Fail()
		}

		if expectedKeypair.PublicKey != Keypair.PublicKey {
			fmt.Println(Keypair.PublicKey)
			t.Fail()
		}
	})
	
}

func TestDecryptGCMPrivatekey (t *testing.T) {
	privateKey, err := utils.ExtractPrivateKeysFromFS(filepath.Join("testFiles", "gcm_privateEtherfiKeystore.json"))
	if err != nil {
		t.Error("Expected no errors in extracting keys, received error: ", err)
	}

	validatorKey, err := utils.DecryptPrivateKeysGCM(privateKey, GetGCMOperatorPassword())
	if err != nil {
		t.Error("Expected no errors in Decrypting keys, received error: ", err)
	}

	if reflect.TypeOf(validatorKey.PrivateKeys).Kind() != reflect.Slice {
		t.Error("Private key is not array ", reflect.TypeOf(validatorKey.PublicKeys).Kind())
	}

	if reflect.TypeOf(validatorKey.PublicKeys).Kind() != reflect.Slice {
		t.Error("Private key is not array, is: ", reflect.TypeOf(validatorKey.PublicKeys).Kind())
	}

	t.Run("Get Index GCM", func(t *testing.T) {

		gcmKeypair := GetGCMKeypair()
		Keypair, err := utils.GetKeyPairByPubKeyIndex("0", validatorKey.PrivateKeys, validatorKey.PublicKeys)
		if err != nil {
			t.Fail()
		}

		if gcmKeypair.PrivateKey != Keypair.PrivateKey {
			fmt.Println(Keypair.PrivateKey)
			t.Fail()
		}

		if gcmKeypair.PublicKey != Keypair.PublicKey {
			fmt.Println(Keypair.PublicKey)
			t.Fail()
		}
	})
}

func TestDecryptValidatorKeysCBC(t *testing.T) {

	config, _ := utils.GetAndCheckConfig()

	IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, "QmX4eYKNXVmBpa4ZqBxKAZNgPa3KpTB1Ub5KBJTRtNqAaf")
	if err != nil {
		t.Error("Expected no errors in IPFS response, received error: ", err)
	}

	keyPair := GetCBCKeypair()

	data := utils.DecryptValidatorKeyInfo(IPFSResponse, keyPair)

	if string(data.ValidatorKeyPassword) != GetCBCValidatorPassword() {
		t.Fail()
	}

	if string(data.KeystoreName) != "keystore-m_12381_3600_0_0_0-1684942379.json" {
		t.Fail()
	}	
}

func TestDecryptValidatorKeysGCM(t *testing.T) {

	config, _ := utils.GetAndCheckConfig()

	IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, "QmUwkoSJBoq8tNG8sNkuxkf6w8ADMrARLqzWqk7qYCDBJw")
	if err != nil {
		t.Error("Expected no errors in IPFS response, received error: ", err)
	}

	keyPair := GetGCMKeypair()

	data := utils.DecryptValidatorKeyInfo(IPFSResponse, keyPair)

	if string(data.ValidatorKeyPassword) != GetGCMValidatorPassword() {
		t.Fail()
	}

	if string(data.KeystoreName) != "keystore-m_12381_3600_0_0_0-1684873868.json" {
		t.Fail()
	}	
}