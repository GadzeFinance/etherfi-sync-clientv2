// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IStakingManagerDepositData is an auto generated low-level Go binding around an user-defined struct.
type IStakingManagerDepositData struct {
	PublicKey                        []byte
	Signature                        []byte
	DepositDataRoot                  [32]byte
	IpfsHashForEncryptedValidatorKey string
}

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_liquidityPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_etherFiNodesManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ethDepositContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_auctionManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_etherFiNodeBeacon\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_roleRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_NODE_CREATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAuctionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateValidatorPubkeyHash\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"confirmAndFundBeaconValidators\",\"inputs\":[{\"name\":\"depositData\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingManager.DepositData[]\",\"components\":[{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ipfsHashForEncryptedValidatorKey\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"validatorSizeWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createBeaconValidators\",\"inputs\":[{\"name\":\"depositData\",\"type\":\"tuple[]\",\"internalType\":\"structIStakingManager.DepositData[]\",\"components\":[{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ipfsHashForEncryptedValidatorKey\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"bidIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositContractEth2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDepositContract\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"etherFiNodeBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractUpgradeableBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"etherFiNodesManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEtherFiNodesManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEtherFiNodeBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"instantiateEtherFiNode\",\"inputs\":[{\"name\":\"_createEigenPod\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidityPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roleRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRoleRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unPauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeEtherFiNode\",\"inputs\":[{\"name\":\"_newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bNftOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tNftOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"validatorPubKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"ipfsHashForEncryptedValidatorKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"linkLegacyValidatorId\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"legacyId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorConfirmed\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"bnftRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tnftRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorCreated\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InactiveBid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectBeaconRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectRole\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDepositData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEtherFiNode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgrade\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnlinkedPubkey\",\"inputs\":[]}]",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// STAKINGMANAGERNODECREATORROLE is a free data retrieval call binding the contract method 0x7cb2abcb.
//
// Solidity: function STAKING_MANAGER_NODE_CREATOR_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGMANAGERNODECREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_MANAGER_NODE_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERNODECREATORROLE is a free data retrieval call binding the contract method 0x7cb2abcb.
//
// Solidity: function STAKING_MANAGER_NODE_CREATOR_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGMANAGERNODECREATORROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGMANAGERNODECREATORROLE(&_StakingManager.CallOpts)
}

// STAKINGMANAGERNODECREATORROLE is a free data retrieval call binding the contract method 0x7cb2abcb.
//
// Solidity: function STAKING_MANAGER_NODE_CREATOR_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGMANAGERNODECREATORROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGMANAGERNODECREATORROLE(&_StakingManager.CallOpts)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_StakingManager *StakingManagerCaller) AuctionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "auctionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_StakingManager *StakingManagerSession) AuctionManager() (common.Address, error) {
	return _StakingManager.Contract.AuctionManager(&_StakingManager.CallOpts)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_StakingManager *StakingManagerCallerSession) AuctionManager() (common.Address, error) {
	return _StakingManager.Contract.AuctionManager(&_StakingManager.CallOpts)
}

// CalculateValidatorPubkeyHash is a free data retrieval call binding the contract method 0x49d58951.
//
// Solidity: function calculateValidatorPubkeyHash(bytes pubkey) pure returns(bytes32)
func (_StakingManager *StakingManagerCaller) CalculateValidatorPubkeyHash(opts *bind.CallOpts, pubkey []byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "calculateValidatorPubkeyHash", pubkey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateValidatorPubkeyHash is a free data retrieval call binding the contract method 0x49d58951.
//
// Solidity: function calculateValidatorPubkeyHash(bytes pubkey) pure returns(bytes32)
func (_StakingManager *StakingManagerSession) CalculateValidatorPubkeyHash(pubkey []byte) ([32]byte, error) {
	return _StakingManager.Contract.CalculateValidatorPubkeyHash(&_StakingManager.CallOpts, pubkey)
}

// CalculateValidatorPubkeyHash is a free data retrieval call binding the contract method 0x49d58951.
//
// Solidity: function calculateValidatorPubkeyHash(bytes pubkey) pure returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) CalculateValidatorPubkeyHash(pubkey []byte) ([32]byte, error) {
	return _StakingManager.Contract.CalculateValidatorPubkeyHash(&_StakingManager.CallOpts, pubkey)
}

