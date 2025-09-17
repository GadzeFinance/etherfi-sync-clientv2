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

// BeaconChainProofsBalanceContainerProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceContainerProof struct {
	BalanceContainerRoot [32]byte
	Proof                []byte
}

// BeaconChainProofsBalanceProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceProof struct {
	PubkeyHash  [32]byte
	BalanceRoot [32]byte
	Proof       []byte
}

// IDelegationManagerTypesQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesQueuedWithdrawalParams struct {
	Strategies           []common.Address
	DepositShares        []*big.Int
	DeprecatedWithdrawer common.Address
}

// IDelegationManagerTypesWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesWithdrawal struct {
	Staker       common.Address
	DelegatedTo  common.Address
	Withdrawer   common.Address
	Nonce        *big.Int
	StartBlock   uint32
	Strategies   []common.Address
	ScaledShares []*big.Int
}

// EtherFiNodesManagerMetaData contains all meta data concerning the EtherFiNodesManager contract.
var EtherFiNodesManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakingManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_roleRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ETHERFI_NODES_MANAGER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETHERFI_NODES_MANAGER_CALL_FORWARDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETHERFI_NODES_MANAGER_EIGENLAYER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressToWithdrawalCredentials\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"allowedForwardedEigenpodCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowedForwardedExternalCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateValidatorPubkeyHash\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedETHWithdrawals\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"etherFiNodeFromPubkeyHash\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEtherFiNode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"etherfiNodeAddress\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardEigenPodCall\",\"inputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardExternalCall\",\"inputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEigenPod\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkLegacyValidatorIds\",\"inputs\":[{\"name\":\"validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"pubkeys\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"linkPubkeyToNode\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nodeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"legacyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueETHWithdrawal\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"__deprecated_withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roleRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRoleRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweepFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unPauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedForwardedEigenpodCalls\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedForwardedExternalCalls\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedForwardedEigenpodCallsUpdated\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedForwardedExternalCallsUpdated\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"_target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PubkeyLinked\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"legacyId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyLinked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForwardedCallNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectRole\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidForwardedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownNode\",\"inputs\":[]}]",
}

// EtherFiNodesManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherFiNodesManagerMetaData.ABI instead.
var EtherFiNodesManagerABI = EtherFiNodesManagerMetaData.ABI

// EtherFiNodesManager is an auto generated Go binding around an Ethereum contract.
type EtherFiNodesManager struct {
	EtherFiNodesManagerCaller     // Read-only binding to the contract
	EtherFiNodesManagerTransactor // Write-only binding to the contract
	EtherFiNodesManagerFilterer   // Log filterer for contract events
}

// EtherFiNodesManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherFiNodesManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodesManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherFiNodesManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodesManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherFiNodesManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodesManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherFiNodesManagerSession struct {
	Contract     *EtherFiNodesManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EtherFiNodesManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherFiNodesManagerCallerSession struct {
	Contract *EtherFiNodesManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EtherFiNodesManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherFiNodesManagerTransactorSession struct {
	Contract     *EtherFiNodesManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EtherFiNodesManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherFiNodesManagerRaw struct {
	Contract *EtherFiNodesManager // Generic contract binding to access the raw methods on
}

// EtherFiNodesManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherFiNodesManagerCallerRaw struct {
	Contract *EtherFiNodesManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EtherFiNodesManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherFiNodesManagerTransactorRaw struct {
	Contract *EtherFiNodesManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherFiNodesManager creates a new instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManager(address common.Address, backend bind.ContractBackend) (*EtherFiNodesManager, error) {
	contract, err := bindEtherFiNodesManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManager{EtherFiNodesManagerCaller: EtherFiNodesManagerCaller{contract: contract}, EtherFiNodesManagerTransactor: EtherFiNodesManagerTransactor{contract: contract}, EtherFiNodesManagerFilterer: EtherFiNodesManagerFilterer{contract: contract}}, nil
}

// NewEtherFiNodesManagerCaller creates a new read-only instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManagerCaller(address common.Address, caller bind.ContractCaller) (*EtherFiNodesManagerCaller, error) {
	contract, err := bindEtherFiNodesManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerCaller{contract: contract}, nil
}

// NewEtherFiNodesManagerTransactor creates a new write-only instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherFiNodesManagerTransactor, error) {
	contract, err := bindEtherFiNodesManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerTransactor{contract: contract}, nil
}

// NewEtherFiNodesManagerFilterer creates a new log filterer instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherFiNodesManagerFilterer, error) {
	contract, err := bindEtherFiNodesManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerFilterer{contract: contract}, nil
}

