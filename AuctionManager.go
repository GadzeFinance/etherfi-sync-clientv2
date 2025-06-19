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

// AuctionManagerMetaData contains all meta data concerning the AuctionManager contract.
var AuctionManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEPRECATED_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPRECATED_protocolRevenueManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProtocolRevenueManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accumulatedRevenue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accumulatedRevenueThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bids\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bidderPubKeyIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bidderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelBid\",\"inputs\":[{\"name\":\"_bidId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelBidBatch\",\"inputs\":[{\"name\":\"_bidIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createBid\",\"inputs\":[{\"name\":\"_bidSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_bidAmountPerBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disableWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBidOwner\",\"inputs\":[{\"name\":\"_bidId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_nodeOperatorManagerContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeOnUpgrade\",\"inputs\":[{\"name\":\"_membershipManagerContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_accumulatedRevenueThreshold\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"_etherFiAdminContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nodeOperatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBidActive\",\"inputs\":[{\"name\":\"_bidId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxBidAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"membershipManagerContractAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minBidAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeOperatorManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINodeOperatorManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfActiveBids\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfBids\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processAuctionFeeTransfer\",\"inputs\":[{\"name\":\"_bidId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reEnterAuction\",\"inputs\":[{\"name\":\"_bidId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAccumulatedRevenueThreshold\",\"inputs\":[{\"name\":\"_newThreshold\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxBidPrice\",\"inputs\":[{\"name\":\"_newMaxBidAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBidPrice\",\"inputs\":[{\"name\":\"_newMinBidAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakingManagerContractAddress\",\"inputs\":[{\"name\":\"_stakingManagerContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingManagerContractAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferAccumulatedRevenue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unPauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAdmin\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isAdmin\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeOperatorManager\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSelectedBidInformation\",\"inputs\":[{\"name\":\"_bidId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWhitelistMinBidAmount\",\"inputs\":[{\"name\":\"_newAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"whitelistBidAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidCancelled\",\"inputs\":[{\"name\":\"bidId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidCreated\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountPerBid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bidIdArray\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"ipfsIndexArray\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidReEnteredAuction\",\"inputs\":[{\"name\":\"bidId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistDisabled\",\"inputs\":[{\"name\":\"whitelistStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistEnabled\",\"inputs\":[{\"name\":\"whitelistStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// AuctionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionManagerMetaData.ABI instead.
var AuctionManagerABI = AuctionManagerMetaData.ABI

// AuctionManager is an auto generated Go binding around an Ethereum contract.
type AuctionManager struct {
	AuctionManagerCaller     // Read-only binding to the contract
	AuctionManagerTransactor // Write-only binding to the contract
	AuctionManagerFilterer   // Log filterer for contract events
}

// AuctionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionManagerSession struct {
	Contract     *AuctionManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionManagerCallerSession struct {
	Contract *AuctionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AuctionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionManagerTransactorSession struct {
	Contract     *AuctionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AuctionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionManagerRaw struct {
	Contract *AuctionManager // Generic contract binding to access the raw methods on
}

// AuctionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionManagerCallerRaw struct {
	Contract *AuctionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionManagerTransactorRaw struct {
	Contract *AuctionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionManager creates a new instance of AuctionManager, bound to a specific deployed contract.
func NewAuctionManager(address common.Address, backend bind.ContractBackend) (*AuctionManager, error) {
	contract, err := bindAuctionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionManager{AuctionManagerCaller: AuctionManagerCaller{contract: contract}, AuctionManagerTransactor: AuctionManagerTransactor{contract: contract}, AuctionManagerFilterer: AuctionManagerFilterer{contract: contract}}, nil
}

// NewAuctionManagerCaller creates a new read-only instance of AuctionManager, bound to a specific deployed contract.
func NewAuctionManagerCaller(address common.Address, caller bind.ContractCaller) (*AuctionManagerCaller, error) {
	contract, err := bindAuctionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerCaller{contract: contract}, nil
}

// NewAuctionManagerTransactor creates a new write-only instance of AuctionManager, bound to a specific deployed contract.
func NewAuctionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionManagerTransactor, error) {
	contract, err := bindAuctionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerTransactor{contract: contract}, nil
}

// NewAuctionManagerFilterer creates a new log filterer instance of AuctionManager, bound to a specific deployed contract.
func NewAuctionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionManagerFilterer, error) {
	contract, err := bindAuctionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerFilterer{contract: contract}, nil
}

// bindAuctionManager binds a generic wrapper to an already deployed contract.
func bindAuctionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionManager *AuctionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionManager.Contract.AuctionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionManager *AuctionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.Contract.AuctionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionManager *AuctionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionManager.Contract.AuctionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionManager *AuctionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionManager *AuctionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionManager *AuctionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionManager.Contract.contract.Transact(opts, method, params...)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_AuctionManager *AuctionManagerCaller) DEPRECATEDAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "DEPRECATED_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_AuctionManager *AuctionManagerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _AuctionManager.Contract.DEPRECATEDAdmin(&_AuctionManager.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _AuctionManager.Contract.DEPRECATEDAdmin(&_AuctionManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_AuctionManager *AuctionManagerCaller) DEPRECATEDProtocolRevenueManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "DEPRECATED_protocolRevenueManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_AuctionManager *AuctionManagerSession) DEPRECATEDProtocolRevenueManager() (common.Address, error) {
	return _AuctionManager.Contract.DEPRECATEDProtocolRevenueManager(&_AuctionManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) DEPRECATEDProtocolRevenueManager() (common.Address, error) {
	return _AuctionManager.Contract.DEPRECATEDProtocolRevenueManager(&_AuctionManager.CallOpts)
}

// AccumulatedRevenue is a free data retrieval call binding the contract method 0xffacef08.
//
// Solidity: function accumulatedRevenue() view returns(uint128)
func (_AuctionManager *AuctionManagerCaller) AccumulatedRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "accumulatedRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRevenue is a free data retrieval call binding the contract method 0xffacef08.
//
// Solidity: function accumulatedRevenue() view returns(uint128)
func (_AuctionManager *AuctionManagerSession) AccumulatedRevenue() (*big.Int, error) {
	return _AuctionManager.Contract.AccumulatedRevenue(&_AuctionManager.CallOpts)
}

// AccumulatedRevenue is a free data retrieval call binding the contract method 0xffacef08.
//
// Solidity: function accumulatedRevenue() view returns(uint128)
func (_AuctionManager *AuctionManagerCallerSession) AccumulatedRevenue() (*big.Int, error) {
	return _AuctionManager.Contract.AccumulatedRevenue(&_AuctionManager.CallOpts)
}

// AccumulatedRevenueThreshold is a free data retrieval call binding the contract method 0x489b677c.
//
// Solidity: function accumulatedRevenueThreshold() view returns(uint128)
func (_AuctionManager *AuctionManagerCaller) AccumulatedRevenueThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "accumulatedRevenueThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRevenueThreshold is a free data retrieval call binding the contract method 0x489b677c.
//
// Solidity: function accumulatedRevenueThreshold() view returns(uint128)
func (_AuctionManager *AuctionManagerSession) AccumulatedRevenueThreshold() (*big.Int, error) {
	return _AuctionManager.Contract.AccumulatedRevenueThreshold(&_AuctionManager.CallOpts)
}

// AccumulatedRevenueThreshold is a free data retrieval call binding the contract method 0x489b677c.
//
// Solidity: function accumulatedRevenueThreshold() view returns(uint128)
func (_AuctionManager *AuctionManagerCallerSession) AccumulatedRevenueThreshold() (*big.Int, error) {
	return _AuctionManager.Contract.AccumulatedRevenueThreshold(&_AuctionManager.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AuctionManager *AuctionManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AuctionManager *AuctionManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _AuctionManager.Contract.Admins(&_AuctionManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AuctionManager *AuctionManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _AuctionManager.Contract.Admins(&_AuctionManager.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 amount, uint64 bidderPubKeyIndex, address bidderAddress, bool isActive)
func (_AuctionManager *AuctionManagerCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount            *big.Int
	BidderPubKeyIndex uint64
	BidderAddress     common.Address
	IsActive          bool
}, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "bids", arg0)

	outstruct := new(struct {
		Amount            *big.Int
		BidderPubKeyIndex uint64
		BidderAddress     common.Address
		IsActive          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidderPubKeyIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BidderAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 amount, uint64 bidderPubKeyIndex, address bidderAddress, bool isActive)
func (_AuctionManager *AuctionManagerSession) Bids(arg0 *big.Int) (struct {
	Amount            *big.Int
	BidderPubKeyIndex uint64
	BidderAddress     common.Address
	IsActive          bool
}, error) {
	return _AuctionManager.Contract.Bids(&_AuctionManager.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 amount, uint64 bidderPubKeyIndex, address bidderAddress, bool isActive)
func (_AuctionManager *AuctionManagerCallerSession) Bids(arg0 *big.Int) (struct {
	Amount            *big.Int
	BidderPubKeyIndex uint64
	BidderAddress     common.Address
	IsActive          bool
}, error) {
	return _AuctionManager.Contract.Bids(&_AuctionManager.CallOpts, arg0)
}

// GetBidOwner is a free data retrieval call binding the contract method 0x860e4784.
//
// Solidity: function getBidOwner(uint256 _bidId) view returns(address)
func (_AuctionManager *AuctionManagerCaller) GetBidOwner(opts *bind.CallOpts, _bidId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "getBidOwner", _bidId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidOwner is a free data retrieval call binding the contract method 0x860e4784.
//
// Solidity: function getBidOwner(uint256 _bidId) view returns(address)
func (_AuctionManager *AuctionManagerSession) GetBidOwner(_bidId *big.Int) (common.Address, error) {
	return _AuctionManager.Contract.GetBidOwner(&_AuctionManager.CallOpts, _bidId)
}

// GetBidOwner is a free data retrieval call binding the contract method 0x860e4784.
//
// Solidity: function getBidOwner(uint256 _bidId) view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) GetBidOwner(_bidId *big.Int) (common.Address, error) {
	return _AuctionManager.Contract.GetBidOwner(&_AuctionManager.CallOpts, _bidId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AuctionManager *AuctionManagerCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AuctionManager *AuctionManagerSession) GetImplementation() (common.Address, error) {
	return _AuctionManager.Contract.GetImplementation(&_AuctionManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) GetImplementation() (common.Address, error) {
	return _AuctionManager.Contract.GetImplementation(&_AuctionManager.CallOpts)
}

// IsBidActive is a free data retrieval call binding the contract method 0xa06287e6.
//
// Solidity: function isBidActive(uint256 _bidId) view returns(bool)
func (_AuctionManager *AuctionManagerCaller) IsBidActive(opts *bind.CallOpts, _bidId *big.Int) (bool, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "isBidActive", _bidId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidActive is a free data retrieval call binding the contract method 0xa06287e6.
//
// Solidity: function isBidActive(uint256 _bidId) view returns(bool)
func (_AuctionManager *AuctionManagerSession) IsBidActive(_bidId *big.Int) (bool, error) {
	return _AuctionManager.Contract.IsBidActive(&_AuctionManager.CallOpts, _bidId)
}

// IsBidActive is a free data retrieval call binding the contract method 0xa06287e6.
//
// Solidity: function isBidActive(uint256 _bidId) view returns(bool)
func (_AuctionManager *AuctionManagerCallerSession) IsBidActive(_bidId *big.Int) (bool, error) {
	return _AuctionManager.Contract.IsBidActive(&_AuctionManager.CallOpts, _bidId)
}

// MaxBidAmount is a free data retrieval call binding the contract method 0x4c9981f6.
//
// Solidity: function maxBidAmount() view returns(uint64)
func (_AuctionManager *AuctionManagerCaller) MaxBidAmount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "maxBidAmount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxBidAmount is a free data retrieval call binding the contract method 0x4c9981f6.
//
// Solidity: function maxBidAmount() view returns(uint64)
func (_AuctionManager *AuctionManagerSession) MaxBidAmount() (uint64, error) {
	return _AuctionManager.Contract.MaxBidAmount(&_AuctionManager.CallOpts)
}

// MaxBidAmount is a free data retrieval call binding the contract method 0x4c9981f6.
//
// Solidity: function maxBidAmount() view returns(uint64)
func (_AuctionManager *AuctionManagerCallerSession) MaxBidAmount() (uint64, error) {
	return _AuctionManager.Contract.MaxBidAmount(&_AuctionManager.CallOpts)
}

// MembershipManagerContractAddress is a free data retrieval call binding the contract method 0x30453392.
//
// Solidity: function membershipManagerContractAddress() view returns(address)
func (_AuctionManager *AuctionManagerCaller) MembershipManagerContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "membershipManagerContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MembershipManagerContractAddress is a free data retrieval call binding the contract method 0x30453392.
//
// Solidity: function membershipManagerContractAddress() view returns(address)
func (_AuctionManager *AuctionManagerSession) MembershipManagerContractAddress() (common.Address, error) {
	return _AuctionManager.Contract.MembershipManagerContractAddress(&_AuctionManager.CallOpts)
}

// MembershipManagerContractAddress is a free data retrieval call binding the contract method 0x30453392.
//
// Solidity: function membershipManagerContractAddress() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) MembershipManagerContractAddress() (common.Address, error) {
	return _AuctionManager.Contract.MembershipManagerContractAddress(&_AuctionManager.CallOpts)
}

// MinBidAmount is a free data retrieval call binding the contract method 0x49751788.
//
// Solidity: function minBidAmount() view returns(uint64)
func (_AuctionManager *AuctionManagerCaller) MinBidAmount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "minBidAmount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinBidAmount is a free data retrieval call binding the contract method 0x49751788.
//
// Solidity: function minBidAmount() view returns(uint64)
func (_AuctionManager *AuctionManagerSession) MinBidAmount() (uint64, error) {
	return _AuctionManager.Contract.MinBidAmount(&_AuctionManager.CallOpts)
}

// MinBidAmount is a free data retrieval call binding the contract method 0x49751788.
//
// Solidity: function minBidAmount() view returns(uint64)
func (_AuctionManager *AuctionManagerCallerSession) MinBidAmount() (uint64, error) {
	return _AuctionManager.Contract.MinBidAmount(&_AuctionManager.CallOpts)
}

// NodeOperatorManager is a free data retrieval call binding the contract method 0xe30f7c4f.
//
// Solidity: function nodeOperatorManager() view returns(address)
func (_AuctionManager *AuctionManagerCaller) NodeOperatorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "nodeOperatorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorManager is a free data retrieval call binding the contract method 0xe30f7c4f.
//
// Solidity: function nodeOperatorManager() view returns(address)
func (_AuctionManager *AuctionManagerSession) NodeOperatorManager() (common.Address, error) {
	return _AuctionManager.Contract.NodeOperatorManager(&_AuctionManager.CallOpts)
}

// NodeOperatorManager is a free data retrieval call binding the contract method 0xe30f7c4f.
//
// Solidity: function nodeOperatorManager() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) NodeOperatorManager() (common.Address, error) {
	return _AuctionManager.Contract.NodeOperatorManager(&_AuctionManager.CallOpts)
}

// NumberOfActiveBids is a free data retrieval call binding the contract method 0x048d885f.
//
// Solidity: function numberOfActiveBids() view returns(uint256)
func (_AuctionManager *AuctionManagerCaller) NumberOfActiveBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "numberOfActiveBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfActiveBids is a free data retrieval call binding the contract method 0x048d885f.
//
// Solidity: function numberOfActiveBids() view returns(uint256)
func (_AuctionManager *AuctionManagerSession) NumberOfActiveBids() (*big.Int, error) {
	return _AuctionManager.Contract.NumberOfActiveBids(&_AuctionManager.CallOpts)
}

// NumberOfActiveBids is a free data retrieval call binding the contract method 0x048d885f.
//
// Solidity: function numberOfActiveBids() view returns(uint256)
func (_AuctionManager *AuctionManagerCallerSession) NumberOfActiveBids() (*big.Int, error) {
	return _AuctionManager.Contract.NumberOfActiveBids(&_AuctionManager.CallOpts)
}

// NumberOfBids is a free data retrieval call binding the contract method 0x17d531fc.
//
// Solidity: function numberOfBids() view returns(uint256)
func (_AuctionManager *AuctionManagerCaller) NumberOfBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "numberOfBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfBids is a free data retrieval call binding the contract method 0x17d531fc.
//
// Solidity: function numberOfBids() view returns(uint256)
func (_AuctionManager *AuctionManagerSession) NumberOfBids() (*big.Int, error) {
	return _AuctionManager.Contract.NumberOfBids(&_AuctionManager.CallOpts)
}

// NumberOfBids is a free data retrieval call binding the contract method 0x17d531fc.
//
// Solidity: function numberOfBids() view returns(uint256)
func (_AuctionManager *AuctionManagerCallerSession) NumberOfBids() (*big.Int, error) {
	return _AuctionManager.Contract.NumberOfBids(&_AuctionManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionManager *AuctionManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionManager *AuctionManagerSession) Owner() (common.Address, error) {
	return _AuctionManager.Contract.Owner(&_AuctionManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) Owner() (common.Address, error) {
	return _AuctionManager.Contract.Owner(&_AuctionManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionManager *AuctionManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionManager *AuctionManagerSession) Paused() (bool, error) {
	return _AuctionManager.Contract.Paused(&_AuctionManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionManager *AuctionManagerCallerSession) Paused() (bool, error) {
	return _AuctionManager.Contract.Paused(&_AuctionManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuctionManager *AuctionManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuctionManager *AuctionManagerSession) ProxiableUUID() ([32]byte, error) {
	return _AuctionManager.Contract.ProxiableUUID(&_AuctionManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuctionManager *AuctionManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AuctionManager.Contract.ProxiableUUID(&_AuctionManager.CallOpts)
}

// StakingManagerContractAddress is a free data retrieval call binding the contract method 0xa24085e1.
//
// Solidity: function stakingManagerContractAddress() view returns(address)
func (_AuctionManager *AuctionManagerCaller) StakingManagerContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "stakingManagerContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingManagerContractAddress is a free data retrieval call binding the contract method 0xa24085e1.
//
// Solidity: function stakingManagerContractAddress() view returns(address)
func (_AuctionManager *AuctionManagerSession) StakingManagerContractAddress() (common.Address, error) {
	return _AuctionManager.Contract.StakingManagerContractAddress(&_AuctionManager.CallOpts)
}

// StakingManagerContractAddress is a free data retrieval call binding the contract method 0xa24085e1.
//
// Solidity: function stakingManagerContractAddress() view returns(address)
func (_AuctionManager *AuctionManagerCallerSession) StakingManagerContractAddress() (common.Address, error) {
	return _AuctionManager.Contract.StakingManagerContractAddress(&_AuctionManager.CallOpts)
}

// WhitelistBidAmount is a free data retrieval call binding the contract method 0x8af11723.
//
// Solidity: function whitelistBidAmount() view returns(uint128)
func (_AuctionManager *AuctionManagerCaller) WhitelistBidAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "whitelistBidAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistBidAmount is a free data retrieval call binding the contract method 0x8af11723.
//
// Solidity: function whitelistBidAmount() view returns(uint128)
func (_AuctionManager *AuctionManagerSession) WhitelistBidAmount() (*big.Int, error) {
	return _AuctionManager.Contract.WhitelistBidAmount(&_AuctionManager.CallOpts)
}

// WhitelistBidAmount is a free data retrieval call binding the contract method 0x8af11723.
//
// Solidity: function whitelistBidAmount() view returns(uint128)
func (_AuctionManager *AuctionManagerCallerSession) WhitelistBidAmount() (*big.Int, error) {
	return _AuctionManager.Contract.WhitelistBidAmount(&_AuctionManager.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_AuctionManager *AuctionManagerCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuctionManager.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_AuctionManager *AuctionManagerSession) WhitelistEnabled() (bool, error) {
	return _AuctionManager.Contract.WhitelistEnabled(&_AuctionManager.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_AuctionManager *AuctionManagerCallerSession) WhitelistEnabled() (bool, error) {
	return _AuctionManager.Contract.WhitelistEnabled(&_AuctionManager.CallOpts)
}

// CancelBid is a paid mutator transaction binding the contract method 0x9703ef35.
//
// Solidity: function cancelBid(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactor) CancelBid(opts *bind.TransactOpts, _bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "cancelBid", _bidId)
}

// CancelBid is a paid mutator transaction binding the contract method 0x9703ef35.
//
// Solidity: function cancelBid(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerSession) CancelBid(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.CancelBid(&_AuctionManager.TransactOpts, _bidId)
}

// CancelBid is a paid mutator transaction binding the contract method 0x9703ef35.
//
// Solidity: function cancelBid(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactorSession) CancelBid(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.CancelBid(&_AuctionManager.TransactOpts, _bidId)
}

// CancelBidBatch is a paid mutator transaction binding the contract method 0x8270f552.
//
// Solidity: function cancelBidBatch(uint256[] _bidIds) returns()
func (_AuctionManager *AuctionManagerTransactor) CancelBidBatch(opts *bind.TransactOpts, _bidIds []*big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "cancelBidBatch", _bidIds)
}

// CancelBidBatch is a paid mutator transaction binding the contract method 0x8270f552.
//
// Solidity: function cancelBidBatch(uint256[] _bidIds) returns()
func (_AuctionManager *AuctionManagerSession) CancelBidBatch(_bidIds []*big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.CancelBidBatch(&_AuctionManager.TransactOpts, _bidIds)
}

// CancelBidBatch is a paid mutator transaction binding the contract method 0x8270f552.
//
// Solidity: function cancelBidBatch(uint256[] _bidIds) returns()
func (_AuctionManager *AuctionManagerTransactorSession) CancelBidBatch(_bidIds []*big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.CancelBidBatch(&_AuctionManager.TransactOpts, _bidIds)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 _bidSize, uint256 _bidAmountPerBid) payable returns(uint256[])
func (_AuctionManager *AuctionManagerTransactor) CreateBid(opts *bind.TransactOpts, _bidSize *big.Int, _bidAmountPerBid *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "createBid", _bidSize, _bidAmountPerBid)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 _bidSize, uint256 _bidAmountPerBid) payable returns(uint256[])
func (_AuctionManager *AuctionManagerSession) CreateBid(_bidSize *big.Int, _bidAmountPerBid *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.CreateBid(&_AuctionManager.TransactOpts, _bidSize, _bidAmountPerBid)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 _bidSize, uint256 _bidAmountPerBid) payable returns(uint256[])
func (_AuctionManager *AuctionManagerTransactorSession) CreateBid(_bidSize *big.Int, _bidAmountPerBid *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.CreateBid(&_AuctionManager.TransactOpts, _bidSize, _bidAmountPerBid)
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_AuctionManager *AuctionManagerTransactor) DisableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "disableWhitelist")
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_AuctionManager *AuctionManagerSession) DisableWhitelist() (*types.Transaction, error) {
	return _AuctionManager.Contract.DisableWhitelist(&_AuctionManager.TransactOpts)
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_AuctionManager *AuctionManagerTransactorSession) DisableWhitelist() (*types.Transaction, error) {
	return _AuctionManager.Contract.DisableWhitelist(&_AuctionManager.TransactOpts)
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_AuctionManager *AuctionManagerTransactor) EnableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "enableWhitelist")
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_AuctionManager *AuctionManagerSession) EnableWhitelist() (*types.Transaction, error) {
	return _AuctionManager.Contract.EnableWhitelist(&_AuctionManager.TransactOpts)
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_AuctionManager *AuctionManagerTransactorSession) EnableWhitelist() (*types.Transaction, error) {
	return _AuctionManager.Contract.EnableWhitelist(&_AuctionManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _nodeOperatorManagerContract) returns()
func (_AuctionManager *AuctionManagerTransactor) Initialize(opts *bind.TransactOpts, _nodeOperatorManagerContract common.Address) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "initialize", _nodeOperatorManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _nodeOperatorManagerContract) returns()
func (_AuctionManager *AuctionManagerSession) Initialize(_nodeOperatorManagerContract common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.Initialize(&_AuctionManager.TransactOpts, _nodeOperatorManagerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _nodeOperatorManagerContract) returns()
func (_AuctionManager *AuctionManagerTransactorSession) Initialize(_nodeOperatorManagerContract common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.Initialize(&_AuctionManager.TransactOpts, _nodeOperatorManagerContract)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0x792f9b76.
//
// Solidity: function initializeOnUpgrade(address _membershipManagerContractAddress, uint128 _accumulatedRevenueThreshold, address _etherFiAdminContractAddress, address _nodeOperatorManagerAddress) returns()
func (_AuctionManager *AuctionManagerTransactor) InitializeOnUpgrade(opts *bind.TransactOpts, _membershipManagerContractAddress common.Address, _accumulatedRevenueThreshold *big.Int, _etherFiAdminContractAddress common.Address, _nodeOperatorManagerAddress common.Address) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "initializeOnUpgrade", _membershipManagerContractAddress, _accumulatedRevenueThreshold, _etherFiAdminContractAddress, _nodeOperatorManagerAddress)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0x792f9b76.
//
// Solidity: function initializeOnUpgrade(address _membershipManagerContractAddress, uint128 _accumulatedRevenueThreshold, address _etherFiAdminContractAddress, address _nodeOperatorManagerAddress) returns()
func (_AuctionManager *AuctionManagerSession) InitializeOnUpgrade(_membershipManagerContractAddress common.Address, _accumulatedRevenueThreshold *big.Int, _etherFiAdminContractAddress common.Address, _nodeOperatorManagerAddress common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.InitializeOnUpgrade(&_AuctionManager.TransactOpts, _membershipManagerContractAddress, _accumulatedRevenueThreshold, _etherFiAdminContractAddress, _nodeOperatorManagerAddress)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0x792f9b76.
//
// Solidity: function initializeOnUpgrade(address _membershipManagerContractAddress, uint128 _accumulatedRevenueThreshold, address _etherFiAdminContractAddress, address _nodeOperatorManagerAddress) returns()
func (_AuctionManager *AuctionManagerTransactorSession) InitializeOnUpgrade(_membershipManagerContractAddress common.Address, _accumulatedRevenueThreshold *big.Int, _etherFiAdminContractAddress common.Address, _nodeOperatorManagerAddress common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.InitializeOnUpgrade(&_AuctionManager.TransactOpts, _membershipManagerContractAddress, _accumulatedRevenueThreshold, _etherFiAdminContractAddress, _nodeOperatorManagerAddress)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_AuctionManager *AuctionManagerTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_AuctionManager *AuctionManagerSession) PauseContract() (*types.Transaction, error) {
	return _AuctionManager.Contract.PauseContract(&_AuctionManager.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_AuctionManager *AuctionManagerTransactorSession) PauseContract() (*types.Transaction, error) {
	return _AuctionManager.Contract.PauseContract(&_AuctionManager.TransactOpts)
}

// ProcessAuctionFeeTransfer is a paid mutator transaction binding the contract method 0xd4e01f71.
//
// Solidity: function processAuctionFeeTransfer(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactor) ProcessAuctionFeeTransfer(opts *bind.TransactOpts, _bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "processAuctionFeeTransfer", _bidId)
}

// ProcessAuctionFeeTransfer is a paid mutator transaction binding the contract method 0xd4e01f71.
//
// Solidity: function processAuctionFeeTransfer(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerSession) ProcessAuctionFeeTransfer(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.ProcessAuctionFeeTransfer(&_AuctionManager.TransactOpts, _bidId)
}

// ProcessAuctionFeeTransfer is a paid mutator transaction binding the contract method 0xd4e01f71.
//
// Solidity: function processAuctionFeeTransfer(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactorSession) ProcessAuctionFeeTransfer(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.ProcessAuctionFeeTransfer(&_AuctionManager.TransactOpts, _bidId)
}

// ReEnterAuction is a paid mutator transaction binding the contract method 0x380c1ef5.
//
// Solidity: function reEnterAuction(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactor) ReEnterAuction(opts *bind.TransactOpts, _bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "reEnterAuction", _bidId)
}

// ReEnterAuction is a paid mutator transaction binding the contract method 0x380c1ef5.
//
// Solidity: function reEnterAuction(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerSession) ReEnterAuction(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.ReEnterAuction(&_AuctionManager.TransactOpts, _bidId)
}

// ReEnterAuction is a paid mutator transaction binding the contract method 0x380c1ef5.
//
// Solidity: function reEnterAuction(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactorSession) ReEnterAuction(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.ReEnterAuction(&_AuctionManager.TransactOpts, _bidId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionManager *AuctionManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionManager *AuctionManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionManager.Contract.RenounceOwnership(&_AuctionManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionManager *AuctionManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionManager.Contract.RenounceOwnership(&_AuctionManager.TransactOpts)
}

// SetAccumulatedRevenueThreshold is a paid mutator transaction binding the contract method 0xe1ce002f.
//
// Solidity: function setAccumulatedRevenueThreshold(uint128 _newThreshold) returns()
func (_AuctionManager *AuctionManagerTransactor) SetAccumulatedRevenueThreshold(opts *bind.TransactOpts, _newThreshold *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "setAccumulatedRevenueThreshold", _newThreshold)
}

// SetAccumulatedRevenueThreshold is a paid mutator transaction binding the contract method 0xe1ce002f.
//
// Solidity: function setAccumulatedRevenueThreshold(uint128 _newThreshold) returns()
func (_AuctionManager *AuctionManagerSession) SetAccumulatedRevenueThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetAccumulatedRevenueThreshold(&_AuctionManager.TransactOpts, _newThreshold)
}

