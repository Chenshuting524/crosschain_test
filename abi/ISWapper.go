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

// ISwapperABI is the input ABI used to generate the binding from.
const ISwapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ISwapperFuncSigs maps the 4-byte function signature to its string representation.
var ISwapperFuncSigs = map[string]string{
	"4672b083": "add_liquidity(address,uint64,uint64,bytes,uint256,uint256,uint256,uint256)",
	"773c9d9c": "remove_liquidity(address,uint64,uint64,bytes,bytes,uint256,uint256,uint256,uint256)",
	"559ebe57": "swap(address,uint64,uint64,bytes,bytes,uint256,uint256,uint256,uint256)",
}

// ISwapperBin is the compiled bytecode used for deploying new contracts.
var ISwapperBin = "0x608060405234801561001057600080fd5b506102eb806100206000396000f3fe6080604052600436106100345760003560e01c80634672b08314610039578063559ebe571461012f578063773c9d9c1461012f575b600080fd5b61011b600480360361010081101561005057600080fd5b6001600160a01b038235169167ffffffffffffffff60208201358116926040830135909116919081019060808101606082013564010000000081111561009557600080fd5b8201836020820111156100a757600080fd5b803590602001918460018302840111640100000000831117156100c957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060408101359060600135610298565b604080519115158252519081900360200190f35b61011b600480360361012081101561014657600080fd5b6001600160a01b038235169167ffffffffffffffff60208201358116926040830135909116919081019060808101606082013564010000000081111561018b57600080fd5b82018360208201111561019d57600080fd5b803590602001918460018302840111640100000000831117156101bf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561021257600080fd5b82018360208201111561022457600080fd5b8035906020019184600183028401116401000000008311171561024657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604081013590606001356102a6565b600098975050505050505050565b6000999850505050505050505056fea2646970667358221220a1b6d5c988662db6c7b6cb04511fec7e38e5fa3d89e3cbff1ba7f1bf5228d0d964736f6c634300060c0033"

// DeployISwapper deploys a new Ethereum contract, binding an instance of ISwapper to it.
func DeployISwapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ISwapper, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ISwapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ISwapper{ISwapperCaller: ISwapperCaller{contract: contract}, ISwapperTransactor: ISwapperTransactor{contract: contract}, ISwapperFilterer: ISwapperFilterer{contract: contract}}, nil
}

// ISwapper is an auto generated Go binding around an Ethereum contract.
type ISwapper struct {
	ISwapperCaller     // Read-only binding to the contract
	ISwapperTransactor // Write-only binding to the contract
	ISwapperFilterer   // Log filterer for contract events
}

// ISwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapperSession struct {
	Contract     *ISwapper         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapperCallerSession struct {
	Contract *ISwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ISwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapperTransactorSession struct {
	Contract     *ISwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapperRaw struct {
	Contract *ISwapper // Generic contract binding to access the raw methods on
}

// ISwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapperCallerRaw struct {
	Contract *ISwapperCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapperTransactorRaw struct {
	Contract *ISwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapper creates a new instance of ISwapper, bound to a specific deployed contract.
func NewISwapper(address common.Address, backend bind.ContractBackend) (*ISwapper, error) {
	contract, err := bindISwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapper{ISwapperCaller: ISwapperCaller{contract: contract}, ISwapperTransactor: ISwapperTransactor{contract: contract}, ISwapperFilterer: ISwapperFilterer{contract: contract}}, nil
}

// NewISwapperCaller creates a new read-only instance of ISwapper, bound to a specific deployed contract.
func NewISwapperCaller(address common.Address, caller bind.ContractCaller) (*ISwapperCaller, error) {
	contract, err := bindISwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapperCaller{contract: contract}, nil
}

// NewISwapperTransactor creates a new write-only instance of ISwapper, bound to a specific deployed contract.
func NewISwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapperTransactor, error) {
	contract, err := bindISwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapperTransactor{contract: contract}, nil
}

// NewISwapperFilterer creates a new log filterer instance of ISwapper, bound to a specific deployed contract.
func NewISwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapperFilterer, error) {
	contract, err := bindISwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapperFilterer{contract: contract}, nil
}

// bindISwapper binds a generic wrapper to an already deployed contract.
func bindISwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapper *ISwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapper.Contract.ISwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapper *ISwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapper.Contract.ISwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapper *ISwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapper.Contract.ISwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapper *ISwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapper *ISwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapper *ISwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapper.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4672b083.
//
// Solidity: function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperTransactor) AddLiquidity(opts *bind.TransactOpts, fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.contract.Transact(opts, "add_liquidity", fromAssetHash, toPoolId, toChainId, toAddress, amount, minOutAmount, fee, id)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4672b083.
//
// Solidity: function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperSession) AddLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.Contract.AddLiquidity(&_ISwapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAddress, amount, minOutAmount, fee, id)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4672b083.
//
// Solidity: function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperTransactorSession) AddLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.Contract.AddLiquidity(&_ISwapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAddress, amount, minOutAmount, fee, id)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x773c9d9c.
//
// Solidity: function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperTransactor) RemoveLiquidity(opts *bind.TransactOpts, fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.contract.Transact(opts, "remove_liquidity", fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x773c9d9c.
//
// Solidity: function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperSession) RemoveLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.Contract.RemoveLiquidity(&_ISwapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x773c9d9c.
//
// Solidity: function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperTransactorSession) RemoveLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.Contract.RemoveLiquidity(&_ISwapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// Swap is a paid mutator transaction binding the contract method 0x559ebe57.
//
// Solidity: function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperTransactor) Swap(opts *bind.TransactOpts, fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.contract.Transact(opts, "swap", fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// Swap is a paid mutator transaction binding the contract method 0x559ebe57.
//
// Solidity: function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperSession) Swap(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.Contract.Swap(&_ISwapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// Swap is a paid mutator transaction binding the contract method 0x559ebe57.
//
// Solidity: function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_ISwapper *ISwapperTransactorSession) Swap(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _ISwapper.Contract.Swap(&_ISwapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