// bindEtherFiNodesManager binds a generic wrapper to an already deployed contract.
func bindEtherFiNodesManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherFiNodesManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherFiNodesManager *EtherFiNodesManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherFiNodesManager.Contract.EtherFiNodesManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherFiNodesManager *EtherFiNodesManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.EtherFiNodesManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherFiNodesManager *EtherFiNodesManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.EtherFiNodesManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherFiNodesManager *EtherFiNodesManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherFiNodesManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.contract.Transact(opts, method, params...)
}

// ETHERFINODESMANAGERADMINROLE is a free data retrieval call binding the contract method 0xc1258f5f.
//
// Solidity: function ETHERFI_NODES_MANAGER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) ETHERFINODESMANAGERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "ETHERFI_NODES_MANAGER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHERFINODESMANAGERADMINROLE is a free data retrieval call binding the contract method 0xc1258f5f.
//
// Solidity: function ETHERFI_NODES_MANAGER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ETHERFINODESMANAGERADMINROLE() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ETHERFINODESMANAGERADMINROLE(&_EtherFiNodesManager.CallOpts)
}

// ETHERFINODESMANAGERADMINROLE is a free data retrieval call binding the contract method 0xc1258f5f.
//
// Solidity: function ETHERFI_NODES_MANAGER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) ETHERFINODESMANAGERADMINROLE() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ETHERFINODESMANAGERADMINROLE(&_EtherFiNodesManager.CallOpts)
}

// ETHERFINODESMANAGERCALLFORWARDERROLE is a free data retrieval call binding the contract method 0xc4be7c72.
//
// Solidity: function ETHERFI_NODES_MANAGER_CALL_FORWARDER_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) ETHERFINODESMANAGERCALLFORWARDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "ETHERFI_NODES_MANAGER_CALL_FORWARDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHERFINODESMANAGERCALLFORWARDERROLE is a free data retrieval call binding the contract method 0xc4be7c72.
//
// Solidity: function ETHERFI_NODES_MANAGER_CALL_FORWARDER_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ETHERFINODESMANAGERCALLFORWARDERROLE() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ETHERFINODESMANAGERCALLFORWARDERROLE(&_EtherFiNodesManager.CallOpts)
}

// ETHERFINODESMANAGERCALLFORWARDERROLE is a free data retrieval call binding the contract method 0xc4be7c72.
//
// Solidity: function ETHERFI_NODES_MANAGER_CALL_FORWARDER_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) ETHERFINODESMANAGERCALLFORWARDERROLE() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ETHERFINODESMANAGERCALLFORWARDERROLE(&_EtherFiNodesManager.CallOpts)
}

// ETHERFINODESMANAGEREIGENLAYERADMINROLE is a free data retrieval call binding the contract method 0x0dfa3b5c.
//
// Solidity: function ETHERFI_NODES_MANAGER_EIGENLAYER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) ETHERFINODESMANAGEREIGENLAYERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "ETHERFI_NODES_MANAGER_EIGENLAYER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHERFINODESMANAGEREIGENLAYERADMINROLE is a free data retrieval call binding the contract method 0x0dfa3b5c.
//
// Solidity: function ETHERFI_NODES_MANAGER_EIGENLAYER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ETHERFINODESMANAGEREIGENLAYERADMINROLE() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ETHERFINODESMANAGEREIGENLAYERADMINROLE(&_EtherFiNodesManager.CallOpts)
}

// ETHERFINODESMANAGEREIGENLAYERADMINROLE is a free data retrieval call binding the contract method 0x0dfa3b5c.
//
// Solidity: function ETHERFI_NODES_MANAGER_EIGENLAYER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) ETHERFINODESMANAGEREIGENLAYERADMINROLE() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ETHERFINODESMANAGEREIGENLAYERADMINROLE(&_EtherFiNodesManager.CallOpts)
}

// AddressToWithdrawalCredentials is a free data retrieval call binding the contract method 0x07a455d3.
//
// Solidity: function addressToWithdrawalCredentials(address addr) pure returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) AddressToWithdrawalCredentials(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "addressToWithdrawalCredentials", addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AddressToWithdrawalCredentials is a free data retrieval call binding the contract method 0x07a455d3.
//
// Solidity: function addressToWithdrawalCredentials(address addr) pure returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AddressToWithdrawalCredentials(addr common.Address) ([]byte, error) {
	return _EtherFiNodesManager.Contract.AddressToWithdrawalCredentials(&_EtherFiNodesManager.CallOpts, addr)
}

