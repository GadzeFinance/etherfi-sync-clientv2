package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	config, err := utils.GetAndCheckConfig("config.json")
	if err != nil {
		return fmt.Errorf("Failed to load config: %w\n", err)
	}

	fmt.Println("Starting EtherFi Sync Client:")
	fmt.Println("Operator Address: ", config.BIDDER)
	fmt.Println("Output directory: ", config.OUTPUT_LOCATION)

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		return fmt.Errorf("failed to open sqlite db: %w", err)
	}
	defer db.Close()

	if err := utils.CreateTable(db); err != nil {
		return fmt.Errorf("crating bids table: %w", err)
	}

	return fetchValidatorKeys(config, db)
}

func fetchValidatorKeys(cfg schemas.Config, db *sql.DB) error {
	var (
		stakingManagerAddr      common.Address
		auctionManagerAddr      common.Address
		etherFiNodesManagerAddr common.Address
	)
	if cfg.NETWORK == "" && (cfg.NETWORK != "hoodi" && cfg.NETWORK != "mainnet") {
		return fmt.Errorf("NETWORK is required in config and must be either 'hoodi' or 'mainnet'")
	}
	if cfg.NETWORK == "mainnet" {
		stakingManagerAddr = common.HexToAddress("0x25e821b7197B146F7713C3b89B6A4D83516B912d")
		auctionManagerAddr = common.HexToAddress("0x00C452aFFee3a17d9Cecc1Bcd2B8d5C7635C4CB9")
		etherFiNodesManagerAddr = common.HexToAddress("0x8B71140AD2e5d1E7018d2a7f8a288BD3CD38916F")
	} else if cfg.NETWORK == "hoodi" {
		stakingManagerAddr = common.HexToAddress("0xDbE50E32Ed95f539F36bA315a75377FBc35aBc12")
		auctionManagerAddr = common.HexToAddress("0x261315c176864cE29D582f38DdA4930ED17CD95A")
		etherFiNodesManagerAddr = common.HexToAddress("0x7579194b8265e3Aa7df451c6BD2aff5B1FC5F945")
	}
	fmt.Println("Network: ", cfg.NETWORK)
	fmt.Println("Staking Manager Address: ", stakingManagerAddr)
	fmt.Println("Auction Manager Address: ", auctionManagerAddr)
	fmt.Println("EtherFi Nodes Manager Address: ", etherFiNodesManagerAddr)
	operator := common.HexToAddress(cfg.BIDDER)
	rpcClient, err := ethclient.Dial(cfg.RPC_URL)
	if err != nil {
		return fmt.Errorf("failed to dial RPC: %w", err)
	}
	stakingManager, _ := NewStakingManager(stakingManagerAddr, rpcClient)
	auctionManager, _ := NewAuctionManager(auctionManagerAddr, rpcClient)
	etherFiNodesManager, _ := NewEtherFiNodesManager(etherFiNodesManagerAddr, rpcClient)

	operatorKeystore, err := utils.ParseKeystoreFile(cfg.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("parsing keystore: %w", err)
	}
	// For compatibility, if the authTag is empty, we know it's CBC mode
	isUsingCBC := false
	if operatorKeystore.AuthTag == "" {
		isUsingCBC = true
	}

	registrationEvents, err := stakingManager.FilterValidatorRegistered(nil, []common.Address{operator}, nil, nil)
	if err != nil {
		panic(err)
	}
	for registrationEvents.Next() {
		event := registrationEvents.Event

		// the original authors did not add a primary key or any sort of key to database :(
		// so we check if there are any number of entries tied to this validator id
		count, err := utils.GetIDCount(db, event.ValidatorId.String())
		if err != nil {
			return fmt.Errorf("failed to check existence of validator in db")
		}
		if count > 0 {
			// we have already downloaded keys for this validator
			continue
		}

		// fetch the bid data from auctionManager
		bid, err := auctionManager.Bids(&bind.CallOpts{}, event.ValidatorId)
		if err != nil {
			return fmt.Errorf("failed fetch bid information for bid %s: %w", event.ValidatorId, err)
		}

		// grab the associated node
		nodeAddr, err := etherFiNodesManager.EtherfiNodeAddress(nil, event.ValidatorId)
		if err != nil {
			return fmt.Errorf("failed to fetch node associated with id %s: %w", event.ValidatorId, err)
		}

		fmt.Println(`Processing stake request for validator: ` + event.ValidatorId.String() + ` and ipfs path: ` + event.IpfsHashForEncryptedValidatorKey)

		IPFSResponse, err := utils.FetchFromIPFS(cfg.IPFS_GATEWAY, event.IpfsHashForEncryptedValidatorKey)
		if err != nil {
			fmt.Println("failed to fetch encrypted key from IPFS, Retrying Once: " + event.IpfsHashForEncryptedValidatorKey)
			IPFSResponse, err = utils.FetchFromIPFS(cfg.IPFS_GATEWAY, event.IpfsHashForEncryptedValidatorKey)
			if err != nil {
				return fmt.Errorf("FetchFromIPFS: %w", err)
			}
		}

		var operatorEncryptionKeys schemas.DecryptedDataJSON
		if isUsingCBC {
			operatorEncryptionKeys, err = utils.DecryptPrivateKeysCBC(operatorKeystore, cfg.PASSWORD)
		} else {
			operatorEncryptionKeys, err = utils.DecryptPrivateKeysGCM(operatorKeystore, cfg.PASSWORD)
		}
		if err != nil {
			return fmt.Errorf("DecryptPrivateKeys: %w", err)
		}

		pubKeyArray := operatorEncryptionKeys.PublicKeys
		privKeyArray := operatorEncryptionKeys.PrivateKeys
		keypairForIndex, err := utils.GetKeyPairByPubKeyIndex(int64(bid.BidderPubKeyIndex), privKeyArray, pubKeyArray)
		if err != nil {
			return fmt.Errorf("GetKeyPairByPubKeyIndex: %w", err)
		}

		data := utils.DecryptValidatorKeyInfo(IPFSResponse, keypairForIndex)
		if err := utils.SaveKeysToFS(cfg.OUTPUT_LOCATION, data, event.ValidatorId.String(), int64(bid.BidderPubKeyIndex), hexutil.Encode(event.ValidatorPubKey), nodeAddr.Hex(), db); err != nil {
			return fmt.Errorf("SaveKeysToFS: %w", err)
		}

	}

	return nil
}