// SetAccumulatedRevenueThreshold is a paid mutator transaction binding the contract method 0xe1ce002f.
//
// Solidity: function setAccumulatedRevenueThreshold(uint128 _newThreshold) returns()
func (_AuctionManager *AuctionManagerTransactorSession) SetAccumulatedRevenueThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetAccumulatedRevenueThreshold(&_AuctionManager.TransactOpts, _newThreshold)
}

// SetMaxBidPrice is a paid mutator transaction binding the contract method 0x7edc80fd.
//
// Solidity: function setMaxBidPrice(uint64 _newMaxBidAmount) returns()
func (_AuctionManager *AuctionManagerTransactor) SetMaxBidPrice(opts *bind.TransactOpts, _newMaxBidAmount uint64) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "setMaxBidPrice", _newMaxBidAmount)
}

// SetMaxBidPrice is a paid mutator transaction binding the contract method 0x7edc80fd.
//
// Solidity: function setMaxBidPrice(uint64 _newMaxBidAmount) returns()
func (_AuctionManager *AuctionManagerSession) SetMaxBidPrice(_newMaxBidAmount uint64) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetMaxBidPrice(&_AuctionManager.TransactOpts, _newMaxBidAmount)
}

// SetMaxBidPrice is a paid mutator transaction binding the contract method 0x7edc80fd.
//
// Solidity: function setMaxBidPrice(uint64 _newMaxBidAmount) returns()
func (_AuctionManager *AuctionManagerTransactorSession) SetMaxBidPrice(_newMaxBidAmount uint64) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetMaxBidPrice(&_AuctionManager.TransactOpts, _newMaxBidAmount)
}