// AddressToWithdrawalCredentials is a free data retrieval call binding the contract method 0x07a455d3.
//
// Solidity: function addressToWithdrawalCredentials(address addr) pure returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) AddressToWithdrawalCredentials(addr common.Address) ([]byte, error) {
	return _EtherFiNodesManager.Contract.AddressToWithdrawalCredentials(&_EtherFiNodesManager.CallOpts, addr)
}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) AllowedForwardedEigenpodCalls(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "allowedForwardedEigenpodCalls", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AllowedForwardedEigenpodCalls(arg0 [4]byte) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedEigenpodCalls(&_EtherFiNodesManager.CallOpts, arg0)
}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) AllowedForwardedEigenpodCalls(arg0 [4]byte) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedEigenpodCalls(&_EtherFiNodesManager.CallOpts, arg0)
}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) AllowedForwardedExternalCalls(opts *bind.CallOpts, arg0 [4]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "allowedForwardedExternalCalls", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AllowedForwardedExternalCalls(arg0 [4]byte, arg1 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedExternalCalls(&_EtherFiNodesManager.CallOpts, arg0, arg1)
}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) AllowedForwardedExternalCalls(arg0 [4]byte, arg1 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedExternalCalls(&_EtherFiNodesManager.CallOpts, arg0, arg1)
}

// CalculateValidatorPubkeyHash is a free data retrieval call binding the contract method 0x49d58951.
//
// Solidity: function calculateValidatorPubkeyHash(bytes pubkey) pure returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) CalculateValidatorPubkeyHash(opts *bind.CallOpts, pubkey []byte) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "calculateValidatorPubkeyHash", pubkey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateValidatorPubkeyHash is a free data retrieval call binding the contract method 0x49d58951.
//
// Solidity: function calculateValidatorPubkeyHash(bytes pubkey) pure returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) CalculateValidatorPubkeyHash(pubkey []byte) ([32]byte, error) {
	return _EtherFiNodesManager.Contract.CalculateValidatorPubkeyHash(&_EtherFiNodesManager.CallOpts, pubkey)
}

// CalculateValidatorPubkeyHash is a free data retrieval call binding the contract method 0x49d58951.
//
// Solidity: function calculateValidatorPubkeyHash(bytes pubkey) pure returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) CalculateValidatorPubkeyHash(pubkey []byte) ([32]byte, error) {
	return _EtherFiNodesManager.Contract.CalculateValidatorPubkeyHash(&_EtherFiNodesManager.CallOpts, pubkey)
}

// EtherFiNodeFromPubkeyHash is a free data retrieval call binding the contract method 0x9055e951.
//
// Solidity: function etherFiNodeFromPubkeyHash(bytes32 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) EtherFiNodeFromPubkeyHash(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "etherFiNodeFromPubkeyHash", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherFiNodeFromPubkeyHash is a free data retrieval call binding the contract method 0x9055e951.
//
// Solidity: function etherFiNodeFromPubkeyHash(bytes32 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) EtherFiNodeFromPubkeyHash(arg0 [32]byte) (common.Address, error) {
	return _EtherFiNodesManager.Contract.EtherFiNodeFromPubkeyHash(&_EtherFiNodesManager.CallOpts, arg0)
}

