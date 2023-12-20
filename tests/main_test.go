package testing

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
)

func TestFetchFromIPFSCBC(t *testing.T) {

	config, err := utils.GetAndCheckConfig("./testdata/config.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, hash := range GetIPFSHashes() {
		IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, hash)
		if err != nil {
			t.Fatalf("Expected no errors in IPFS response, received error: %v", err)
		}

		fieldNames := []string{
			"EncryptedKeystoreName",
			"EncryptedValidatorKey",
			"EncryptedPassword",
			"StakerPublicKey",
			"NodeOperatorPublicKey",
			"EtherfiDesktopAppVersion",
		}

		for _, fieldName := range fieldNames {
			value := reflect.ValueOf(*IPFSResponse)
			field := value.FieldByName(fieldName)
			if !field.IsValid() {
				t.Error("Invalid field ", field, ". With value ", value)
			}
		}

	}
}

func TestDecryptCBCPrivatekey(t *testing.T) {
	privateKey, err := utils.ParseKeystoreFile(filepath.Join("testdata", "cbc_privateEtherfiKeystore.json"))
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

func TestDecryptGCMPrivatekey(t *testing.T) {
	privateKey, err := utils.ParseKeystoreFile(filepath.Join("testdata", "gcm_privateEtherfiKeystore.json"))
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

	config, err := utils.GetAndCheckConfig("./testdata/config.json")
	if err != nil {
		t.Fatal(err)
	}

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

	config, err := utils.GetAndCheckConfig("./testdata/config.json")
	if err != nil {
		t.Fatal(err)
	}

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

func GetIPFSHashes() []string {
	ipfsHashes := []string{
		"QmXA3uT5wnXfwMYbEajFUNHDPv9qENrfW3quPL9KkCNrE4",
		"QmX4eYKNXVmBpa4ZqBxKAZNgPa3KpTB1Ub5KBJTRtNqAaf",
		"QmdEytzWJ2atAToYD57jShkPqns5FaktrPnZib1A4FwHgk",
		"QmUJ5xsHBPz4HBHBVNYRu8h6Xnpawtk6rxzvr89m6Q57sF",
	}

	return ipfsHashes
}

func GetCBCOperatorPassword() string {
	return "$M00THOp3rat0R"
}

func GetCBCKeypair() schemas.KeyPair {
	return schemas.KeyPair{
		PrivateKey: "84882960453863968714531524381150657937041302799814311443266907307945660872829",
		PublicKey:  "04a600b60d602e2ffd8b77ffd18812d2ce938d4421281fb0bb47c28f54a1562f66fa364099c8ce9d2270a044341c00ac6f3047faeda9b251a109f4d0dfff388c98",
	}
}

func GetGCMOperatorPassword() string {
	return "Password123!"
}

func GetGCMKeypair() schemas.KeyPair {
	return schemas.KeyPair{
		PrivateKey: "38331824479263245210020673306200128332224251395654410960657076080784808684342",
		PublicKey:  "04f262c21a97f93bf361645e9bb23b6b36a9bdff68e579f20b6653025ac5edc465005cac64f91ef82a735be8bfe6577e3b400a362b498576c62217c4a513cc8d79",
	}
}

func GetCBCValidatorPassword() string {
	return "CrazyNewPassword157!"
}

func GetGCMValidatorPassword() string {
	return "lU8BKjqlN6K8yDPYZIiF"
}