// SetMinBidPrice is a paid mutator transaction binding the contract method 0x702685c9.
//
// Solidity: function setMinBidPrice(uint64 _newMinBidAmount) returns()
func (_AuctionManager *AuctionManagerTransactor) SetMinBidPrice(opts *bind.TransactOpts, _newMinBidAmount uint64) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "setMinBidPrice", _newMinBidAmount)
}

// SetMinBidPrice is a paid mutator transaction binding the contract method 0x702685c9.
//
// Solidity: function setMinBidPrice(uint64 _newMinBidAmount) returns()
func (_AuctionManager *AuctionManagerSession) SetMinBidPrice(_newMinBidAmount uint64) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetMinBidPrice(&_AuctionManager.TransactOpts, _newMinBidAmount)
}

// SetMinBidPrice is a paid mutator transaction binding the contract method 0x702685c9.
//
// Solidity: function setMinBidPrice(uint64 _newMinBidAmount) returns()
func (_AuctionManager *AuctionManagerTransactorSession) SetMinBidPrice(_newMinBidAmount uint64) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetMinBidPrice(&_AuctionManager.TransactOpts, _newMinBidAmount)
}

// SetStakingManagerContractAddress is a paid mutator transaction binding the contract method 0xdbdcedd2.
//
// Solidity: function setStakingManagerContractAddress(address _stakingManagerContractAddress) returns()
func (_AuctionManager *AuctionManagerTransactor) SetStakingManagerContractAddress(opts *bind.TransactOpts, _stakingManagerContractAddress common.Address) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "setStakingManagerContractAddress", _stakingManagerContractAddress)
}