// DepositContractEth2 is a free data retrieval call binding the contract method 0x82699d5a.
//
// Solidity: function depositContractEth2() view returns(address)
func (_StakingManager *StakingManagerCaller) DepositContractEth2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "depositContractEth2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContractEth2 is a free data retrieval call binding the contract method 0x82699d5a.
//
// Solidity: function depositContractEth2() view returns(address)
func (_StakingManager *StakingManagerSession) DepositContractEth2() (common.Address, error) {
	return _StakingManager.Contract.DepositContractEth2(&_StakingManager.CallOpts)
}

// DepositContractEth2 is a free data retrieval call binding the contract method 0x82699d5a.
//
// Solidity: function depositContractEth2() view returns(address)
func (_StakingManager *StakingManagerCallerSession) DepositContractEth2() (common.Address, error) {
	return _StakingManager.Contract.DepositContractEth2(&_StakingManager.CallOpts)
}

// EtherFiNodeBeacon is a free data retrieval call binding the contract method 0xb149e447.
//
// Solidity: function etherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerCaller) EtherFiNodeBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "etherFiNodeBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherFiNodeBeacon is a free data retrieval call binding the contract method 0xb149e447.
//
// Solidity: function etherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerSession) EtherFiNodeBeacon() (common.Address, error) {
	return _StakingManager.Contract.EtherFiNodeBeacon(&_StakingManager.CallOpts)
}

// EtherFiNodeBeacon is a free data retrieval call binding the contract method 0xb149e447.
//
// Solidity: function etherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerCallerSession) EtherFiNodeBeacon() (common.Address, error) {
	return _StakingManager.Contract.EtherFiNodeBeacon(&_StakingManager.CallOpts)
}

// EtherFiNodesManager is a free data retrieval call binding the contract method 0x089acd98.
//
// Solidity: function etherFiNodesManager() view returns(address)
func (_StakingManager *StakingManagerCaller) EtherFiNodesManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "etherFiNodesManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherFiNodesManager is a free data retrieval call binding the contract method 0x089acd98.
//
// Solidity: function etherFiNodesManager() view returns(address)
func (_StakingManager *StakingManagerSession) EtherFiNodesManager() (common.Address, error) {
	return _StakingManager.Contract.EtherFiNodesManager(&_StakingManager.CallOpts)
}

// EtherFiNodesManager is a free data retrieval call binding the contract method 0x089acd98.
//
// Solidity: function etherFiNodesManager() view returns(address)
func (_StakingManager *StakingManagerCallerSession) EtherFiNodesManager() (common.Address, error) {
	return _StakingManager.Contract.EtherFiNodesManager(&_StakingManager.CallOpts)
}

// GetEtherFiNodeBeacon is a free data retrieval call binding the contract method 0xe9c99b6b.
//
// Solidity: function getEtherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerCaller) GetEtherFiNodeBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getEtherFiNodeBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEtherFiNodeBeacon is a free data retrieval call binding the contract method 0xe9c99b6b.
//
// Solidity: function getEtherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerSession) GetEtherFiNodeBeacon() (common.Address, error) {
	return _StakingManager.Contract.GetEtherFiNodeBeacon(&_StakingManager.CallOpts)
}

