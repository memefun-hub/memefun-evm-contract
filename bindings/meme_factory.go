// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meme

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

// MemeFactoryMetaData contains all meta data concerning the MemeFactory contract.
var MemeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"TokenBuy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"icon\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"discord\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"github\",\"type\":\"string\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"K\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dy\",\"type\":\"uint256\"}],\"name\":\"calculateFundsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateTokensReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icon\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discord\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"github\",\"type\":\"string\"}],\"name\":\"createMeme\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingGoal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingGoal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_createFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_createFee\",\"type\":\"uint256\"}],\"name\":\"setCreateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"name\":\"setFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFundingGoal\",\"type\":\"uint256\"}],\"name\":\"setFundingGoal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"enumMemeFactory.TokenState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MemeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MemeFactoryMetaData.ABI instead.
var MemeFactoryABI = MemeFactoryMetaData.ABI

// MemeFactory is an auto generated Go binding around an Ethereum contract.
type MemeFactory struct {
	MemeFactoryCaller     // Read-only binding to the contract
	MemeFactoryTransactor // Write-only binding to the contract
	MemeFactoryFilterer   // Log filterer for contract events
}

// MemeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemeFactorySession struct {
	Contract     *MemeFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemeFactoryCallerSession struct {
	Contract *MemeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MemeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemeFactoryTransactorSession struct {
	Contract     *MemeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MemeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemeFactoryRaw struct {
	Contract *MemeFactory // Generic contract binding to access the raw methods on
}

// MemeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemeFactoryCallerRaw struct {
	Contract *MemeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemeFactoryTransactorRaw struct {
	Contract *MemeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemeFactory creates a new instance of MemeFactory, bound to a specific deployed contract.
func NewMemeFactory(address common.Address, backend bind.ContractBackend) (*MemeFactory, error) {
	contract, err := bindMemeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemeFactory{MemeFactoryCaller: MemeFactoryCaller{contract: contract}, MemeFactoryTransactor: MemeFactoryTransactor{contract: contract}, MemeFactoryFilterer: MemeFactoryFilterer{contract: contract}}, nil
}

// NewMemeFactoryCaller creates a new read-only instance of MemeFactory, bound to a specific deployed contract.
func NewMemeFactoryCaller(address common.Address, caller bind.ContractCaller) (*MemeFactoryCaller, error) {
	contract, err := bindMemeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryCaller{contract: contract}, nil
}

// NewMemeFactoryTransactor creates a new write-only instance of MemeFactory, bound to a specific deployed contract.
func NewMemeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemeFactoryTransactor, error) {
	contract, err := bindMemeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryTransactor{contract: contract}, nil
}

// NewMemeFactoryFilterer creates a new log filterer instance of MemeFactory, bound to a specific deployed contract.
func NewMemeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemeFactoryFilterer, error) {
	contract, err := bindMemeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryFilterer{contract: contract}, nil
}

// bindMemeFactory binds a generic wrapper to an already deployed contract.
func bindMemeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemeFactory *MemeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemeFactory.Contract.MemeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemeFactory *MemeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemeFactory.Contract.MemeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemeFactory *MemeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemeFactory.Contract.MemeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemeFactory *MemeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemeFactory *MemeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemeFactory *MemeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemeFactory.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_MemeFactory *MemeFactorySession) FEEDENOMINATOR() (*big.Int, error) {
	return _MemeFactory.Contract.FEEDENOMINATOR(&_MemeFactory.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _MemeFactory.Contract.FEEDENOMINATOR(&_MemeFactory.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xa932492f.
//
// Solidity: function K() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "K")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xa932492f.
//
// Solidity: function K() view returns(uint256)
func (_MemeFactory *MemeFactorySession) K() (*big.Int, error) {
	return _MemeFactory.Contract.K(&_MemeFactory.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xa932492f.
//
// Solidity: function K() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) K() (*big.Int, error) {
	return _MemeFactory.Contract.K(&_MemeFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MemeFactory *MemeFactoryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MemeFactory *MemeFactorySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MemeFactory.Contract.UPGRADEINTERFACEVERSION(&_MemeFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MemeFactory *MemeFactoryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MemeFactory.Contract.UPGRADEINTERFACEVERSION(&_MemeFactory.CallOpts)
}

// CalculateFundsReceived is a free data retrieval call binding the contract method 0x7accf009.
//
// Solidity: function calculateFundsReceived(uint256 K, uint256 x, uint256 y, uint256 dy) pure returns(uint256)
func (_MemeFactory *MemeFactoryCaller) CalculateFundsReceived(opts *bind.CallOpts, K *big.Int, x *big.Int, y *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "calculateFundsReceived", K, x, y, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFundsReceived is a free data retrieval call binding the contract method 0x7accf009.
//
// Solidity: function calculateFundsReceived(uint256 K, uint256 x, uint256 y, uint256 dy) pure returns(uint256)
func (_MemeFactory *MemeFactorySession) CalculateFundsReceived(K *big.Int, x *big.Int, y *big.Int, dy *big.Int) (*big.Int, error) {
	return _MemeFactory.Contract.CalculateFundsReceived(&_MemeFactory.CallOpts, K, x, y, dy)
}

// CalculateFundsReceived is a free data retrieval call binding the contract method 0x7accf009.
//
// Solidity: function calculateFundsReceived(uint256 K, uint256 x, uint256 y, uint256 dy) pure returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) CalculateFundsReceived(K *big.Int, x *big.Int, y *big.Int, dy *big.Int) (*big.Int, error) {
	return _MemeFactory.Contract.CalculateFundsReceived(&_MemeFactory.CallOpts, K, x, y, dy)
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0xf290f5c8.
//
// Solidity: function calculateTokensReceived(uint256 K, uint256 x, uint256 y, uint256 dx) pure returns(uint256)
func (_MemeFactory *MemeFactoryCaller) CalculateTokensReceived(opts *bind.CallOpts, K *big.Int, x *big.Int, y *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "calculateTokensReceived", K, x, y, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0xf290f5c8.
//
// Solidity: function calculateTokensReceived(uint256 K, uint256 x, uint256 y, uint256 dx) pure returns(uint256)
func (_MemeFactory *MemeFactorySession) CalculateTokensReceived(K *big.Int, x *big.Int, y *big.Int, dx *big.Int) (*big.Int, error) {
	return _MemeFactory.Contract.CalculateTokensReceived(&_MemeFactory.CallOpts, K, x, y, dx)
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0xf290f5c8.
//
// Solidity: function calculateTokensReceived(uint256 K, uint256 x, uint256 y, uint256 dx) pure returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) CalculateTokensReceived(K *big.Int, x *big.Int, y *big.Int, dx *big.Int) (*big.Int, error) {
	return _MemeFactory.Contract.CalculateTokensReceived(&_MemeFactory.CallOpts, K, x, y, dx)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address ) view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) Collateral(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "collateral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address ) view returns(uint256)
func (_MemeFactory *MemeFactorySession) Collateral(arg0 common.Address) (*big.Int, error) {
	return _MemeFactory.Contract.Collateral(&_MemeFactory.CallOpts, arg0)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address ) view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) Collateral(arg0 common.Address) (*big.Int, error) {
	return _MemeFactory.Contract.Collateral(&_MemeFactory.CallOpts, arg0)
}

// CreateFee is a free data retrieval call binding the contract method 0xb2db919b.
//
// Solidity: function createFee() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) CreateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "createFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreateFee is a free data retrieval call binding the contract method 0xb2db919b.
//
// Solidity: function createFee() view returns(uint256)
func (_MemeFactory *MemeFactorySession) CreateFee() (*big.Int, error) {
	return _MemeFactory.Contract.CreateFee(&_MemeFactory.CallOpts)
}

// CreateFee is a free data retrieval call binding the contract method 0xb2db919b.
//
// Solidity: function createFee() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) CreateFee() (*big.Int, error) {
	return _MemeFactory.Contract.CreateFee(&_MemeFactory.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_MemeFactory *MemeFactorySession) Fee() (*big.Int, error) {
	return _MemeFactory.Contract.Fee(&_MemeFactory.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) Fee() (*big.Int, error) {
	return _MemeFactory.Contract.Fee(&_MemeFactory.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) FeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "feePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_MemeFactory *MemeFactorySession) FeePercent() (*big.Int, error) {
	return _MemeFactory.Contract.FeePercent(&_MemeFactory.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) FeePercent() (*big.Int, error) {
	return _MemeFactory.Contract.FeePercent(&_MemeFactory.CallOpts)
}

// FundingGoal is a free data retrieval call binding the contract method 0x7a3a0e84.
//
// Solidity: function fundingGoal() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) FundingGoal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "fundingGoal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingGoal is a free data retrieval call binding the contract method 0x7a3a0e84.
//
// Solidity: function fundingGoal() view returns(uint256)
func (_MemeFactory *MemeFactorySession) FundingGoal() (*big.Int, error) {
	return _MemeFactory.Contract.FundingGoal(&_MemeFactory.CallOpts)
}

// FundingGoal is a free data retrieval call binding the contract method 0x7a3a0e84.
//
// Solidity: function fundingGoal() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) FundingGoal() (*big.Int, error) {
	return _MemeFactory.Contract.FundingGoal(&_MemeFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MemeFactory *MemeFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MemeFactory *MemeFactorySession) Owner() (common.Address, error) {
	return _MemeFactory.Contract.Owner(&_MemeFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MemeFactory *MemeFactoryCallerSession) Owner() (common.Address, error) {
	return _MemeFactory.Contract.Owner(&_MemeFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MemeFactory *MemeFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MemeFactory *MemeFactorySession) ProxiableUUID() ([32]byte, error) {
	return _MemeFactory.Contract.ProxiableUUID(&_MemeFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MemeFactory *MemeFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MemeFactory.Contract.ProxiableUUID(&_MemeFactory.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x953909f0.
//
// Solidity: function tokenX(address ) view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) TokenX(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "tokenX", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenX is a free data retrieval call binding the contract method 0x953909f0.
//
// Solidity: function tokenX(address ) view returns(uint256)
func (_MemeFactory *MemeFactorySession) TokenX(arg0 common.Address) (*big.Int, error) {
	return _MemeFactory.Contract.TokenX(&_MemeFactory.CallOpts, arg0)
}

// TokenX is a free data retrieval call binding the contract method 0x953909f0.
//
// Solidity: function tokenX(address ) view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) TokenX(arg0 common.Address) (*big.Int, error) {
	return _MemeFactory.Contract.TokenX(&_MemeFactory.CallOpts, arg0)
}

// TokenY is a free data retrieval call binding the contract method 0xc2013721.
//
// Solidity: function tokenY(address ) view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) TokenY(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "tokenY", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenY is a free data retrieval call binding the contract method 0xc2013721.
//
// Solidity: function tokenY(address ) view returns(uint256)
func (_MemeFactory *MemeFactorySession) TokenY(arg0 common.Address) (*big.Int, error) {
	return _MemeFactory.Contract.TokenY(&_MemeFactory.CallOpts, arg0)
}

// TokenY is a free data retrieval call binding the contract method 0xc2013721.
//
// Solidity: function tokenY(address ) view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) TokenY(arg0 common.Address) (*big.Int, error) {
	return _MemeFactory.Contract.TokenY(&_MemeFactory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint8)
func (_MemeFactory *MemeFactoryCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint8)
func (_MemeFactory *MemeFactorySession) Tokens(arg0 common.Address) (uint8, error) {
	return _MemeFactory.Contract.Tokens(&_MemeFactory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint8)
func (_MemeFactory *MemeFactoryCallerSession) Tokens(arg0 common.Address) (uint8, error) {
	return _MemeFactory.Contract.Tokens(&_MemeFactory.CallOpts, arg0)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_MemeFactory *MemeFactorySession) X() (*big.Int, error) {
	return _MemeFactory.Contract.X(&_MemeFactory.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) X() (*big.Int, error) {
	return _MemeFactory.Contract.X(&_MemeFactory.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_MemeFactory *MemeFactoryCaller) Y(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MemeFactory.contract.Call(opts, &out, "y")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_MemeFactory *MemeFactorySession) Y() (*big.Int, error) {
	return _MemeFactory.Contract.Y(&_MemeFactory.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_MemeFactory *MemeFactoryCallerSession) Y() (*big.Int, error) {
	return _MemeFactory.Contract.Y(&_MemeFactory.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address tokenAddress) payable returns()
func (_MemeFactory *MemeFactoryTransactor) Buy(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "buy", tokenAddress)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address tokenAddress) payable returns()
func (_MemeFactory *MemeFactorySession) Buy(tokenAddress common.Address) (*types.Transaction, error) {
	return _MemeFactory.Contract.Buy(&_MemeFactory.TransactOpts, tokenAddress)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address tokenAddress) payable returns()
func (_MemeFactory *MemeFactoryTransactorSession) Buy(tokenAddress common.Address) (*types.Transaction, error) {
	return _MemeFactory.Contract.Buy(&_MemeFactory.TransactOpts, tokenAddress)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_MemeFactory *MemeFactoryTransactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_MemeFactory *MemeFactorySession) ClaimFee() (*types.Transaction, error) {
	return _MemeFactory.Contract.ClaimFee(&_MemeFactory.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_MemeFactory *MemeFactoryTransactorSession) ClaimFee() (*types.Transaction, error) {
	return _MemeFactory.Contract.ClaimFee(&_MemeFactory.TransactOpts)
}

// CreateMeme is a paid mutator transaction binding the contract method 0xeffd76b7.
//
// Solidity: function createMeme(string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github) payable returns(address)
func (_MemeFactory *MemeFactoryTransactor) CreateMeme(opts *bind.TransactOpts, name string, symbol string, description string, icon string, website string, twitter string, telegram string, discord string, github string) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "createMeme", name, symbol, description, icon, website, twitter, telegram, discord, github)
}

// CreateMeme is a paid mutator transaction binding the contract method 0xeffd76b7.
//
// Solidity: function createMeme(string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github) payable returns(address)
func (_MemeFactory *MemeFactorySession) CreateMeme(name string, symbol string, description string, icon string, website string, twitter string, telegram string, discord string, github string) (*types.Transaction, error) {
	return _MemeFactory.Contract.CreateMeme(&_MemeFactory.TransactOpts, name, symbol, description, icon, website, twitter, telegram, discord, github)
}

// CreateMeme is a paid mutator transaction binding the contract method 0xeffd76b7.
//
// Solidity: function createMeme(string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github) payable returns(address)
func (_MemeFactory *MemeFactoryTransactorSession) CreateMeme(name string, symbol string, description string, icon string, website string, twitter string, telegram string, discord string, github string) (*types.Transaction, error) {
	return _MemeFactory.Contract.CreateMeme(&_MemeFactory.TransactOpts, name, symbol, description, icon, website, twitter, telegram, discord, github)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _x, uint256 _y, uint256 _fundingGoal, uint256 _createFee, uint256 _feePercent) returns()
func (_MemeFactory *MemeFactoryTransactor) Initialize(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _fundingGoal *big.Int, _createFee *big.Int, _feePercent *big.Int) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "initialize", _x, _y, _fundingGoal, _createFee, _feePercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _x, uint256 _y, uint256 _fundingGoal, uint256 _createFee, uint256 _feePercent) returns()
func (_MemeFactory *MemeFactorySession) Initialize(_x *big.Int, _y *big.Int, _fundingGoal *big.Int, _createFee *big.Int, _feePercent *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.Initialize(&_MemeFactory.TransactOpts, _x, _y, _fundingGoal, _createFee, _feePercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _x, uint256 _y, uint256 _fundingGoal, uint256 _createFee, uint256 _feePercent) returns()
func (_MemeFactory *MemeFactoryTransactorSession) Initialize(_x *big.Int, _y *big.Int, _fundingGoal *big.Int, _createFee *big.Int, _feePercent *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.Initialize(&_MemeFactory.TransactOpts, _x, _y, _fundingGoal, _createFee, _feePercent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MemeFactory *MemeFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MemeFactory *MemeFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _MemeFactory.Contract.RenounceOwnership(&_MemeFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MemeFactory *MemeFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MemeFactory.Contract.RenounceOwnership(&_MemeFactory.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tokenAddress, uint256 amount) returns()
func (_MemeFactory *MemeFactoryTransactor) Sell(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "sell", tokenAddress, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tokenAddress, uint256 amount) returns()
func (_MemeFactory *MemeFactorySession) Sell(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.Sell(&_MemeFactory.TransactOpts, tokenAddress, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tokenAddress, uint256 amount) returns()
func (_MemeFactory *MemeFactoryTransactorSession) Sell(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.Sell(&_MemeFactory.TransactOpts, tokenAddress, amount)
}

// SetCreateFee is a paid mutator transaction binding the contract method 0x62914849.
//
// Solidity: function setCreateFee(uint256 _createFee) returns()
func (_MemeFactory *MemeFactoryTransactor) SetCreateFee(opts *bind.TransactOpts, _createFee *big.Int) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "setCreateFee", _createFee)
}

// SetCreateFee is a paid mutator transaction binding the contract method 0x62914849.
//
// Solidity: function setCreateFee(uint256 _createFee) returns()
func (_MemeFactory *MemeFactorySession) SetCreateFee(_createFee *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.SetCreateFee(&_MemeFactory.TransactOpts, _createFee)
}

// SetCreateFee is a paid mutator transaction binding the contract method 0x62914849.
//
// Solidity: function setCreateFee(uint256 _createFee) returns()
func (_MemeFactory *MemeFactoryTransactorSession) SetCreateFee(_createFee *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.SetCreateFee(&_MemeFactory.TransactOpts, _createFee)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x7ce3489b.
//
// Solidity: function setFeePercent(uint256 _feePercent) returns()
func (_MemeFactory *MemeFactoryTransactor) SetFeePercent(opts *bind.TransactOpts, _feePercent *big.Int) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "setFeePercent", _feePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x7ce3489b.
//
// Solidity: function setFeePercent(uint256 _feePercent) returns()
func (_MemeFactory *MemeFactorySession) SetFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.SetFeePercent(&_MemeFactory.TransactOpts, _feePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x7ce3489b.
//
// Solidity: function setFeePercent(uint256 _feePercent) returns()
func (_MemeFactory *MemeFactoryTransactorSession) SetFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.SetFeePercent(&_MemeFactory.TransactOpts, _feePercent)
}

// SetFundingGoal is a paid mutator transaction binding the contract method 0x576eac66.
//
// Solidity: function setFundingGoal(uint256 newFundingGoal) returns()
func (_MemeFactory *MemeFactoryTransactor) SetFundingGoal(opts *bind.TransactOpts, newFundingGoal *big.Int) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "setFundingGoal", newFundingGoal)
}

// SetFundingGoal is a paid mutator transaction binding the contract method 0x576eac66.
//
// Solidity: function setFundingGoal(uint256 newFundingGoal) returns()
func (_MemeFactory *MemeFactorySession) SetFundingGoal(newFundingGoal *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.SetFundingGoal(&_MemeFactory.TransactOpts, newFundingGoal)
}

// SetFundingGoal is a paid mutator transaction binding the contract method 0x576eac66.
//
// Solidity: function setFundingGoal(uint256 newFundingGoal) returns()
func (_MemeFactory *MemeFactoryTransactorSession) SetFundingGoal(newFundingGoal *big.Int) (*types.Transaction, error) {
	return _MemeFactory.Contract.SetFundingGoal(&_MemeFactory.TransactOpts, newFundingGoal)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MemeFactory *MemeFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MemeFactory *MemeFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MemeFactory.Contract.TransferOwnership(&_MemeFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MemeFactory *MemeFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MemeFactory.Contract.TransferOwnership(&_MemeFactory.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MemeFactory *MemeFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MemeFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MemeFactory *MemeFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MemeFactory.Contract.UpgradeToAndCall(&_MemeFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MemeFactory *MemeFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MemeFactory.Contract.UpgradeToAndCall(&_MemeFactory.TransactOpts, newImplementation, data)
}

// MemeFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MemeFactory contract.
type MemeFactoryInitializedIterator struct {
	Event *MemeFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *MemeFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeFactoryInitialized)
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
		it.Event = new(MemeFactoryInitialized)
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
func (it *MemeFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeFactoryInitialized represents a Initialized event raised by the MemeFactory contract.
type MemeFactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MemeFactory *MemeFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*MemeFactoryInitializedIterator, error) {

	logs, sub, err := _MemeFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MemeFactoryInitializedIterator{contract: _MemeFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MemeFactory *MemeFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MemeFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _MemeFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeFactoryInitialized)
				if err := _MemeFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MemeFactory *MemeFactoryFilterer) ParseInitialized(log types.Log) (*MemeFactoryInitialized, error) {
	event := new(MemeFactoryInitialized)
	if err := _MemeFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MemeFactory contract.
type MemeFactoryOwnershipTransferredIterator struct {
	Event *MemeFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MemeFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeFactoryOwnershipTransferred)
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
		it.Event = new(MemeFactoryOwnershipTransferred)
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
func (it *MemeFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the MemeFactory contract.
type MemeFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MemeFactory *MemeFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MemeFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MemeFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryOwnershipTransferredIterator{contract: _MemeFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MemeFactory *MemeFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MemeFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MemeFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeFactoryOwnershipTransferred)
				if err := _MemeFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MemeFactory *MemeFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*MemeFactoryOwnershipTransferred, error) {
	event := new(MemeFactoryOwnershipTransferred)
	if err := _MemeFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeFactoryTokenBuyIterator is returned from FilterTokenBuy and is used to iterate over the raw logs and unpacked data for TokenBuy events raised by the MemeFactory contract.
type MemeFactoryTokenBuyIterator struct {
	Event *MemeFactoryTokenBuy // Event containing the contract specifics and raw log

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
func (it *MemeFactoryTokenBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeFactoryTokenBuy)
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
		it.Event = new(MemeFactoryTokenBuy)
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
func (it *MemeFactoryTokenBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeFactoryTokenBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeFactoryTokenBuy represents a TokenBuy event raised by the MemeFactory contract.
type MemeFactoryTokenBuy struct {
	Token       common.Address
	Timestamp   *big.Int
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	FeeAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenBuy is a free log retrieval operation binding the contract event 0x6c6bf9b07a057cb0766b2b0aefeb3ed76d844744cc34e081c60bb230ef57f65d.
//
// Solidity: event TokenBuy(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_MemeFactory *MemeFactoryFilterer) FilterTokenBuy(opts *bind.FilterOpts, token []common.Address) (*MemeFactoryTokenBuyIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MemeFactory.contract.FilterLogs(opts, "TokenBuy", tokenRule)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryTokenBuyIterator{contract: _MemeFactory.contract, event: "TokenBuy", logs: logs, sub: sub}, nil
}

// WatchTokenBuy is a free log subscription operation binding the contract event 0x6c6bf9b07a057cb0766b2b0aefeb3ed76d844744cc34e081c60bb230ef57f65d.
//
// Solidity: event TokenBuy(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_MemeFactory *MemeFactoryFilterer) WatchTokenBuy(opts *bind.WatchOpts, sink chan<- *MemeFactoryTokenBuy, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MemeFactory.contract.WatchLogs(opts, "TokenBuy", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeFactoryTokenBuy)
				if err := _MemeFactory.contract.UnpackLog(event, "TokenBuy", log); err != nil {
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

// ParseTokenBuy is a log parse operation binding the contract event 0x6c6bf9b07a057cb0766b2b0aefeb3ed76d844744cc34e081c60bb230ef57f65d.
//
// Solidity: event TokenBuy(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_MemeFactory *MemeFactoryFilterer) ParseTokenBuy(log types.Log) (*MemeFactoryTokenBuy, error) {
	event := new(MemeFactoryTokenBuy)
	if err := _MemeFactory.contract.UnpackLog(event, "TokenBuy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeFactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the MemeFactory contract.
type MemeFactoryTokenCreatedIterator struct {
	Event *MemeFactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *MemeFactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeFactoryTokenCreated)
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
		it.Event = new(MemeFactoryTokenCreated)
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
func (it *MemeFactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeFactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeFactoryTokenCreated represents a TokenCreated event raised by the MemeFactory contract.
type MemeFactoryTokenCreated struct {
	Token       common.Address
	Creator     common.Address
	Timestamp   *big.Int
	Name        string
	Symbol      string
	Description string
	Icon        string
	Website     string
	Twitter     string
	Telegram    string
	Discord     string
	Github      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0x1a2bd37fcc163f5d288090a0f9d97dc7dcf61b340620c5da1049f9cc413a9708.
//
// Solidity: event TokenCreated(address indexed token, address indexed creator, uint256 timestamp, string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github)
func (_MemeFactory *MemeFactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts, token []common.Address, creator []common.Address) (*MemeFactoryTokenCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _MemeFactory.contract.FilterLogs(opts, "TokenCreated", tokenRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryTokenCreatedIterator{contract: _MemeFactory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0x1a2bd37fcc163f5d288090a0f9d97dc7dcf61b340620c5da1049f9cc413a9708.
//
// Solidity: event TokenCreated(address indexed token, address indexed creator, uint256 timestamp, string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github)
func (_MemeFactory *MemeFactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *MemeFactoryTokenCreated, token []common.Address, creator []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _MemeFactory.contract.WatchLogs(opts, "TokenCreated", tokenRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeFactoryTokenCreated)
				if err := _MemeFactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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

// ParseTokenCreated is a log parse operation binding the contract event 0x1a2bd37fcc163f5d288090a0f9d97dc7dcf61b340620c5da1049f9cc413a9708.
//
// Solidity: event TokenCreated(address indexed token, address indexed creator, uint256 timestamp, string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github)
func (_MemeFactory *MemeFactoryFilterer) ParseTokenCreated(log types.Log) (*MemeFactoryTokenCreated, error) {
	event := new(MemeFactoryTokenCreated)
	if err := _MemeFactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeFactoryTokenSellIterator is returned from FilterTokenSell and is used to iterate over the raw logs and unpacked data for TokenSell events raised by the MemeFactory contract.
type MemeFactoryTokenSellIterator struct {
	Event *MemeFactoryTokenSell // Event containing the contract specifics and raw log

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
func (it *MemeFactoryTokenSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeFactoryTokenSell)
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
		it.Event = new(MemeFactoryTokenSell)
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
func (it *MemeFactoryTokenSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeFactoryTokenSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeFactoryTokenSell represents a TokenSell event raised by the MemeFactory contract.
type MemeFactoryTokenSell struct {
	Token       common.Address
	Timestamp   *big.Int
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	FeeAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenSell is a free log retrieval operation binding the contract event 0xb40ebf7368df31a675c3b3a523e81f6b28b5cf6e4b32586ed0bc6a86052d8991.
//
// Solidity: event TokenSell(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_MemeFactory *MemeFactoryFilterer) FilterTokenSell(opts *bind.FilterOpts, token []common.Address) (*MemeFactoryTokenSellIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MemeFactory.contract.FilterLogs(opts, "TokenSell", tokenRule)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryTokenSellIterator{contract: _MemeFactory.contract, event: "TokenSell", logs: logs, sub: sub}, nil
}

// WatchTokenSell is a free log subscription operation binding the contract event 0xb40ebf7368df31a675c3b3a523e81f6b28b5cf6e4b32586ed0bc6a86052d8991.
//
// Solidity: event TokenSell(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_MemeFactory *MemeFactoryFilterer) WatchTokenSell(opts *bind.WatchOpts, sink chan<- *MemeFactoryTokenSell, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MemeFactory.contract.WatchLogs(opts, "TokenSell", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeFactoryTokenSell)
				if err := _MemeFactory.contract.UnpackLog(event, "TokenSell", log); err != nil {
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

// ParseTokenSell is a log parse operation binding the contract event 0xb40ebf7368df31a675c3b3a523e81f6b28b5cf6e4b32586ed0bc6a86052d8991.
//
// Solidity: event TokenSell(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_MemeFactory *MemeFactoryFilterer) ParseTokenSell(log types.Log) (*MemeFactoryTokenSell, error) {
	event := new(MemeFactoryTokenSell)
	if err := _MemeFactory.contract.UnpackLog(event, "TokenSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MemeFactory contract.
type MemeFactoryUpgradedIterator struct {
	Event *MemeFactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *MemeFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeFactoryUpgraded)
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
		it.Event = new(MemeFactoryUpgraded)
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
func (it *MemeFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeFactoryUpgraded represents a Upgraded event raised by the MemeFactory contract.
type MemeFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MemeFactory *MemeFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MemeFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MemeFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MemeFactoryUpgradedIterator{contract: _MemeFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MemeFactory *MemeFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MemeFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MemeFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeFactoryUpgraded)
				if err := _MemeFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MemeFactory *MemeFactoryFilterer) ParseUpgraded(log types.Log) (*MemeFactoryUpgraded, error) {
	event := new(MemeFactoryUpgraded)
	if err := _MemeFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