// SetStakingManagerContractAddress is a paid mutator transaction binding the contract method 0xdbdcedd2.
//
// Solidity: function setStakingManagerContractAddress(address _stakingManagerContractAddress) returns()
func (_AuctionManager *AuctionManagerSession) SetStakingManagerContractAddress(_stakingManagerContractAddress common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetStakingManagerContractAddress(&_AuctionManager.TransactOpts, _stakingManagerContractAddress)
}

// SetStakingManagerContractAddress is a paid mutator transaction binding the contract method 0xdbdcedd2.
//
// Solidity: function setStakingManagerContractAddress(address _stakingManagerContractAddress) returns()
func (_AuctionManager *AuctionManagerTransactorSession) SetStakingManagerContractAddress(_stakingManagerContractAddress common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.SetStakingManagerContractAddress(&_AuctionManager.TransactOpts, _stakingManagerContractAddress)
}

// TransferAccumulatedRevenue is a paid mutator transaction binding the contract method 0x861fe104.
//
// Solidity: function transferAccumulatedRevenue() returns()
func (_AuctionManager *AuctionManagerTransactor) TransferAccumulatedRevenue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "transferAccumulatedRevenue")
}

// TransferAccumulatedRevenue is a paid mutator transaction binding the contract method 0x861fe104.
//
// Solidity: function transferAccumulatedRevenue() returns()
func (_AuctionManager *AuctionManagerSession) TransferAccumulatedRevenue() (*types.Transaction, error) {
	return _AuctionManager.Contract.TransferAccumulatedRevenue(&_AuctionManager.TransactOpts)
}

// TransferAccumulatedRevenue is a paid mutator transaction binding the contract method 0x861fe104.
//
// Solidity: function transferAccumulatedRevenue() returns()
func (_AuctionManager *AuctionManagerTransactorSession) TransferAccumulatedRevenue() (*types.Transaction, error) {
	return _AuctionManager.Contract.TransferAccumulatedRevenue(&_AuctionManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionManager *AuctionManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionManager *AuctionManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.TransferOwnership(&_AuctionManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionManager *AuctionManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.TransferOwnership(&_AuctionManager.TransactOpts, newOwner)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_AuctionManager *AuctionManagerTransactor) UnPauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "unPauseContract")
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_AuctionManager *AuctionManagerSession) UnPauseContract() (*types.Transaction, error) {
	return _AuctionManager.Contract.UnPauseContract(&_AuctionManager.TransactOpts)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_AuctionManager *AuctionManagerTransactorSession) UnPauseContract() (*types.Transaction, error) {
	return _AuctionManager.Contract.UnPauseContract(&_AuctionManager.TransactOpts)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AuctionManager *AuctionManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AuctionManager *AuctionManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateAdmin(&_AuctionManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AuctionManager *AuctionManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateAdmin(&_AuctionManager.TransactOpts, _address, _isAdmin)
}

// UpdateNodeOperatorManager is a paid mutator transaction binding the contract method 0xb084eae5.
//
// Solidity: function updateNodeOperatorManager(address _address) returns()
func (_AuctionManager *AuctionManagerTransactor) UpdateNodeOperatorManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "updateNodeOperatorManager", _address)
}