// GetEtherFiNodeBeacon is a free data retrieval call binding the contract method 0xe9c99b6b.
//
// Solidity: function getEtherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetEtherFiNodeBeacon() (common.Address, error) {
	return _StakingManager.Contract.GetEtherFiNodeBeacon(&_StakingManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StakingManager *StakingManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StakingManager *StakingManagerSession) Implementation() (common.Address, error) {
	return _StakingManager.Contract.Implementation(&_StakingManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Implementation() (common.Address, error) {
	return _StakingManager.Contract.Implementation(&_StakingManager.CallOpts)
}

// InitialDepositAmount is a free data retrieval call binding the contract method 0x21edd097.
//
// Solidity: function initialDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) InitialDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "initialDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialDepositAmount is a free data retrieval call binding the contract method 0x21edd097.
//
// Solidity: function initialDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) InitialDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.InitialDepositAmount(&_StakingManager.CallOpts)
}

// InitialDepositAmount is a free data retrieval call binding the contract method 0x21edd097.
//
// Solidity: function initialDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) InitialDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.InitialDepositAmount(&_StakingManager.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_StakingManager *StakingManagerCaller) LiquidityPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "liquidityPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_StakingManager *StakingManagerSession) LiquidityPool() (common.Address, error) {
	return _StakingManager.Contract.LiquidityPool(&_StakingManager.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_StakingManager *StakingManagerCallerSession) LiquidityPool() (common.Address, error) {
	return _StakingManager.Contract.LiquidityPool(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerSession) Paused() (bool, error) {
	return _StakingManager.Contract.Paused(&_StakingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) Paused() (bool, error) {
	return _StakingManager.Contract.Paused(&_StakingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingManager.Contract.ProxiableUUID(&_StakingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingManager.Contract.ProxiableUUID(&_StakingManager.CallOpts)
}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_StakingManager *StakingManagerCaller) RoleRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "roleRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_StakingManager *StakingManagerSession) RoleRegistry() (common.Address, error) {
	return _StakingManager.Contract.RoleRegistry(&_StakingManager.CallOpts)
}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_StakingManager *StakingManagerCallerSession) RoleRegistry() (common.Address, error) {
	return _StakingManager.Contract.RoleRegistry(&_StakingManager.CallOpts)
}

// ConfirmAndFundBeaconValidators is a paid mutator transaction binding the contract method 0xbaaff116.
//
// Solidity: function confirmAndFundBeaconValidators((bytes,bytes,bytes32,string)[] depositData, uint256 validatorSizeWei) payable returns()
func (_StakingManager *StakingManagerTransactor) ConfirmAndFundBeaconValidators(opts *bind.TransactOpts, depositData []IStakingManagerDepositData, validatorSizeWei *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "confirmAndFundBeaconValidators", depositData, validatorSizeWei)
}

// ConfirmAndFundBeaconValidators is a paid mutator transaction binding the contract method 0xbaaff116.
//
// Solidity: function confirmAndFundBeaconValidators((bytes,bytes,bytes32,string)[] depositData, uint256 validatorSizeWei) payable returns()
func (_StakingManager *StakingManagerSession) ConfirmAndFundBeaconValidators(depositData []IStakingManagerDepositData, validatorSizeWei *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ConfirmAndFundBeaconValidators(&_StakingManager.TransactOpts, depositData, validatorSizeWei)
}

// ConfirmAndFundBeaconValidators is a paid mutator transaction binding the contract method 0xbaaff116.
//
// Solidity: function confirmAndFundBeaconValidators((bytes,bytes,bytes32,string)[] depositData, uint256 validatorSizeWei) payable returns()
func (_StakingManager *StakingManagerTransactorSession) ConfirmAndFundBeaconValidators(depositData []IStakingManagerDepositData, validatorSizeWei *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ConfirmAndFundBeaconValidators(&_StakingManager.TransactOpts, depositData, validatorSizeWei)
}

// CreateBeaconValidators is a paid mutator transaction binding the contract method 0xb71205d4.
//
// Solidity: function createBeaconValidators((bytes,bytes,bytes32,string)[] depositData, uint256[] bidIds, address etherFiNode) payable returns()
func (_StakingManager *StakingManagerTransactor) CreateBeaconValidators(opts *bind.TransactOpts, depositData []IStakingManagerDepositData, bidIds []*big.Int, etherFiNode common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "createBeaconValidators", depositData, bidIds, etherFiNode)
}

