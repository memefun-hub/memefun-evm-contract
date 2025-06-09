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

// MemeMetaData contains all meta data concerning the Meme contract.
var MemeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"TokenBuy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"icon\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"discord\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"github\",\"type\":\"string\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"K\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dy\",\"type\":\"uint256\"}],\"name\":\"calculateFundsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateTokensReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icon\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegram\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discord\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"github\",\"type\":\"string\"}],\"name\":\"createMeme\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingGoal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingGoal\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_createFee\",\"type\":\"uint256\"}],\"name\":\"setCreateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"name\":\"setFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFundingGoal\",\"type\":\"uint256\"}],\"name\":\"setFundingGoal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"enumMemeFactory.TokenState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MemeABI is the input ABI used to generate the binding from.
// Deprecated: Use MemeMetaData.ABI instead.
var MemeABI = MemeMetaData.ABI

// Meme is an auto generated Go binding around an Ethereum contract.
type Meme struct {
	MemeCaller     // Read-only binding to the contract
	MemeTransactor // Write-only binding to the contract
	MemeFilterer   // Log filterer for contract events
}

// MemeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemeSession struct {
	Contract     *Meme             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemeCallerSession struct {
	Contract *MemeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemeTransactorSession struct {
	Contract     *MemeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemeRaw struct {
	Contract *Meme // Generic contract binding to access the raw methods on
}

// MemeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemeCallerRaw struct {
	Contract *MemeCaller // Generic read-only contract binding to access the raw methods on
}

// MemeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemeTransactorRaw struct {
	Contract *MemeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeme creates a new instance of Meme, bound to a specific deployed contract.
func NewMeme(address common.Address, backend bind.ContractBackend) (*Meme, error) {
	contract, err := bindMeme(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Meme{MemeCaller: MemeCaller{contract: contract}, MemeTransactor: MemeTransactor{contract: contract}, MemeFilterer: MemeFilterer{contract: contract}}, nil
}

// NewMemeCaller creates a new read-only instance of Meme, bound to a specific deployed contract.
func NewMemeCaller(address common.Address, caller bind.ContractCaller) (*MemeCaller, error) {
	contract, err := bindMeme(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemeCaller{contract: contract}, nil
}

// NewMemeTransactor creates a new write-only instance of Meme, bound to a specific deployed contract.
func NewMemeTransactor(address common.Address, transactor bind.ContractTransactor) (*MemeTransactor, error) {
	contract, err := bindMeme(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemeTransactor{contract: contract}, nil
}

// NewMemeFilterer creates a new log filterer instance of Meme, bound to a specific deployed contract.
func NewMemeFilterer(address common.Address, filterer bind.ContractFilterer) (*MemeFilterer, error) {
	contract, err := bindMeme(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemeFilterer{contract: contract}, nil
}

// bindMeme binds a generic wrapper to an already deployed contract.
func bindMeme(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meme *MemeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meme.Contract.MemeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meme *MemeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meme.Contract.MemeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meme *MemeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meme.Contract.MemeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meme *MemeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meme.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meme *MemeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meme.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meme *MemeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meme.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Meme *MemeCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Meme *MemeSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Meme.Contract.FEEDENOMINATOR(&_Meme.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Meme *MemeCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Meme.Contract.FEEDENOMINATOR(&_Meme.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xa932492f.
//
// Solidity: function K() view returns(uint256)
func (_Meme *MemeCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "K")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xa932492f.
//
// Solidity: function K() view returns(uint256)
func (_Meme *MemeSession) K() (*big.Int, error) {
	return _Meme.Contract.K(&_Meme.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xa932492f.
//
// Solidity: function K() view returns(uint256)
func (_Meme *MemeCallerSession) K() (*big.Int, error) {
	return _Meme.Contract.K(&_Meme.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Meme *MemeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Meme *MemeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Meme.Contract.UPGRADEINTERFACEVERSION(&_Meme.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Meme *MemeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Meme.Contract.UPGRADEINTERFACEVERSION(&_Meme.CallOpts)
}

// CalculateFundsReceived is a free data retrieval call binding the contract method 0x7accf009.
//
// Solidity: function calculateFundsReceived(uint256 K, uint256 x, uint256 y, uint256 dy) pure returns(uint256)
func (_Meme *MemeCaller) CalculateFundsReceived(opts *bind.CallOpts, K *big.Int, x *big.Int, y *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "calculateFundsReceived", K, x, y, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFundsReceived is a free data retrieval call binding the contract method 0x7accf009.
//
// Solidity: function calculateFundsReceived(uint256 K, uint256 x, uint256 y, uint256 dy) pure returns(uint256)
func (_Meme *MemeSession) CalculateFundsReceived(K *big.Int, x *big.Int, y *big.Int, dy *big.Int) (*big.Int, error) {
	return _Meme.Contract.CalculateFundsReceived(&_Meme.CallOpts, K, x, y, dy)
}

// CalculateFundsReceived is a free data retrieval call binding the contract method 0x7accf009.
//
// Solidity: function calculateFundsReceived(uint256 K, uint256 x, uint256 y, uint256 dy) pure returns(uint256)
func (_Meme *MemeCallerSession) CalculateFundsReceived(K *big.Int, x *big.Int, y *big.Int, dy *big.Int) (*big.Int, error) {
	return _Meme.Contract.CalculateFundsReceived(&_Meme.CallOpts, K, x, y, dy)
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0xf290f5c8.
//
// Solidity: function calculateTokensReceived(uint256 K, uint256 x, uint256 y, uint256 dx) pure returns(uint256)
func (_Meme *MemeCaller) CalculateTokensReceived(opts *bind.CallOpts, K *big.Int, x *big.Int, y *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "calculateTokensReceived", K, x, y, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0xf290f5c8.
//
// Solidity: function calculateTokensReceived(uint256 K, uint256 x, uint256 y, uint256 dx) pure returns(uint256)
func (_Meme *MemeSession) CalculateTokensReceived(K *big.Int, x *big.Int, y *big.Int, dx *big.Int) (*big.Int, error) {
	return _Meme.Contract.CalculateTokensReceived(&_Meme.CallOpts, K, x, y, dx)
}

// CalculateTokensReceived is a free data retrieval call binding the contract method 0xf290f5c8.
//
// Solidity: function calculateTokensReceived(uint256 K, uint256 x, uint256 y, uint256 dx) pure returns(uint256)
func (_Meme *MemeCallerSession) CalculateTokensReceived(K *big.Int, x *big.Int, y *big.Int, dx *big.Int) (*big.Int, error) {
	return _Meme.Contract.CalculateTokensReceived(&_Meme.CallOpts, K, x, y, dx)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address ) view returns(uint256)
func (_Meme *MemeCaller) Collateral(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "collateral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address ) view returns(uint256)
func (_Meme *MemeSession) Collateral(arg0 common.Address) (*big.Int, error) {
	return _Meme.Contract.Collateral(&_Meme.CallOpts, arg0)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address ) view returns(uint256)
func (_Meme *MemeCallerSession) Collateral(arg0 common.Address) (*big.Int, error) {
	return _Meme.Contract.Collateral(&_Meme.CallOpts, arg0)
}

// CreateFee is a free data retrieval call binding the contract method 0xb2db919b.
//
// Solidity: function createFee() view returns(uint256)
func (_Meme *MemeCaller) CreateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "createFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreateFee is a free data retrieval call binding the contract method 0xb2db919b.
//
// Solidity: function createFee() view returns(uint256)
func (_Meme *MemeSession) CreateFee() (*big.Int, error) {
	return _Meme.Contract.CreateFee(&_Meme.CallOpts)
}

// CreateFee is a free data retrieval call binding the contract method 0xb2db919b.
//
// Solidity: function createFee() view returns(uint256)
func (_Meme *MemeCallerSession) CreateFee() (*big.Int, error) {
	return _Meme.Contract.CreateFee(&_Meme.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Meme *MemeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Meme *MemeSession) Fee() (*big.Int, error) {
	return _Meme.Contract.Fee(&_Meme.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Meme *MemeCallerSession) Fee() (*big.Int, error) {
	return _Meme.Contract.Fee(&_Meme.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_Meme *MemeCaller) FeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "feePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_Meme *MemeSession) FeePercent() (*big.Int, error) {
	return _Meme.Contract.FeePercent(&_Meme.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_Meme *MemeCallerSession) FeePercent() (*big.Int, error) {
	return _Meme.Contract.FeePercent(&_Meme.CallOpts)
}

// FundingGoal is a free data retrieval call binding the contract method 0x7a3a0e84.
//
// Solidity: function fundingGoal() view returns(uint256)
func (_Meme *MemeCaller) FundingGoal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "fundingGoal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingGoal is a free data retrieval call binding the contract method 0x7a3a0e84.
//
// Solidity: function fundingGoal() view returns(uint256)
func (_Meme *MemeSession) FundingGoal() (*big.Int, error) {
	return _Meme.Contract.FundingGoal(&_Meme.CallOpts)
}

// FundingGoal is a free data retrieval call binding the contract method 0x7a3a0e84.
//
// Solidity: function fundingGoal() view returns(uint256)
func (_Meme *MemeCallerSession) FundingGoal() (*big.Int, error) {
	return _Meme.Contract.FundingGoal(&_Meme.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meme *MemeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meme *MemeSession) Owner() (common.Address, error) {
	return _Meme.Contract.Owner(&_Meme.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meme *MemeCallerSession) Owner() (common.Address, error) {
	return _Meme.Contract.Owner(&_Meme.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Meme *MemeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Meme *MemeSession) ProxiableUUID() ([32]byte, error) {
	return _Meme.Contract.ProxiableUUID(&_Meme.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Meme *MemeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Meme.Contract.ProxiableUUID(&_Meme.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint8)
func (_Meme *MemeCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint8)
func (_Meme *MemeSession) Tokens(arg0 common.Address) (uint8, error) {
	return _Meme.Contract.Tokens(&_Meme.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint8)
func (_Meme *MemeCallerSession) Tokens(arg0 common.Address) (uint8, error) {
	return _Meme.Contract.Tokens(&_Meme.CallOpts, arg0)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Meme *MemeCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Meme *MemeSession) X() (*big.Int, error) {
	return _Meme.Contract.X(&_Meme.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Meme *MemeCallerSession) X() (*big.Int, error) {
	return _Meme.Contract.X(&_Meme.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_Meme *MemeCaller) Y(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meme.contract.Call(opts, &out, "y")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_Meme *MemeSession) Y() (*big.Int, error) {
	return _Meme.Contract.Y(&_Meme.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(uint256)
func (_Meme *MemeCallerSession) Y() (*big.Int, error) {
	return _Meme.Contract.Y(&_Meme.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address tokenAddress) payable returns()
func (_Meme *MemeTransactor) Buy(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "buy", tokenAddress)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address tokenAddress) payable returns()
func (_Meme *MemeSession) Buy(tokenAddress common.Address) (*types.Transaction, error) {
	return _Meme.Contract.Buy(&_Meme.TransactOpts, tokenAddress)
}

// Buy is a paid mutator transaction binding the contract method 0xf088d547.
//
// Solidity: function buy(address tokenAddress) payable returns()
func (_Meme *MemeTransactorSession) Buy(tokenAddress common.Address) (*types.Transaction, error) {
	return _Meme.Contract.Buy(&_Meme.TransactOpts, tokenAddress)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_Meme *MemeTransactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_Meme *MemeSession) ClaimFee() (*types.Transaction, error) {
	return _Meme.Contract.ClaimFee(&_Meme.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_Meme *MemeTransactorSession) ClaimFee() (*types.Transaction, error) {
	return _Meme.Contract.ClaimFee(&_Meme.TransactOpts)
}

// CreateMeme is a paid mutator transaction binding the contract method 0xeffd76b7.
//
// Solidity: function createMeme(string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github) payable returns(address)
func (_Meme *MemeTransactor) CreateMeme(opts *bind.TransactOpts, name string, symbol string, description string, icon string, website string, twitter string, telegram string, discord string, github string) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "createMeme", name, symbol, description, icon, website, twitter, telegram, discord, github)
}

// CreateMeme is a paid mutator transaction binding the contract method 0xeffd76b7.
//
// Solidity: function createMeme(string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github) payable returns(address)
func (_Meme *MemeSession) CreateMeme(name string, symbol string, description string, icon string, website string, twitter string, telegram string, discord string, github string) (*types.Transaction, error) {
	return _Meme.Contract.CreateMeme(&_Meme.TransactOpts, name, symbol, description, icon, website, twitter, telegram, discord, github)
}

// CreateMeme is a paid mutator transaction binding the contract method 0xeffd76b7.
//
// Solidity: function createMeme(string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github) payable returns(address)
func (_Meme *MemeTransactorSession) CreateMeme(name string, symbol string, description string, icon string, website string, twitter string, telegram string, discord string, github string) (*types.Transaction, error) {
	return _Meme.Contract.CreateMeme(&_Meme.TransactOpts, name, symbol, description, icon, website, twitter, telegram, discord, github)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _x, uint256 _y, uint256 _fundingGoal) returns()
func (_Meme *MemeTransactor) Initialize(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _fundingGoal *big.Int) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "initialize", _x, _y, _fundingGoal)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _x, uint256 _y, uint256 _fundingGoal) returns()
func (_Meme *MemeSession) Initialize(_x *big.Int, _y *big.Int, _fundingGoal *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.Initialize(&_Meme.TransactOpts, _x, _y, _fundingGoal)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _x, uint256 _y, uint256 _fundingGoal) returns()
func (_Meme *MemeTransactorSession) Initialize(_x *big.Int, _y *big.Int, _fundingGoal *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.Initialize(&_Meme.TransactOpts, _x, _y, _fundingGoal)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Meme *MemeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Meme *MemeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Meme.Contract.RenounceOwnership(&_Meme.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Meme *MemeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Meme.Contract.RenounceOwnership(&_Meme.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tokenAddress, uint256 amount) returns()
func (_Meme *MemeTransactor) Sell(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "sell", tokenAddress, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tokenAddress, uint256 amount) returns()
func (_Meme *MemeSession) Sell(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.Sell(&_Meme.TransactOpts, tokenAddress, amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tokenAddress, uint256 amount) returns()
func (_Meme *MemeTransactorSession) Sell(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.Sell(&_Meme.TransactOpts, tokenAddress, amount)
}

// SetCreateFee is a paid mutator transaction binding the contract method 0x62914849.
//
// Solidity: function setCreateFee(uint256 _createFee) returns()
func (_Meme *MemeTransactor) SetCreateFee(opts *bind.TransactOpts, _createFee *big.Int) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "setCreateFee", _createFee)
}

// SetCreateFee is a paid mutator transaction binding the contract method 0x62914849.
//
// Solidity: function setCreateFee(uint256 _createFee) returns()
func (_Meme *MemeSession) SetCreateFee(_createFee *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.SetCreateFee(&_Meme.TransactOpts, _createFee)
}

// SetCreateFee is a paid mutator transaction binding the contract method 0x62914849.
//
// Solidity: function setCreateFee(uint256 _createFee) returns()
func (_Meme *MemeTransactorSession) SetCreateFee(_createFee *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.SetCreateFee(&_Meme.TransactOpts, _createFee)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x7ce3489b.
//
// Solidity: function setFeePercent(uint256 _feePercent) returns()
func (_Meme *MemeTransactor) SetFeePercent(opts *bind.TransactOpts, _feePercent *big.Int) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "setFeePercent", _feePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x7ce3489b.
//
// Solidity: function setFeePercent(uint256 _feePercent) returns()
func (_Meme *MemeSession) SetFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.SetFeePercent(&_Meme.TransactOpts, _feePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0x7ce3489b.
//
// Solidity: function setFeePercent(uint256 _feePercent) returns()
func (_Meme *MemeTransactorSession) SetFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.SetFeePercent(&_Meme.TransactOpts, _feePercent)
}

// SetFundingGoal is a paid mutator transaction binding the contract method 0x576eac66.
//
// Solidity: function setFundingGoal(uint256 newFundingGoal) returns()
func (_Meme *MemeTransactor) SetFundingGoal(opts *bind.TransactOpts, newFundingGoal *big.Int) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "setFundingGoal", newFundingGoal)
}

// SetFundingGoal is a paid mutator transaction binding the contract method 0x576eac66.
//
// Solidity: function setFundingGoal(uint256 newFundingGoal) returns()
func (_Meme *MemeSession) SetFundingGoal(newFundingGoal *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.SetFundingGoal(&_Meme.TransactOpts, newFundingGoal)
}

// SetFundingGoal is a paid mutator transaction binding the contract method 0x576eac66.
//
// Solidity: function setFundingGoal(uint256 newFundingGoal) returns()
func (_Meme *MemeTransactorSession) SetFundingGoal(newFundingGoal *big.Int) (*types.Transaction, error) {
	return _Meme.Contract.SetFundingGoal(&_Meme.TransactOpts, newFundingGoal)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meme *MemeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meme *MemeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Meme.Contract.TransferOwnership(&_Meme.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meme *MemeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Meme.Contract.TransferOwnership(&_Meme.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Meme *MemeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Meme.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Meme *MemeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Meme.Contract.UpgradeToAndCall(&_Meme.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Meme *MemeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Meme.Contract.UpgradeToAndCall(&_Meme.TransactOpts, newImplementation, data)
}

// MemeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Meme contract.
type MemeInitializedIterator struct {
	Event *MemeInitialized // Event containing the contract specifics and raw log

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
func (it *MemeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeInitialized)
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
		it.Event = new(MemeInitialized)
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
func (it *MemeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeInitialized represents a Initialized event raised by the Meme contract.
type MemeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Meme *MemeFilterer) FilterInitialized(opts *bind.FilterOpts) (*MemeInitializedIterator, error) {

	logs, sub, err := _Meme.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MemeInitializedIterator{contract: _Meme.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Meme *MemeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MemeInitialized) (event.Subscription, error) {

	logs, sub, err := _Meme.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeInitialized)
				if err := _Meme.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Meme *MemeFilterer) ParseInitialized(log types.Log) (*MemeInitialized, error) {
	event := new(MemeInitialized)
	if err := _Meme.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Meme contract.
type MemeOwnershipTransferredIterator struct {
	Event *MemeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MemeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeOwnershipTransferred)
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
		it.Event = new(MemeOwnershipTransferred)
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
func (it *MemeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeOwnershipTransferred represents a OwnershipTransferred event raised by the Meme contract.
type MemeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Meme *MemeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MemeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Meme.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MemeOwnershipTransferredIterator{contract: _Meme.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Meme *MemeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MemeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Meme.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeOwnershipTransferred)
				if err := _Meme.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Meme *MemeFilterer) ParseOwnershipTransferred(log types.Log) (*MemeOwnershipTransferred, error) {
	event := new(MemeOwnershipTransferred)
	if err := _Meme.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeTokenBuyIterator is returned from FilterTokenBuy and is used to iterate over the raw logs and unpacked data for TokenBuy events raised by the Meme contract.
type MemeTokenBuyIterator struct {
	Event *MemeTokenBuy // Event containing the contract specifics and raw log

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
func (it *MemeTokenBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeTokenBuy)
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
		it.Event = new(MemeTokenBuy)
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
func (it *MemeTokenBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeTokenBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeTokenBuy represents a TokenBuy event raised by the Meme contract.
type MemeTokenBuy struct {
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
func (_Meme *MemeFilterer) FilterTokenBuy(opts *bind.FilterOpts, token []common.Address) (*MemeTokenBuyIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Meme.contract.FilterLogs(opts, "TokenBuy", tokenRule)
	if err != nil {
		return nil, err
	}
	return &MemeTokenBuyIterator{contract: _Meme.contract, event: "TokenBuy", logs: logs, sub: sub}, nil
}

// WatchTokenBuy is a free log subscription operation binding the contract event 0x6c6bf9b07a057cb0766b2b0aefeb3ed76d844744cc34e081c60bb230ef57f65d.
//
// Solidity: event TokenBuy(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_Meme *MemeFilterer) WatchTokenBuy(opts *bind.WatchOpts, sink chan<- *MemeTokenBuy, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Meme.contract.WatchLogs(opts, "TokenBuy", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeTokenBuy)
				if err := _Meme.contract.UnpackLog(event, "TokenBuy", log); err != nil {
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
func (_Meme *MemeFilterer) ParseTokenBuy(log types.Log) (*MemeTokenBuy, error) {
	event := new(MemeTokenBuy)
	if err := _Meme.contract.UnpackLog(event, "TokenBuy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the Meme contract.
type MemeTokenCreatedIterator struct {
	Event *MemeTokenCreated // Event containing the contract specifics and raw log

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
func (it *MemeTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeTokenCreated)
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
		it.Event = new(MemeTokenCreated)
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
func (it *MemeTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeTokenCreated represents a TokenCreated event raised by the Meme contract.
type MemeTokenCreated struct {
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
func (_Meme *MemeFilterer) FilterTokenCreated(opts *bind.FilterOpts, token []common.Address, creator []common.Address) (*MemeTokenCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Meme.contract.FilterLogs(opts, "TokenCreated", tokenRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &MemeTokenCreatedIterator{contract: _Meme.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0x1a2bd37fcc163f5d288090a0f9d97dc7dcf61b340620c5da1049f9cc413a9708.
//
// Solidity: event TokenCreated(address indexed token, address indexed creator, uint256 timestamp, string name, string symbol, string description, string icon, string website, string twitter, string telegram, string discord, string github)
func (_Meme *MemeFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *MemeTokenCreated, token []common.Address, creator []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Meme.contract.WatchLogs(opts, "TokenCreated", tokenRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeTokenCreated)
				if err := _Meme.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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
func (_Meme *MemeFilterer) ParseTokenCreated(log types.Log) (*MemeTokenCreated, error) {
	event := new(MemeTokenCreated)
	if err := _Meme.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeTokenSellIterator is returned from FilterTokenSell and is used to iterate over the raw logs and unpacked data for TokenSell events raised by the Meme contract.
type MemeTokenSellIterator struct {
	Event *MemeTokenSell // Event containing the contract specifics and raw log

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
func (it *MemeTokenSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeTokenSell)
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
		it.Event = new(MemeTokenSell)
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
func (it *MemeTokenSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeTokenSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeTokenSell represents a TokenSell event raised by the Meme contract.
type MemeTokenSell struct {
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
func (_Meme *MemeFilterer) FilterTokenSell(opts *bind.FilterOpts, token []common.Address) (*MemeTokenSellIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Meme.contract.FilterLogs(opts, "TokenSell", tokenRule)
	if err != nil {
		return nil, err
	}
	return &MemeTokenSellIterator{contract: _Meme.contract, event: "TokenSell", logs: logs, sub: sub}, nil
}

// WatchTokenSell is a free log subscription operation binding the contract event 0xb40ebf7368df31a675c3b3a523e81f6b28b5cf6e4b32586ed0bc6a86052d8991.
//
// Solidity: event TokenSell(address indexed token, uint256 timestamp, uint256 baseAmount, uint256 quoteAmount, uint256 feeAmount)
func (_Meme *MemeFilterer) WatchTokenSell(opts *bind.WatchOpts, sink chan<- *MemeTokenSell, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Meme.contract.WatchLogs(opts, "TokenSell", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeTokenSell)
				if err := _Meme.contract.UnpackLog(event, "TokenSell", log); err != nil {
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
func (_Meme *MemeFilterer) ParseTokenSell(log types.Log) (*MemeTokenSell, error) {
	event := new(MemeTokenSell)
	if err := _Meme.contract.UnpackLog(event, "TokenSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Meme contract.
type MemeUpgradedIterator struct {
	Event *MemeUpgraded // Event containing the contract specifics and raw log

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
func (it *MemeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemeUpgraded)
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
		it.Event = new(MemeUpgraded)
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
func (it *MemeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemeUpgraded represents a Upgraded event raised by the Meme contract.
type MemeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Meme *MemeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MemeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Meme.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MemeUpgradedIterator{contract: _Meme.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Meme *MemeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MemeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Meme.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemeUpgraded)
				if err := _Meme.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Meme *MemeFilterer) ParseUpgraded(log types.Log) (*MemeUpgraded, error) {
	event := new(MemeUpgraded)
	if err := _Meme.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