// UpdateNodeOperatorManager is a paid mutator transaction binding the contract method 0xb084eae5.
//
// Solidity: function updateNodeOperatorManager(address _address) returns()
func (_AuctionManager *AuctionManagerSession) UpdateNodeOperatorManager(_address common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateNodeOperatorManager(&_AuctionManager.TransactOpts, _address)
}

// UpdateNodeOperatorManager is a paid mutator transaction binding the contract method 0xb084eae5.
//
// Solidity: function updateNodeOperatorManager(address _address) returns()
func (_AuctionManager *AuctionManagerTransactorSession) UpdateNodeOperatorManager(_address common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateNodeOperatorManager(&_AuctionManager.TransactOpts, _address)
}

// UpdateSelectedBidInformation is a paid mutator transaction binding the contract method 0xdfd269de.
//
// Solidity: function updateSelectedBidInformation(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactor) UpdateSelectedBidInformation(opts *bind.TransactOpts, _bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "updateSelectedBidInformation", _bidId)
}

// UpdateSelectedBidInformation is a paid mutator transaction binding the contract method 0xdfd269de.
//
// Solidity: function updateSelectedBidInformation(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerSession) UpdateSelectedBidInformation(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateSelectedBidInformation(&_AuctionManager.TransactOpts, _bidId)
}

// UpdateSelectedBidInformation is a paid mutator transaction binding the contract method 0xdfd269de.
//
// Solidity: function updateSelectedBidInformation(uint256 _bidId) returns()
func (_AuctionManager *AuctionManagerTransactorSession) UpdateSelectedBidInformation(_bidId *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateSelectedBidInformation(&_AuctionManager.TransactOpts, _bidId)
}

// UpdateWhitelistMinBidAmount is a paid mutator transaction binding the contract method 0x069ef863.
//
// Solidity: function updateWhitelistMinBidAmount(uint128 _newAmount) returns()
func (_AuctionManager *AuctionManagerTransactor) UpdateWhitelistMinBidAmount(opts *bind.TransactOpts, _newAmount *big.Int) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "updateWhitelistMinBidAmount", _newAmount)
}

// UpdateWhitelistMinBidAmount is a paid mutator transaction binding the contract method 0x069ef863.
//
// Solidity: function updateWhitelistMinBidAmount(uint128 _newAmount) returns()
func (_AuctionManager *AuctionManagerSession) UpdateWhitelistMinBidAmount(_newAmount *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateWhitelistMinBidAmount(&_AuctionManager.TransactOpts, _newAmount)
}