// EtherFiNodeFromPubkeyHash is a free data retrieval call binding the contract method 0x9055e951.
//
// Solidity: function etherFiNodeFromPubkeyHash(bytes32 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) EtherFiNodeFromPubkeyHash(arg0 [32]byte) (common.Address, error) {
	return _EtherFiNodesManager.Contract.EtherFiNodeFromPubkeyHash(&_EtherFiNodesManager.CallOpts, arg0)
}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 id) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) EtherfiNodeAddress(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "etherfiNodeAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 id) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) EtherfiNodeAddress(id *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.EtherfiNodeAddress(&_EtherFiNodesManager.CallOpts, id)
}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 id) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) EtherfiNodeAddress(id *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.EtherfiNodeAddress(&_EtherFiNodesManager.CallOpts, id)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 id) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetEigenPod(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getEigenPod", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 id) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetEigenPod(id *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetEigenPod(&_EtherFiNodesManager.CallOpts, id)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 id) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetEigenPod(id *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetEigenPod(&_EtherFiNodesManager.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Owner() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Owner(&_EtherFiNodesManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Owner() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Owner(&_EtherFiNodesManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Paused() (bool, error) {
	return _EtherFiNodesManager.Contract.Paused(&_EtherFiNodesManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Paused() (bool, error) {
	return _EtherFiNodesManager.Contract.Paused(&_EtherFiNodesManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ProxiableUUID(&_EtherFiNodesManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ProxiableUUID(&_EtherFiNodesManager.CallOpts)
}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) RoleRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "roleRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) RoleRegistry() (common.Address, error) {
	return _EtherFiNodesManager.Contract.RoleRegistry(&_EtherFiNodesManager.CallOpts)
}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) RoleRegistry() (common.Address, error) {
	return _EtherFiNodesManager.Contract.RoleRegistry(&_EtherFiNodesManager.CallOpts)
}

// StakingManager is a free data retrieval call binding the contract method 0x22828cc2.
//
// Solidity: function stakingManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) StakingManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "stakingManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingManager is a free data retrieval call binding the contract method 0x22828cc2.
//
// Solidity: function stakingManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) StakingManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.StakingManager(&_EtherFiNodesManager.CallOpts)
}

// StakingManager is a free data retrieval call binding the contract method 0x22828cc2.
//
// Solidity: function stakingManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) StakingManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.StakingManager(&_EtherFiNodesManager.CallOpts)
}

// CompleteQueuedETHWithdrawals is a paid mutator transaction binding the contract method 0x84973b72.
//
// Solidity: function completeQueuedETHWithdrawals(uint256 id, bool receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) CompleteQueuedETHWithdrawals(opts *bind.TransactOpts, id *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "completeQueuedETHWithdrawals", id, receiveAsTokens)
}

// CompleteQueuedETHWithdrawals is a paid mutator transaction binding the contract method 0x84973b72.
//
// Solidity: function completeQueuedETHWithdrawals(uint256 id, bool receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) CompleteQueuedETHWithdrawals(id *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.CompleteQueuedETHWithdrawals(&_EtherFiNodesManager.TransactOpts, id, receiveAsTokens)
}

// CompleteQueuedETHWithdrawals is a paid mutator transaction binding the contract method 0x84973b72.
//
// Solidity: function completeQueuedETHWithdrawals(uint256 id, bool receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) CompleteQueuedETHWithdrawals(id *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.CompleteQueuedETHWithdrawals(&_EtherFiNodesManager.TransactOpts, id, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xabe0c96f.
//
// Solidity: function completeQueuedWithdrawals(uint256 id, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, id *big.Int, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "completeQueuedWithdrawals", id, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xabe0c96f.
//
// Solidity: function completeQueuedWithdrawals(uint256 id, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) CompleteQueuedWithdrawals(id *big.Int, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.CompleteQueuedWithdrawals(&_EtherFiNodesManager.TransactOpts, id, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xabe0c96f.
//
// Solidity: function completeQueuedWithdrawals(uint256 id, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) CompleteQueuedWithdrawals(id *big.Int, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.CompleteQueuedWithdrawals(&_EtherFiNodesManager.TransactOpts, id, withdrawals, tokens, receiveAsTokens)
}

// ForwardEigenPodCall is a paid mutator transaction binding the contract method 0xd67db2e5.
//
// Solidity: function forwardEigenPodCall(uint256[] ids, bytes[] data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ForwardEigenPodCall(opts *bind.TransactOpts, ids []*big.Int, data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "forwardEigenPodCall", ids, data)
}

// ForwardEigenPodCall is a paid mutator transaction binding the contract method 0xd67db2e5.
//
// Solidity: function forwardEigenPodCall(uint256[] ids, bytes[] data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ForwardEigenPodCall(ids []*big.Int, data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardEigenPodCall(&_EtherFiNodesManager.TransactOpts, ids, data)
}

// ForwardEigenPodCall is a paid mutator transaction binding the contract method 0xd67db2e5.
//
// Solidity: function forwardEigenPodCall(uint256[] ids, bytes[] data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ForwardEigenPodCall(ids []*big.Int, data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardEigenPodCall(&_EtherFiNodesManager.TransactOpts, ids, data)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] ids, bytes[] data, address target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ForwardExternalCall(opts *bind.TransactOpts, ids []*big.Int, data [][]byte, target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "forwardExternalCall", ids, data, target)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] ids, bytes[] data, address target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ForwardExternalCall(ids []*big.Int, data [][]byte, target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardExternalCall(&_EtherFiNodesManager.TransactOpts, ids, data, target)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] ids, bytes[] data, address target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ForwardExternalCall(ids []*big.Int, data [][]byte, target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardExternalCall(&_EtherFiNodesManager.TransactOpts, ids, data, target)
}

