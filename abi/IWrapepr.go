// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IWrapperABI is the input ABI used to generate the binding from.
const IWrapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWrapperFuncSigs maps the 4-byte function signature to its string representation.
var IWrapperFuncSigs = map[string]string{
	"60de1a9b": "lock(address,uint64,bytes,uint256,uint256,uint256)",
}

// IWrapperBin is the compiled bytecode used for deploying new contracts.
var IWrapperBin = "0x608060405234801561001057600080fd5b50610140806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806360de1a9b14610030575b600080fd5b610100600480360360c081101561004657600080fd5b6001600160a01b038235169167ffffffffffffffff6020820135169181019060608101604082013564010000000081111561008057600080fd5b82018360208201111561009257600080fd5b803590602001918460018302840111640100000000831117156100b457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060400135610102565b005b50505050505056fea26469706673582212201040952f5e72944c05f425134e3d514df4bc3c9b94e2cddad4da3c6e28f3623464736f6c634300060c0033"

// DeployIWrapper deploys a new Ethereum contract, binding an instance of IWrapper to it.
func DeployIWrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IWrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(IWrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IWrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IWrapper{IWrapperCaller: IWrapperCaller{contract: contract}, IWrapperTransactor: IWrapperTransactor{contract: contract}, IWrapperFilterer: IWrapperFilterer{contract: contract}}, nil
}

// IWrapper is an auto generated Go binding around an Ethereum contract.
type IWrapper struct {
	IWrapperCaller     // Read-only binding to the contract
	IWrapperTransactor // Write-only binding to the contract
	IWrapperFilterer   // Log filterer for contract events
}

// IWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWrapperSession struct {
	Contract     *IWrapper         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWrapperCallerSession struct {
	Contract *IWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWrapperTransactorSession struct {
	Contract     *IWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWrapperRaw struct {
	Contract *IWrapper // Generic contract binding to access the raw methods on
}

// IWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWrapperCallerRaw struct {
	Contract *IWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// IWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWrapperTransactorRaw struct {
	Contract *IWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWrapper creates a new instance of IWrapper, bound to a specific deployed contract.
func NewIWrapper(address common.Address, backend bind.ContractBackend) (*IWrapper, error) {
	contract, err := bindIWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWrapper{IWrapperCaller: IWrapperCaller{contract: contract}, IWrapperTransactor: IWrapperTransactor{contract: contract}, IWrapperFilterer: IWrapperFilterer{contract: contract}}, nil
}

// NewIWrapperCaller creates a new read-only instance of IWrapper, bound to a specific deployed contract.
func NewIWrapperCaller(address common.Address, caller bind.ContractCaller) (*IWrapperCaller, error) {
	contract, err := bindIWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWrapperCaller{contract: contract}, nil
}

// NewIWrapperTransactor creates a new write-only instance of IWrapper, bound to a specific deployed contract.
func NewIWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IWrapperTransactor, error) {
	contract, err := bindIWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWrapperTransactor{contract: contract}, nil
}

// NewIWrapperFilterer creates a new log filterer instance of IWrapper, bound to a specific deployed contract.
func NewIWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IWrapperFilterer, error) {
	contract, err := bindIWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWrapperFilterer{contract: contract}, nil
}

// bindIWrapper binds a generic wrapper to an already deployed contract.
func bindIWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrapper *IWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrapper.Contract.IWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrapper *IWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrapper.Contract.IWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrapper *IWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrapper.Contract.IWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrapper *IWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrapper *IWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrapper *IWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrapper.Contract.contract.Transact(opts, method, params...)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_IWrapper *IWrapperTransactor) Lock(opts *bind.TransactOpts, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IWrapper.contract.Transact(opts, "lock", fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_IWrapper *IWrapperSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IWrapper.Contract.Lock(&_IWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_IWrapper *IWrapperTransactorSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IWrapper.Contract.Lock(&_IWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