// UpdateWhitelistMinBidAmount is a paid mutator transaction binding the contract method 0x069ef863.
//
// Solidity: function updateWhitelistMinBidAmount(uint128 _newAmount) returns()
func (_AuctionManager *AuctionManagerTransactorSession) UpdateWhitelistMinBidAmount(_newAmount *big.Int) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpdateWhitelistMinBidAmount(&_AuctionManager.TransactOpts, _newAmount)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AuctionManager *AuctionManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AuctionManager *AuctionManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpgradeTo(&_AuctionManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AuctionManager *AuctionManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpgradeTo(&_AuctionManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuctionManager *AuctionManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuctionManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuctionManager *AuctionManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpgradeToAndCall(&_AuctionManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuctionManager *AuctionManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuctionManager.Contract.UpgradeToAndCall(&_AuctionManager.TransactOpts, newImplementation, data)
}

// AuctionManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AuctionManager contract.
type AuctionManagerAdminChangedIterator struct {
	Event *AuctionManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *AuctionManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerAdminChanged)
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
		it.Event = new(AuctionManagerAdminChanged)
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
func (it *AuctionManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerAdminChanged represents a AdminChanged event raised by the AuctionManager contract.
type AuctionManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AuctionManager *AuctionManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AuctionManagerAdminChangedIterator, error) {

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AuctionManagerAdminChangedIterator{contract: _AuctionManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AuctionManager *AuctionManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AuctionManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerAdminChanged)
				if err := _AuctionManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParseAdminChanged(log types.Log) (*AuctionManagerAdminChanged, error) {
	event := new(AuctionManagerAdminChanged)
	if err := _AuctionManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AuctionManager contract.
type AuctionManagerBeaconUpgradedIterator struct {
	Event *AuctionManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AuctionManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerBeaconUpgraded)
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
		it.Event = new(AuctionManagerBeaconUpgraded)
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
func (it *AuctionManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerBeaconUpgraded represents a BeaconUpgraded event raised by the AuctionManager contract.
type AuctionManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AuctionManager *AuctionManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AuctionManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerBeaconUpgradedIterator{contract: _AuctionManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AuctionManager *AuctionManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AuctionManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerBeaconUpgraded)
				if err := _AuctionManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParseBeaconUpgraded(log types.Log) (*AuctionManagerBeaconUpgraded, error) {
	event := new(AuctionManagerBeaconUpgraded)
	if err := _AuctionManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerBidCancelledIterator is returned from FilterBidCancelled and is used to iterate over the raw logs and unpacked data for BidCancelled events raised by the AuctionManager contract.
type AuctionManagerBidCancelledIterator struct {
	Event *AuctionManagerBidCancelled // Event containing the contract specifics and raw log

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
func (it *AuctionManagerBidCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerBidCancelled)
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
		it.Event = new(AuctionManagerBidCancelled)
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
func (it *AuctionManagerBidCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerBidCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerBidCancelled represents a BidCancelled event raised by the AuctionManager contract.
type AuctionManagerBidCancelled struct {
	BidId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBidCancelled is a free log retrieval operation binding the contract event 0xc1546e394b1975212fe013e7e6995653585f44e568c407d1157483f7d4b94581.
//
// Solidity: event BidCancelled(uint256 indexed bidId)
func (_AuctionManager *AuctionManagerFilterer) FilterBidCancelled(opts *bind.FilterOpts, bidId []*big.Int) (*AuctionManagerBidCancelledIterator, error) {

	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "BidCancelled", bidIdRule)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerBidCancelledIterator{contract: _AuctionManager.contract, event: "BidCancelled", logs: logs, sub: sub}, nil
}

// WatchBidCancelled is a free log subscription operation binding the contract event 0xc1546e394b1975212fe013e7e6995653585f44e568c407d1157483f7d4b94581.
//
// Solidity: event BidCancelled(uint256 indexed bidId)
func (_AuctionManager *AuctionManagerFilterer) WatchBidCancelled(opts *bind.WatchOpts, sink chan<- *AuctionManagerBidCancelled, bidId []*big.Int) (event.Subscription, error) {

	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "BidCancelled", bidIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerBidCancelled)
				if err := _AuctionManager.contract.UnpackLog(event, "BidCancelled", log); err != nil {
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

// ParseBidCancelled is a log parse operation binding the contract event 0xc1546e394b1975212fe013e7e6995653585f44e568c407d1157483f7d4b94581.
//
// Solidity: event BidCancelled(uint256 indexed bidId)
func (_AuctionManager *AuctionManagerFilterer) ParseBidCancelled(log types.Log) (*AuctionManagerBidCancelled, error) {
	event := new(AuctionManagerBidCancelled)
	if err := _AuctionManager.contract.UnpackLog(event, "BidCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerBidCreatedIterator is returned from FilterBidCreated and is used to iterate over the raw logs and unpacked data for BidCreated events raised by the AuctionManager contract.
type AuctionManagerBidCreatedIterator struct {
	Event *AuctionManagerBidCreated // Event containing the contract specifics and raw log

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
func (it *AuctionManagerBidCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerBidCreated)
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
		it.Event = new(AuctionManagerBidCreated)
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
func (it *AuctionManagerBidCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerBidCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerBidCreated represents a BidCreated event raised by the AuctionManager contract.
type AuctionManagerBidCreated struct {
	Bidder         common.Address
	AmountPerBid   *big.Int
	BidIdArray     []*big.Int
	IpfsIndexArray []uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBidCreated is a free log retrieval operation binding the contract event 0x5ac259179cc7e5e7986fdb545b3ac088788061a3924596f10ee30c42c79fcd4c.
//
// Solidity: event BidCreated(address indexed bidder, uint256 amountPerBid, uint256[] bidIdArray, uint64[] ipfsIndexArray)
func (_AuctionManager *AuctionManagerFilterer) FilterBidCreated(opts *bind.FilterOpts, bidder []common.Address) (*AuctionManagerBidCreatedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "BidCreated", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerBidCreatedIterator{contract: _AuctionManager.contract, event: "BidCreated", logs: logs, sub: sub}, nil
}

// WatchBidCreated is a free log subscription operation binding the contract event 0x5ac259179cc7e5e7986fdb545b3ac088788061a3924596f10ee30c42c79fcd4c.
//
// Solidity: event BidCreated(address indexed bidder, uint256 amountPerBid, uint256[] bidIdArray, uint64[] ipfsIndexArray)
func (_AuctionManager *AuctionManagerFilterer) WatchBidCreated(opts *bind.WatchOpts, sink chan<- *AuctionManagerBidCreated, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "BidCreated", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerBidCreated)
				if err := _AuctionManager.contract.UnpackLog(event, "BidCreated", log); err != nil {
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

// ParseBidCreated is a log parse operation binding the contract event 0x5ac259179cc7e5e7986fdb545b3ac088788061a3924596f10ee30c42c79fcd4c.
//
// Solidity: event BidCreated(address indexed bidder, uint256 amountPerBid, uint256[] bidIdArray, uint64[] ipfsIndexArray)
func (_AuctionManager *AuctionManagerFilterer) ParseBidCreated(log types.Log) (*AuctionManagerBidCreated, error) {
	event := new(AuctionManagerBidCreated)
	if err := _AuctionManager.contract.UnpackLog(event, "BidCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerBidReEnteredAuctionIterator is returned from FilterBidReEnteredAuction and is used to iterate over the raw logs and unpacked data for BidReEnteredAuction events raised by the AuctionManager contract.
type AuctionManagerBidReEnteredAuctionIterator struct {
	Event *AuctionManagerBidReEnteredAuction // Event containing the contract specifics and raw log

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
func (it *AuctionManagerBidReEnteredAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerBidReEnteredAuction)
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
		it.Event = new(AuctionManagerBidReEnteredAuction)
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
func (it *AuctionManagerBidReEnteredAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerBidReEnteredAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerBidReEnteredAuction represents a BidReEnteredAuction event raised by the AuctionManager contract.
type AuctionManagerBidReEnteredAuction struct {
	BidId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBidReEnteredAuction is a free log retrieval operation binding the contract event 0x1802028266951666513c5a864fe410f4b304e79bf39af3297ab000391b75c495.
//
// Solidity: event BidReEnteredAuction(uint256 indexed bidId)
func (_AuctionManager *AuctionManagerFilterer) FilterBidReEnteredAuction(opts *bind.FilterOpts, bidId []*big.Int) (*AuctionManagerBidReEnteredAuctionIterator, error) {

	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "BidReEnteredAuction", bidIdRule)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerBidReEnteredAuctionIterator{contract: _AuctionManager.contract, event: "BidReEnteredAuction", logs: logs, sub: sub}, nil
}

// WatchBidReEnteredAuction is a free log subscription operation binding the contract event 0x1802028266951666513c5a864fe410f4b304e79bf39af3297ab000391b75c495.
//
// Solidity: event BidReEnteredAuction(uint256 indexed bidId)
func (_AuctionManager *AuctionManagerFilterer) WatchBidReEnteredAuction(opts *bind.WatchOpts, sink chan<- *AuctionManagerBidReEnteredAuction, bidId []*big.Int) (event.Subscription, error) {

	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "BidReEnteredAuction", bidIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerBidReEnteredAuction)
				if err := _AuctionManager.contract.UnpackLog(event, "BidReEnteredAuction", log); err != nil {
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

// ParseBidReEnteredAuction is a log parse operation binding the contract event 0x1802028266951666513c5a864fe410f4b304e79bf39af3297ab000391b75c495.
//
// Solidity: event BidReEnteredAuction(uint256 indexed bidId)
func (_AuctionManager *AuctionManagerFilterer) ParseBidReEnteredAuction(log types.Log) (*AuctionManagerBidReEnteredAuction, error) {
	event := new(AuctionManagerBidReEnteredAuction)
	if err := _AuctionManager.contract.UnpackLog(event, "BidReEnteredAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AuctionManager contract.
type AuctionManagerInitializedIterator struct {
	Event *AuctionManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AuctionManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerInitialized)
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
		it.Event = new(AuctionManagerInitialized)
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
func (it *AuctionManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerInitialized represents a Initialized event raised by the AuctionManager contract.
type AuctionManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AuctionManager *AuctionManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuctionManagerInitializedIterator, error) {

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuctionManagerInitializedIterator{contract: _AuctionManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AuctionManager *AuctionManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuctionManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerInitialized)
				if err := _AuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParseInitialized(log types.Log) (*AuctionManagerInitialized, error) {
	event := new(AuctionManagerInitialized)
	if err := _AuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionManager contract.
type AuctionManagerOwnershipTransferredIterator struct {
	Event *AuctionManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerOwnershipTransferred)
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
		it.Event = new(AuctionManagerOwnershipTransferred)
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
func (it *AuctionManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionManager contract.
type AuctionManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionManager *AuctionManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerOwnershipTransferredIterator{contract: _AuctionManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionManager *AuctionManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerOwnershipTransferred)
				if err := _AuctionManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionManagerOwnershipTransferred, error) {
	event := new(AuctionManagerOwnershipTransferred)
	if err := _AuctionManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AuctionManager contract.
type AuctionManagerPausedIterator struct {
	Event *AuctionManagerPaused // Event containing the contract specifics and raw log

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
func (it *AuctionManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerPaused)
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
		it.Event = new(AuctionManagerPaused)
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
func (it *AuctionManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerPaused represents a Paused event raised by the AuctionManager contract.
type AuctionManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionManager *AuctionManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*AuctionManagerPausedIterator, error) {

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AuctionManagerPausedIterator{contract: _AuctionManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionManager *AuctionManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AuctionManagerPaused) (event.Subscription, error) {

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerPaused)
				if err := _AuctionManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParsePaused(log types.Log) (*AuctionManagerPaused, error) {
	event := new(AuctionManagerPaused)
	if err := _AuctionManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AuctionManager contract.
type AuctionManagerUnpausedIterator struct {
	Event *AuctionManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *AuctionManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerUnpaused)
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
		it.Event = new(AuctionManagerUnpaused)
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
func (it *AuctionManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerUnpaused represents a Unpaused event raised by the AuctionManager contract.
type AuctionManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionManager *AuctionManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AuctionManagerUnpausedIterator, error) {

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AuctionManagerUnpausedIterator{contract: _AuctionManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionManager *AuctionManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AuctionManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerUnpaused)
				if err := _AuctionManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParseUnpaused(log types.Log) (*AuctionManagerUnpaused, error) {
	event := new(AuctionManagerUnpaused)
	if err := _AuctionManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AuctionManager contract.
type AuctionManagerUpgradedIterator struct {
	Event *AuctionManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *AuctionManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerUpgraded)
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
		it.Event = new(AuctionManagerUpgraded)
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
func (it *AuctionManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerUpgraded represents a Upgraded event raised by the AuctionManager contract.
type AuctionManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuctionManager *AuctionManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuctionManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuctionManagerUpgradedIterator{contract: _AuctionManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuctionManager *AuctionManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuctionManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerUpgraded)
				if err := _AuctionManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AuctionManager *AuctionManagerFilterer) ParseUpgraded(log types.Log) (*AuctionManagerUpgraded, error) {
	event := new(AuctionManagerUpgraded)
	if err := _AuctionManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerWhitelistDisabledIterator is returned from FilterWhitelistDisabled and is used to iterate over the raw logs and unpacked data for WhitelistDisabled events raised by the AuctionManager contract.
type AuctionManagerWhitelistDisabledIterator struct {
	Event *AuctionManagerWhitelistDisabled // Event containing the contract specifics and raw log

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
func (it *AuctionManagerWhitelistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerWhitelistDisabled)
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
		it.Event = new(AuctionManagerWhitelistDisabled)
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
func (it *AuctionManagerWhitelistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerWhitelistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerWhitelistDisabled represents a WhitelistDisabled event raised by the AuctionManager contract.
type AuctionManagerWhitelistDisabled struct {
	WhitelistStatus bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWhitelistDisabled is a free log retrieval operation binding the contract event 0x3339dbd4fb585c70a2b8d61f33bc30da050195b8b986a441dd33624b88f87970.
//
// Solidity: event WhitelistDisabled(bool whitelistStatus)
func (_AuctionManager *AuctionManagerFilterer) FilterWhitelistDisabled(opts *bind.FilterOpts) (*AuctionManagerWhitelistDisabledIterator, error) {

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "WhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return &AuctionManagerWhitelistDisabledIterator{contract: _AuctionManager.contract, event: "WhitelistDisabled", logs: logs, sub: sub}, nil
}

// WatchWhitelistDisabled is a free log subscription operation binding the contract event 0x3339dbd4fb585c70a2b8d61f33bc30da050195b8b986a441dd33624b88f87970.
//
// Solidity: event WhitelistDisabled(bool whitelistStatus)
func (_AuctionManager *AuctionManagerFilterer) WatchWhitelistDisabled(opts *bind.WatchOpts, sink chan<- *AuctionManagerWhitelistDisabled) (event.Subscription, error) {

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "WhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerWhitelistDisabled)
				if err := _AuctionManager.contract.UnpackLog(event, "WhitelistDisabled", log); err != nil {
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

// ParseWhitelistDisabled is a log parse operation binding the contract event 0x3339dbd4fb585c70a2b8d61f33bc30da050195b8b986a441dd33624b88f87970.
//
// Solidity: event WhitelistDisabled(bool whitelistStatus)
func (_AuctionManager *AuctionManagerFilterer) ParseWhitelistDisabled(log types.Log) (*AuctionManagerWhitelistDisabled, error) {
	event := new(AuctionManagerWhitelistDisabled)
	if err := _AuctionManager.contract.UnpackLog(event, "WhitelistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionManagerWhitelistEnabledIterator is returned from FilterWhitelistEnabled and is used to iterate over the raw logs and unpacked data for WhitelistEnabled events raised by the AuctionManager contract.
type AuctionManagerWhitelistEnabledIterator struct {
	Event *AuctionManagerWhitelistEnabled // Event containing the contract specifics and raw log

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
func (it *AuctionManagerWhitelistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionManagerWhitelistEnabled)
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
		it.Event = new(AuctionManagerWhitelistEnabled)
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
func (it *AuctionManagerWhitelistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionManagerWhitelistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionManagerWhitelistEnabled represents a WhitelistEnabled event raised by the AuctionManager contract.
type AuctionManagerWhitelistEnabled struct {
	WhitelistStatus bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWhitelistEnabled is a free log retrieval operation binding the contract event 0x9bea0dd3cae4438dc4c54c3110002aedc380f4075b6edae73ae0536105a2008a.
//
// Solidity: event WhitelistEnabled(bool whitelistStatus)
func (_AuctionManager *AuctionManagerFilterer) FilterWhitelistEnabled(opts *bind.FilterOpts) (*AuctionManagerWhitelistEnabledIterator, error) {

	logs, sub, err := _AuctionManager.contract.FilterLogs(opts, "WhitelistEnabled")
	if err != nil {
		return nil, err
	}
	return &AuctionManagerWhitelistEnabledIterator{contract: _AuctionManager.contract, event: "WhitelistEnabled", logs: logs, sub: sub}, nil
}

// WatchWhitelistEnabled is a free log subscription operation binding the contract event 0x9bea0dd3cae4438dc4c54c3110002aedc380f4075b6edae73ae0536105a2008a.
//
// Solidity: event WhitelistEnabled(bool whitelistStatus)
func (_AuctionManager *AuctionManagerFilterer) WatchWhitelistEnabled(opts *bind.WatchOpts, sink chan<- *AuctionManagerWhitelistEnabled) (event.Subscription, error) {

	logs, sub, err := _AuctionManager.contract.WatchLogs(opts, "WhitelistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionManagerWhitelistEnabled)
				if err := _AuctionManager.contract.UnpackLog(event, "WhitelistEnabled", log); err != nil {
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

// ParseWhitelistEnabled is a log parse operation binding the contract event 0x9bea0dd3cae4438dc4c54c3110002aedc380f4075b6edae73ae0536105a2008a.
//
// Solidity: event WhitelistEnabled(bool whitelistStatus)
func (_AuctionManager *AuctionManagerFilterer) ParseWhitelistEnabled(log types.Log) (*AuctionManagerWhitelistEnabled, error) {
	event := new(AuctionManagerWhitelistEnabled)
	if err := _AuctionManager.contract.UnpackLog(event, "WhitelistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