// LinkLegacyValidatorIds is a paid mutator transaction binding the contract method 0x83294396.
//
// Solidity: function linkLegacyValidatorIds(uint256[] validatorIds, bytes[] pubkeys) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) LinkLegacyValidatorIds(opts *bind.TransactOpts, validatorIds []*big.Int, pubkeys [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "linkLegacyValidatorIds", validatorIds, pubkeys)
}

// LinkLegacyValidatorIds is a paid mutator transaction binding the contract method 0x83294396.
//
// Solidity: function linkLegacyValidatorIds(uint256[] validatorIds, bytes[] pubkeys) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) LinkLegacyValidatorIds(validatorIds []*big.Int, pubkeys [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.LinkLegacyValidatorIds(&_EtherFiNodesManager.TransactOpts, validatorIds, pubkeys)
}

// LinkLegacyValidatorIds is a paid mutator transaction binding the contract method 0x83294396.
//
// Solidity: function linkLegacyValidatorIds(uint256[] validatorIds, bytes[] pubkeys) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) LinkLegacyValidatorIds(validatorIds []*big.Int, pubkeys [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.LinkLegacyValidatorIds(&_EtherFiNodesManager.TransactOpts, validatorIds, pubkeys)
}

// LinkPubkeyToNode is a paid mutator transaction binding the contract method 0x5f077d1c.
//
// Solidity: function linkPubkeyToNode(bytes pubkey, address nodeAddress, uint256 legacyId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) LinkPubkeyToNode(opts *bind.TransactOpts, pubkey []byte, nodeAddress common.Address, legacyId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "linkPubkeyToNode", pubkey, nodeAddress, legacyId)
}

// LinkPubkeyToNode is a paid mutator transaction binding the contract method 0x5f077d1c.
//
// Solidity: function linkPubkeyToNode(bytes pubkey, address nodeAddress, uint256 legacyId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) LinkPubkeyToNode(pubkey []byte, nodeAddress common.Address, legacyId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.LinkPubkeyToNode(&_EtherFiNodesManager.TransactOpts, pubkey, nodeAddress, legacyId)
}

// LinkPubkeyToNode is a paid mutator transaction binding the contract method 0x5f077d1c.
//
// Solidity: function linkPubkeyToNode(bytes pubkey, address nodeAddress, uint256 legacyId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) LinkPubkeyToNode(pubkey []byte, nodeAddress common.Address, legacyId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.LinkPubkeyToNode(&_EtherFiNodesManager.TransactOpts, pubkey, nodeAddress, legacyId)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) PauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.PauseContract(&_EtherFiNodesManager.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) PauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.PauseContract(&_EtherFiNodesManager.TransactOpts)
}

// QueueETHWithdrawal is a paid mutator transaction binding the contract method 0xc3a9e20e.
//
// Solidity: function queueETHWithdrawal(uint256 id, uint256 amount) returns(bytes32 withdrawalRoot)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) QueueETHWithdrawal(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "queueETHWithdrawal", id, amount)
}

// QueueETHWithdrawal is a paid mutator transaction binding the contract method 0xc3a9e20e.
//
// Solidity: function queueETHWithdrawal(uint256 id, uint256 amount) returns(bytes32 withdrawalRoot)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) QueueETHWithdrawal(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.QueueETHWithdrawal(&_EtherFiNodesManager.TransactOpts, id, amount)
}

