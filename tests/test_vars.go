package testing

import "github.com/GadzeFinance/etherfi-sync-clientv2/schemas"


func GetIPFSHashes() []string {
	ipfsHashes := []string {
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
		PublicKey: "04a600b60d602e2ffd8b77ffd18812d2ce938d4421281fb0bb47c28f54a1562f66fa364099c8ce9d2270a044341c00ac6f3047faeda9b251a109f4d0dfff388c98",
	}
}

func GetGCMOperatorPassword() string {
	return "Password123!"
}

func GetGCMKeypair() schemas.KeyPair {
	return schemas.KeyPair{
		PrivateKey: "38331824479263245210020673306200128332224251395654410960657076080784808684342",
		PublicKey: "04f262c21a97f93bf361645e9bb23b6b36a9bdff68e579f20b6653025ac5edc465005cac64f91ef82a735be8bfe6577e3b400a362b498576c62217c4a513cc8d79",
	}
}

func GetCBCValidatorPassword() string {
	return "CrazyNewPassword157!"
}

func GetGCMValidatorPassword() string {
	return "lU8BKjqlN6K8yDPYZIiF"
}