// CreateBeaconValidators is a paid mutator transaction binding the contract method 0xb71205d4.
//
// Solidity: function createBeaconValidators((bytes,bytes,bytes32,string)[] depositData, uint256[] bidIds, address etherFiNode) payable returns()
func (_StakingManager *StakingManagerSession) CreateBeaconValidators(depositData []IStakingManagerDepositData, bidIds []*big.Int, etherFiNode common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.CreateBeaconValidators(&_StakingManager.TransactOpts, depositData, bidIds, etherFiNode)
}

// CreateBeaconValidators is a paid mutator transaction binding the contract method 0xb71205d4.
//
// Solidity: function createBeaconValidators((bytes,bytes,bytes32,string)[] depositData, uint256[] bidIds, address etherFiNode) payable returns()
func (_StakingManager *StakingManagerTransactorSession) CreateBeaconValidators(depositData []IStakingManagerDepositData, bidIds []*big.Int, etherFiNode common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.CreateBeaconValidators(&_StakingManager.TransactOpts, depositData, bidIds, etherFiNode)
}

// InstantiateEtherFiNode is a paid mutator transaction binding the contract method 0xaeeb9556.
//
// Solidity: function instantiateEtherFiNode(bool _createEigenPod) returns(address)
func (_StakingManager *StakingManagerTransactor) InstantiateEtherFiNode(opts *bind.TransactOpts, _createEigenPod bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "instantiateEtherFiNode", _createEigenPod)
}

// InstantiateEtherFiNode is a paid mutator transaction binding the contract method 0xaeeb9556.
//
// Solidity: function instantiateEtherFiNode(bool _createEigenPod) returns(address)
func (_StakingManager *StakingManagerSession) InstantiateEtherFiNode(_createEigenPod bool) (*types.Transaction, error) {
	return _StakingManager.Contract.InstantiateEtherFiNode(&_StakingManager.TransactOpts, _createEigenPod)
}

// InstantiateEtherFiNode is a paid mutator transaction binding the contract method 0xaeeb9556.
//
// Solidity: function instantiateEtherFiNode(bool _createEigenPod) returns(address)
func (_StakingManager *StakingManagerTransactorSession) InstantiateEtherFiNode(_createEigenPod bool) (*types.Transaction, error) {
	return _StakingManager.Contract.InstantiateEtherFiNode(&_StakingManager.TransactOpts, _createEigenPod)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_StakingManager *StakingManagerTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_StakingManager *StakingManagerSession) PauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.PauseContract(&_StakingManager.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_StakingManager *StakingManagerTransactorSession) PauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.PauseContract(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_StakingManager *StakingManagerTransactor) UnPauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unPauseContract")
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_StakingManager *StakingManagerSession) UnPauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.UnPauseContract(&_StakingManager.TransactOpts)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_StakingManager *StakingManagerTransactorSession) UnPauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.UnPauseContract(&_StakingManager.TransactOpts)
}

// UpgradeEtherFiNode is a paid mutator transaction binding the contract method 0x49370974.
//
// Solidity: function upgradeEtherFiNode(address _newImplementation) returns()
func (_StakingManager *StakingManagerTransactor) UpgradeEtherFiNode(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeEtherFiNode", _newImplementation)
}

// UpgradeEtherFiNode is a paid mutator transaction binding the contract method 0x49370974.
//
// Solidity: function upgradeEtherFiNode(address _newImplementation) returns()
func (_StakingManager *StakingManagerSession) UpgradeEtherFiNode(_newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeEtherFiNode(&_StakingManager.TransactOpts, _newImplementation)
}