// QueueETHWithdrawal is a paid mutator transaction binding the contract method 0xc3a9e20e.
//
// Solidity: function queueETHWithdrawal(uint256 id, uint256 amount) returns(bytes32 withdrawalRoot)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) QueueETHWithdrawal(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.QueueETHWithdrawal(&_EtherFiNodesManager.TransactOpts, id, amount)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0031e778.
//
// Solidity: function queueWithdrawals(uint256 id, (address[],uint256[],address)[] params) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) QueueWithdrawals(opts *bind.TransactOpts, id *big.Int, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "queueWithdrawals", id, params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0031e778.
//
// Solidity: function queueWithdrawals(uint256 id, (address[],uint256[],address)[] params) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) QueueWithdrawals(id *big.Int, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.QueueWithdrawals(&_EtherFiNodesManager.TransactOpts, id, params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0031e778.
//
// Solidity: function queueWithdrawals(uint256 id, (address[],uint256[],address)[] params) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) QueueWithdrawals(id *big.Int, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.QueueWithdrawals(&_EtherFiNodesManager.TransactOpts, id, params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.RenounceOwnership(&_EtherFiNodesManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.RenounceOwnership(&_EtherFiNodesManager.TransactOpts)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 id, address proofSubmitter) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SetProofSubmitter(opts *bind.TransactOpts, id *big.Int, proofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "setProofSubmitter", id, proofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 id, address proofSubmitter) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SetProofSubmitter(id *big.Int, proofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetProofSubmitter(&_EtherFiNodesManager.TransactOpts, id, proofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 id, address proofSubmitter) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SetProofSubmitter(id *big.Int, proofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetProofSubmitter(&_EtherFiNodesManager.TransactOpts, id, proofSubmitter)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x0567b237.
//
// Solidity: function startCheckpoint(uint256 id) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) StartCheckpoint(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "startCheckpoint", id)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x0567b237.
//
// Solidity: function startCheckpoint(uint256 id) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) StartCheckpoint(id *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.StartCheckpoint(&_EtherFiNodesManager.TransactOpts, id)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x0567b237.
//
// Solidity: function startCheckpoint(uint256 id) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) StartCheckpoint(id *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.StartCheckpoint(&_EtherFiNodesManager.TransactOpts, id)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xe5858232.
//
// Solidity: function sweepFunds(uint256 id) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SweepFunds(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "sweepFunds", id)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xe5858232.
//
// Solidity: function sweepFunds(uint256 id) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SweepFunds(id *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SweepFunds(&_EtherFiNodesManager.TransactOpts, id)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xe5858232.
//
// Solidity: function sweepFunds(uint256 id) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SweepFunds(id *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SweepFunds(&_EtherFiNodesManager.TransactOpts, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.TransferOwnership(&_EtherFiNodesManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.TransferOwnership(&_EtherFiNodesManager.TransactOpts, newOwner)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UnPauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "unPauseContract")
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UnPauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UnPauseContract(&_EtherFiNodesManager.TransactOpts)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UnPauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UnPauseContract(&_EtherFiNodesManager.TransactOpts)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 selector, bool allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateAllowedForwardedEigenpodCalls(opts *bind.TransactOpts, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateAllowedForwardedEigenpodCalls", selector, allowed)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 selector, bool allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateAllowedForwardedEigenpodCalls(selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedEigenpodCalls(&_EtherFiNodesManager.TransactOpts, selector, allowed)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 selector, bool allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateAllowedForwardedEigenpodCalls(selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedEigenpodCalls(&_EtherFiNodesManager.TransactOpts, selector, allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 selector, address target, bool allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateAllowedForwardedExternalCalls(opts *bind.TransactOpts, selector [4]byte, target common.Address, allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateAllowedForwardedExternalCalls", selector, target, allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 selector, address target, bool allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateAllowedForwardedExternalCalls(selector [4]byte, target common.Address, allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedExternalCalls(&_EtherFiNodesManager.TransactOpts, selector, target, allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 selector, address target, bool allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateAllowedForwardedExternalCalls(selector [4]byte, target common.Address, allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedExternalCalls(&_EtherFiNodesManager.TransactOpts, selector, target, allowed)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeTo(&_EtherFiNodesManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeTo(&_EtherFiNodesManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeToAndCall(&_EtherFiNodesManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeToAndCall(&_EtherFiNodesManager.TransactOpts, newImplementation, data)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0x20f3b1d4.
//
// Solidity: function verifyCheckpointProofs(uint256 id, (bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) VerifyCheckpointProofs(opts *bind.TransactOpts, id *big.Int, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "verifyCheckpointProofs", id, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0x20f3b1d4.
//
// Solidity: function verifyCheckpointProofs(uint256 id, (bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) VerifyCheckpointProofs(id *big.Int, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.VerifyCheckpointProofs(&_EtherFiNodesManager.TransactOpts, id, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0x20f3b1d4.
//
// Solidity: function verifyCheckpointProofs(uint256 id, (bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) VerifyCheckpointProofs(id *big.Int, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.VerifyCheckpointProofs(&_EtherFiNodesManager.TransactOpts, id, balanceContainerProof, proofs)
}

// EtherFiNodesManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAdminChangedIterator struct {
	Event *EtherFiNodesManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerAdminChanged)
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
		it.Event = new(EtherFiNodesManagerAdminChanged)
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
func (it *EtherFiNodesManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerAdminChanged represents a AdminChanged event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EtherFiNodesManagerAdminChangedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerAdminChangedIterator{contract: _EtherFiNodesManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerAdminChanged)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseAdminChanged(log types.Log) (*EtherFiNodesManagerAdminChanged, error) {
	event := new(EtherFiNodesManagerAdminChanged)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator is returned from FilterAllowedForwardedEigenpodCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedForwardedEigenpodCallsUpdated events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator struct {
	Event *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
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
		it.Event = new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
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
func (it *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated represents a AllowedForwardedEigenpodCallsUpdated event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated struct {
	Selector [4]byte
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedForwardedEigenpodCallsUpdated is a free log retrieval operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterAllowedForwardedEigenpodCallsUpdated(opts *bind.FilterOpts, selector [][4]byte) (*EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "AllowedForwardedEigenpodCallsUpdated", selectorRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator{contract: _EtherFiNodesManager.contract, event: "AllowedForwardedEigenpodCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedForwardedEigenpodCallsUpdated is a free log subscription operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchAllowedForwardedEigenpodCallsUpdated(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated, selector [][4]byte) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "AllowedForwardedEigenpodCallsUpdated", selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedEigenpodCallsUpdated", log); err != nil {
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

// ParseAllowedForwardedEigenpodCallsUpdated is a log parse operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseAllowedForwardedEigenpodCallsUpdated(log types.Log) (*EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated, error) {
	event := new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedEigenpodCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator is returned from FilterAllowedForwardedExternalCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedForwardedExternalCallsUpdated events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator struct {
	Event *EtherFiNodesManagerAllowedForwardedExternalCallsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
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
		it.Event = new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
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
func (it *EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerAllowedForwardedExternalCallsUpdated represents a AllowedForwardedExternalCallsUpdated event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedExternalCallsUpdated struct {
	Selector [4]byte
	Target   common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedForwardedExternalCallsUpdated is a free log retrieval operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterAllowedForwardedExternalCallsUpdated(opts *bind.FilterOpts, selector [][4]byte, _target []common.Address) (*EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "AllowedForwardedExternalCallsUpdated", selectorRule, _targetRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator{contract: _EtherFiNodesManager.contract, event: "AllowedForwardedExternalCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedForwardedExternalCallsUpdated is a free log subscription operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchAllowedForwardedExternalCallsUpdated(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerAllowedForwardedExternalCallsUpdated, selector [][4]byte, _target []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "AllowedForwardedExternalCallsUpdated", selectorRule, _targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedExternalCallsUpdated", log); err != nil {
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

// ParseAllowedForwardedExternalCallsUpdated is a log parse operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseAllowedForwardedExternalCallsUpdated(log types.Log) (*EtherFiNodesManagerAllowedForwardedExternalCallsUpdated, error) {
	event := new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedExternalCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerBeaconUpgradedIterator struct {
	Event *EtherFiNodesManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerBeaconUpgraded)
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
		it.Event = new(EtherFiNodesManagerBeaconUpgraded)
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
func (it *EtherFiNodesManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerBeaconUpgraded represents a BeaconUpgraded event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EtherFiNodesManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerBeaconUpgradedIterator{contract: _EtherFiNodesManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerBeaconUpgraded)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseBeaconUpgraded(log types.Log) (*EtherFiNodesManagerBeaconUpgraded, error) {
	event := new(EtherFiNodesManagerBeaconUpgraded)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerInitializedIterator struct {
	Event *EtherFiNodesManagerInitialized // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerInitialized)
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
		it.Event = new(EtherFiNodesManagerInitialized)
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
func (it *EtherFiNodesManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerInitialized represents a Initialized event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EtherFiNodesManagerInitializedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerInitializedIterator{contract: _EtherFiNodesManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerInitialized)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseInitialized(log types.Log) (*EtherFiNodesManagerInitialized, error) {
	event := new(EtherFiNodesManagerInitialized)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerOwnershipTransferredIterator struct {
	Event *EtherFiNodesManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerOwnershipTransferred)
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
		it.Event = new(EtherFiNodesManagerOwnershipTransferred)
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
func (it *EtherFiNodesManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EtherFiNodesManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerOwnershipTransferredIterator{contract: _EtherFiNodesManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerOwnershipTransferred)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EtherFiNodesManagerOwnershipTransferred, error) {
	event := new(EtherFiNodesManagerOwnershipTransferred)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPausedIterator struct {
	Event *EtherFiNodesManagerPaused // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerPaused)
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
		it.Event = new(EtherFiNodesManagerPaused)
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
func (it *EtherFiNodesManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerPaused represents a Paused event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EtherFiNodesManagerPausedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerPausedIterator{contract: _EtherFiNodesManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerPaused)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParsePaused(log types.Log) (*EtherFiNodesManagerPaused, error) {
	event := new(EtherFiNodesManagerPaused)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerPubkeyLinkedIterator is returned from FilterPubkeyLinked and is used to iterate over the raw logs and unpacked data for PubkeyLinked events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPubkeyLinkedIterator struct {
	Event *EtherFiNodesManagerPubkeyLinked // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerPubkeyLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerPubkeyLinked)
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
		it.Event = new(EtherFiNodesManagerPubkeyLinked)
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
func (it *EtherFiNodesManagerPubkeyLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerPubkeyLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerPubkeyLinked represents a PubkeyLinked event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPubkeyLinked struct {
	PubkeyHash  [32]byte
	NodeAddress common.Address
	LegacyId    *big.Int
	Pubkey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPubkeyLinked is a free log retrieval operation binding the contract event 0x5e525a525cf73653f769c8305dc71a68b85b0e62e3cc5258fe187ff9fd3e5cb9.
//
// Solidity: event PubkeyLinked(bytes32 indexed pubkeyHash, address indexed nodeAddress, uint256 indexed legacyId, bytes pubkey)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterPubkeyLinked(opts *bind.FilterOpts, pubkeyHash [][32]byte, nodeAddress []common.Address, legacyId []*big.Int) (*EtherFiNodesManagerPubkeyLinkedIterator, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}
	var legacyIdRule []interface{}
	for _, legacyIdItem := range legacyId {
		legacyIdRule = append(legacyIdRule, legacyIdItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "PubkeyLinked", pubkeyHashRule, nodeAddressRule, legacyIdRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerPubkeyLinkedIterator{contract: _EtherFiNodesManager.contract, event: "PubkeyLinked", logs: logs, sub: sub}, nil
}

// WatchPubkeyLinked is a free log subscription operation binding the contract event 0x5e525a525cf73653f769c8305dc71a68b85b0e62e3cc5258fe187ff9fd3e5cb9.
//
// Solidity: event PubkeyLinked(bytes32 indexed pubkeyHash, address indexed nodeAddress, uint256 indexed legacyId, bytes pubkey)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchPubkeyLinked(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerPubkeyLinked, pubkeyHash [][32]byte, nodeAddress []common.Address, legacyId []*big.Int) (event.Subscription, error) {

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}
	var legacyIdRule []interface{}
	for _, legacyIdItem := range legacyId {
		legacyIdRule = append(legacyIdRule, legacyIdItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "PubkeyLinked", pubkeyHashRule, nodeAddressRule, legacyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerPubkeyLinked)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "PubkeyLinked", log); err != nil {
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

// ParsePubkeyLinked is a log parse operation binding the contract event 0x5e525a525cf73653f769c8305dc71a68b85b0e62e3cc5258fe187ff9fd3e5cb9.
//
// Solidity: event PubkeyLinked(bytes32 indexed pubkeyHash, address indexed nodeAddress, uint256 indexed legacyId, bytes pubkey)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParsePubkeyLinked(log types.Log) (*EtherFiNodesManagerPubkeyLinked, error) {
	event := new(EtherFiNodesManagerPubkeyLinked)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "PubkeyLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUnpausedIterator struct {
	Event *EtherFiNodesManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerUnpaused)
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
		it.Event = new(EtherFiNodesManagerUnpaused)
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
func (it *EtherFiNodesManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerUnpaused represents a Unpaused event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EtherFiNodesManagerUnpausedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerUnpausedIterator{contract: _EtherFiNodesManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerUnpaused)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseUnpaused(log types.Log) (*EtherFiNodesManagerUnpaused, error) {
	event := new(EtherFiNodesManagerUnpaused)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUpgradedIterator struct {
	Event *EtherFiNodesManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerUpgraded)
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
		it.Event = new(EtherFiNodesManagerUpgraded)
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
func (it *EtherFiNodesManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerUpgraded represents a Upgraded event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EtherFiNodesManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerUpgradedIterator{contract: _EtherFiNodesManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerUpgraded)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseUpgraded(log types.Log) (*EtherFiNodesManagerUpgraded, error) {
	event := new(EtherFiNodesManagerUpgraded)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