// UpgradeEtherFiNode is a paid mutator transaction binding the contract method 0x49370974.
//
// Solidity: function upgradeEtherFiNode(address _newImplementation) returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeEtherFiNode(_newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeEtherFiNode(&_StakingManager.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingManager *StakingManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingManager *StakingManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeTo(&_StakingManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeTo(&_StakingManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeToAndCall(&_StakingManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeToAndCall(&_StakingManager.TransactOpts, newImplementation, data)
}

// StakingManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StakingManager contract.
type StakingManagerAdminChangedIterator struct {
	Event *StakingManagerAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAdminChanged represents a AdminChanged event raised by the StakingManager contract.
type StakingManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingManager *StakingManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StakingManagerAdminChangedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAdminChangedIterator{contract: _StakingManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingManager *StakingManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAdminChanged)
				if err := _StakingManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingManager *StakingManagerFilterer) ParseAdminChanged(log types.Log) (*StakingManagerAdminChanged, error) {
	event := new(StakingManagerAdminChanged)
	if err := _StakingManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the StakingManager contract.
type StakingManagerBeaconUpgradedIterator struct {
	Event *StakingManagerBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerBeaconUpgraded represents a BeaconUpgraded event raised by the StakingManager contract.
type StakingManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingManager *StakingManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StakingManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerBeaconUpgradedIterator{contract: _StakingManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingManager *StakingManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StakingManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerBeaconUpgraded)
				if err := _StakingManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingManager *StakingManagerFilterer) ParseBeaconUpgraded(log types.Log) (*StakingManagerBeaconUpgraded, error) {
	event := new(StakingManagerBeaconUpgraded)
	if err := _StakingManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingManager contract.
type StakingManagerInitializedIterator struct {
	Event *StakingManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialized represents a Initialized event raised by the StakingManager contract.
type StakingManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingManager *StakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingManagerInitializedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitializedIterator{contract: _StakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingManager *StakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialized)
				if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingManager *StakingManagerFilterer) ParseInitialized(log types.Log) (*StakingManagerInitialized, error) {
	event := new(StakingManagerInitialized)
	if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingManager contract.
type StakingManagerPausedIterator struct {
	Event *StakingManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerPaused represents a Paused event raised by the StakingManager contract.
type StakingManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingManagerPausedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingManagerPausedIterator{contract: _StakingManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingManagerPaused) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerPaused)
				if err := _StakingManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) ParsePaused(log types.Log) (*StakingManagerPaused, error) {
	event := new(StakingManagerPaused)
	if err := _StakingManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingManager contract.
type StakingManagerUnpausedIterator struct {
	Event *StakingManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnpaused represents a Unpaused event raised by the StakingManager contract.
type StakingManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingManagerUnpausedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnpausedIterator{contract: _StakingManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnpaused)
				if err := _StakingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) ParseUnpaused(log types.Log) (*StakingManagerUnpaused, error) {
	event := new(StakingManagerUnpaused)
	if err := _StakingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakingManager contract.
type StakingManagerUpgradedIterator struct {
	Event *StakingManagerUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUpgraded represents a Upgraded event raised by the StakingManager contract.
type StakingManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUpgradedIterator{contract: _StakingManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUpgraded)
				if err := _StakingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) ParseUpgraded(log types.Log) (*StakingManagerUpgraded, error) {
	event := new(StakingManagerUpgraded)
	if err := _StakingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the StakingManager contract.
type StakingManagerValidatorRegisteredIterator struct {
	Event *StakingManagerValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerValidatorRegistered represents a ValidatorRegistered event raised by the StakingManager contract.
type StakingManagerValidatorRegistered struct {
	Operator                         common.Address
	BNftOwner                        common.Address
	TNftOwner                        common.Address
	ValidatorId                      *big.Int
	ValidatorPubKey                  []byte
	IpfsHashForEncryptedValidatorKey string
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x0b43d988cd5ab75ae318de41d6871d4b26efe57c3f3975331873f4dc073041fc.
//
// Solidity: event ValidatorRegistered(address indexed operator, address indexed bNftOwner, address indexed tNftOwner, uint256 validatorId, bytes validatorPubKey, string ipfsHashForEncryptedValidatorKey)
func (_StakingManager *StakingManagerFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, operator []common.Address, bNftOwner []common.Address, tNftOwner []common.Address) (*StakingManagerValidatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var bNftOwnerRule []interface{}
	for _, bNftOwnerItem := range bNftOwner {
		bNftOwnerRule = append(bNftOwnerRule, bNftOwnerItem)
	}
	var tNftOwnerRule []interface{}
	for _, tNftOwnerItem := range tNftOwner {
		tNftOwnerRule = append(tNftOwnerRule, tNftOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ValidatorRegistered", operatorRule, bNftOwnerRule, tNftOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerValidatorRegisteredIterator{contract: _StakingManager.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x0b43d988cd5ab75ae318de41d6871d4b26efe57c3f3975331873f4dc073041fc.
//
// Solidity: event ValidatorRegistered(address indexed operator, address indexed bNftOwner, address indexed tNftOwner, uint256 validatorId, bytes validatorPubKey, string ipfsHashForEncryptedValidatorKey)
func (_StakingManager *StakingManagerFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *StakingManagerValidatorRegistered, operator []common.Address, bNftOwner []common.Address, tNftOwner []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var bNftOwnerRule []interface{}
	for _, bNftOwnerItem := range bNftOwner {
		bNftOwnerRule = append(bNftOwnerRule, bNftOwnerItem)
	}
	var tNftOwnerRule []interface{}
	for _, tNftOwnerItem := range tNftOwner {
		tNftOwnerRule = append(tNftOwnerRule, tNftOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ValidatorRegistered", operatorRule, bNftOwnerRule, tNftOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerValidatorRegistered)
				if err := _StakingManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0x0b43d988cd5ab75ae318de41d6871d4b26efe57c3f3975331873f4dc073041fc.
//
// Solidity: event ValidatorRegistered(address indexed operator, address indexed bNftOwner, address indexed tNftOwner, uint256 validatorId, bytes validatorPubKey, string ipfsHashForEncryptedValidatorKey)
func (_StakingManager *StakingManagerFilterer) ParseValidatorRegistered(log types.Log) (*StakingManagerValidatorRegistered, error) {
	event := new(StakingManagerValidatorRegistered)
	if err := _StakingManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerLinkLegacyValidatorIdIterator is returned from FilterLinkLegacyValidatorId and is used to iterate over the raw logs and unpacked data for LinkLegacyValidatorId events raised by the StakingManager contract.
type StakingManagerLinkLegacyValidatorIdIterator struct {
	Event *StakingManagerLinkLegacyValidatorId // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerLinkLegacyValidatorIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerLinkLegacyValidatorId)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerLinkLegacyValidatorId)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerLinkLegacyValidatorIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerLinkLegacyValidatorIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerLinkLegacyValidatorId represents a LinkLegacyValidatorId event raised by the StakingManager contract.
type StakingManagerLinkLegacyValidatorId struct {
	PubkeyHash [32]byte
	LegacyId   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLinkLegacyValidatorId is a free log retrieval operation binding the contract event 0x694b36e569cec3c2b5f38987ad85ac588f196dc235cf2610ba2fd737fdea7e41.
//
// Solidity: event linkLegacyValidatorId(bytes32 indexed pubkeyHash, uint256 indexed legacyId)
func (_StakingManager *StakingManagerFilterer) FilterLinkLegacyValidatorId(opts *bind.FilterOpts, pubkeyHash [][32]byte, legacyId []*big.Int) (*StakingManagerLinkLegacyValidatorIdIterator, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var legacyIdRule []interface{}
	for _, legacyIdItem := range legacyId {
		legacyIdRule = append(legacyIdRule, legacyIdItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "linkLegacyValidatorId", pubkeyHashRule, legacyIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerLinkLegacyValidatorIdIterator{contract: _StakingManager.contract, event: "linkLegacyValidatorId", logs: logs, sub: sub}, nil
}

// WatchLinkLegacyValidatorId is a free log subscription operation binding the contract event 0x694b36e569cec3c2b5f38987ad85ac588f196dc235cf2610ba2fd737fdea7e41.
//
// Solidity: event linkLegacyValidatorId(bytes32 indexed pubkeyHash, uint256 indexed legacyId)
func (_StakingManager *StakingManagerFilterer) WatchLinkLegacyValidatorId(opts *bind.WatchOpts, sink chan<- *StakingManagerLinkLegacyValidatorId, pubkeyHash [][32]byte, legacyId []*big.Int) (event.Subscription, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var legacyIdRule []interface{}
	for _, legacyIdItem := range legacyId {
		legacyIdRule = append(legacyIdRule, legacyIdItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "linkLegacyValidatorId", pubkeyHashRule, legacyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerLinkLegacyValidatorId)
				if err := _StakingManager.contract.UnpackLog(event, "linkLegacyValidatorId", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLinkLegacyValidatorId is a log parse operation binding the contract event 0x694b36e569cec3c2b5f38987ad85ac588f196dc235cf2610ba2fd737fdea7e41.
//
// Solidity: event linkLegacyValidatorId(bytes32 indexed pubkeyHash, uint256 indexed legacyId)
func (_StakingManager *StakingManagerFilterer) ParseLinkLegacyValidatorId(log types.Log) (*StakingManagerLinkLegacyValidatorId, error) {
	event := new(StakingManagerLinkLegacyValidatorId)
	if err := _StakingManager.contract.UnpackLog(event, "linkLegacyValidatorId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerValidatorConfirmedIterator is returned from FilterValidatorConfirmed and is used to iterate over the raw logs and unpacked data for ValidatorConfirmed events raised by the StakingManager contract.
type StakingManagerValidatorConfirmedIterator struct {
	Event *StakingManagerValidatorConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerValidatorConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerValidatorConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerValidatorConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerValidatorConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerValidatorConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerValidatorConfirmed represents a ValidatorConfirmed event raised by the StakingManager contract.
type StakingManagerValidatorConfirmed struct {
	PubkeyHash    [32]byte
	BnftRecipient common.Address
	TnftRecipient common.Address
	Pubkey        []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorConfirmed is a free log retrieval operation binding the contract event 0x8150d115af2aea30a11022232987b82d5e748d48efba6addc26c07377c399a48.
//
// Solidity: event validatorConfirmed(bytes32 indexed pubkeyHash, address indexed bnftRecipient, address indexed tnftRecipient, bytes pubkey)
func (_StakingManager *StakingManagerFilterer) FilterValidatorConfirmed(opts *bind.FilterOpts, pubkeyHash [][32]byte, bnftRecipient []common.Address, tnftRecipient []common.Address) (*StakingManagerValidatorConfirmedIterator, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var bnftRecipientRule []interface{}
	for _, bnftRecipientItem := range bnftRecipient {
		bnftRecipientRule = append(bnftRecipientRule, bnftRecipientItem)
	}
	var tnftRecipientRule []interface{}
	for _, tnftRecipientItem := range tnftRecipient {
		tnftRecipientRule = append(tnftRecipientRule, tnftRecipientItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "validatorConfirmed", pubkeyHashRule, bnftRecipientRule, tnftRecipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerValidatorConfirmedIterator{contract: _StakingManager.contract, event: "validatorConfirmed", logs: logs, sub: sub}, nil
}

// WatchValidatorConfirmed is a free log subscription operation binding the contract event 0x8150d115af2aea30a11022232987b82d5e748d48efba6addc26c07377c399a48.
//
// Solidity: event validatorConfirmed(bytes32 indexed pubkeyHash, address indexed bnftRecipient, address indexed tnftRecipient, bytes pubkey)
func (_StakingManager *StakingManagerFilterer) WatchValidatorConfirmed(opts *bind.WatchOpts, sink chan<- *StakingManagerValidatorConfirmed, pubkeyHash [][32]byte, bnftRecipient []common.Address, tnftRecipient []common.Address) (event.Subscription, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var bnftRecipientRule []interface{}
	for _, bnftRecipientItem := range bnftRecipient {
		bnftRecipientRule = append(bnftRecipientRule, bnftRecipientItem)
	}
	var tnftRecipientRule []interface{}
	for _, tnftRecipientItem := range tnftRecipient {
		tnftRecipientRule = append(tnftRecipientRule, tnftRecipientItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "validatorConfirmed", pubkeyHashRule, bnftRecipientRule, tnftRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerValidatorConfirmed)
				if err := _StakingManager.contract.UnpackLog(event, "validatorConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorConfirmed is a log parse operation binding the contract event 0x8150d115af2aea30a11022232987b82d5e748d48efba6addc26c07377c399a48.
//
// Solidity: event validatorConfirmed(bytes32 indexed pubkeyHash, address indexed bnftRecipient, address indexed tnftRecipient, bytes pubkey)
func (_StakingManager *StakingManagerFilterer) ParseValidatorConfirmed(log types.Log) (*StakingManagerValidatorConfirmed, error) {
	event := new(StakingManagerValidatorConfirmed)
	if err := _StakingManager.contract.UnpackLog(event, "validatorConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerValidatorCreatedIterator is returned from FilterValidatorCreated and is used to iterate over the raw logs and unpacked data for ValidatorCreated events raised by the StakingManager contract.
type StakingManagerValidatorCreatedIterator struct {
	Event *StakingManagerValidatorCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingManagerValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerValidatorCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingManagerValidatorCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingManagerValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerValidatorCreated represents a ValidatorCreated event raised by the StakingManager contract.
type StakingManagerValidatorCreated struct {
	PubkeyHash  [32]byte
	EtherFiNode common.Address
	Pubkey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorCreated is a free log retrieval operation binding the contract event 0x145257af791e94a9baa6d143d80bb91506202de6e8b49814e221785063248a1b.
//
// Solidity: event validatorCreated(bytes32 indexed pubkeyHash, address indexed etherFiNode, bytes pubkey)
func (_StakingManager *StakingManagerFilterer) FilterValidatorCreated(opts *bind.FilterOpts, pubkeyHash [][32]byte, etherFiNode []common.Address) (*StakingManagerValidatorCreatedIterator, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "validatorCreated", pubkeyHashRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerValidatorCreatedIterator{contract: _StakingManager.contract, event: "validatorCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorCreated is a free log subscription operation binding the contract event 0x145257af791e94a9baa6d143d80bb91506202de6e8b49814e221785063248a1b.
//
// Solidity: event validatorCreated(bytes32 indexed pubkeyHash, address indexed etherFiNode, bytes pubkey)
func (_StakingManager *StakingManagerFilterer) WatchValidatorCreated(opts *bind.WatchOpts, sink chan<- *StakingManagerValidatorCreated, pubkeyHash [][32]byte, etherFiNode []common.Address) (event.Subscription, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "validatorCreated", pubkeyHashRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerValidatorCreated)
				if err := _StakingManager.contract.UnpackLog(event, "validatorCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorCreated is a log parse operation binding the contract event 0x145257af791e94a9baa6d143d80bb91506202de6e8b49814e221785063248a1b.
//
// Solidity: event validatorCreated(bytes32 indexed pubkeyHash, address indexed etherFiNode, bytes pubkey)
func (_StakingManager *StakingManagerFilterer) ParseValidatorCreated(log types.Log) (*StakingManagerValidatorCreated, error) {
	event := new(StakingManagerValidatorCreated)
	if err := _StakingManager.contract.UnpackLog(event, "validatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
