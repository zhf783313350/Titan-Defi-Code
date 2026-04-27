// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// AuctionMarketMetaData contains all meta data concerning the AuctionMarket contract.
var AuctionMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minPriceInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"highestBidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[{\"name\":\"_auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minPriceUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[{\"name\":\"_auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalAuctions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minPriceInUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidPlaced\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// AuctionMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMarketMetaData.ABI instead.
var AuctionMarketABI = AuctionMarketMetaData.ABI

// AuctionMarket is an auto generated Go binding around an Ethereum contract.
type AuctionMarket struct {
	AuctionMarketCaller     // Read-only binding to the contract
	AuctionMarketTransactor // Write-only binding to the contract
	AuctionMarketFilterer   // Log filterer for contract events
}

// AuctionMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionMarketSession struct {
	Contract     *AuctionMarket    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionMarketCallerSession struct {
	Contract *AuctionMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AuctionMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionMarketTransactorSession struct {
	Contract     *AuctionMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AuctionMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionMarketRaw struct {
	Contract *AuctionMarket // Generic contract binding to access the raw methods on
}

// AuctionMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionMarketCallerRaw struct {
	Contract *AuctionMarketCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionMarketTransactorRaw struct {
	Contract *AuctionMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionMarket creates a new instance of AuctionMarket, bound to a specific deployed contract.
func NewAuctionMarket(address common.Address, backend bind.ContractBackend) (*AuctionMarket, error) {
	contract, err := bindAuctionMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionMarket{AuctionMarketCaller: AuctionMarketCaller{contract: contract}, AuctionMarketTransactor: AuctionMarketTransactor{contract: contract}, AuctionMarketFilterer: AuctionMarketFilterer{contract: contract}}, nil
}

// NewAuctionMarketCaller creates a new read-only instance of AuctionMarket, bound to a specific deployed contract.
func NewAuctionMarketCaller(address common.Address, caller bind.ContractCaller) (*AuctionMarketCaller, error) {
	contract, err := bindAuctionMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketCaller{contract: contract}, nil
}

// NewAuctionMarketTransactor creates a new write-only instance of AuctionMarket, bound to a specific deployed contract.
func NewAuctionMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionMarketTransactor, error) {
	contract, err := bindAuctionMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketTransactor{contract: contract}, nil
}

// NewAuctionMarketFilterer creates a new log filterer instance of AuctionMarket, bound to a specific deployed contract.
func NewAuctionMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionMarketFilterer, error) {
	contract, err := bindAuctionMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketFilterer{contract: contract}, nil
}

// bindAuctionMarket binds a generic wrapper to an already deployed contract.
func bindAuctionMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionMarket *AuctionMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionMarket.Contract.AuctionMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionMarket *AuctionMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarket.Contract.AuctionMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionMarket *AuctionMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionMarket.Contract.AuctionMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionMarket *AuctionMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionMarket *AuctionMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionMarket *AuctionMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionMarket.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AuctionMarket *AuctionMarketCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuctionMarket.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AuctionMarket *AuctionMarketSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AuctionMarket.Contract.UPGRADEINTERFACEVERSION(&_AuctionMarket.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AuctionMarket *AuctionMarketCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AuctionMarket.Contract.UPGRADEINTERFACEVERSION(&_AuctionMarket.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftAddr, uint256 tokenId, uint256 minPriceInUsd, address highestBidder, uint256 highestBidAmount, uint256 endTime, bool isActive)
func (_AuctionMarket *AuctionMarketCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller           common.Address
	NftAddr          common.Address
	TokenId          *big.Int
	MinPriceInUsd    *big.Int
	HighestBidder    common.Address
	HighestBidAmount *big.Int
	EndTime          *big.Int
	IsActive         bool
}, error) {
	var out []interface{}
	err := _AuctionMarket.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Seller           common.Address
		NftAddr          common.Address
		TokenId          *big.Int
		MinPriceInUsd    *big.Int
		HighestBidder    common.Address
		HighestBidAmount *big.Int
		EndTime          *big.Int
		IsActive         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinPriceInUsd = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.HighestBidAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftAddr, uint256 tokenId, uint256 minPriceInUsd, address highestBidder, uint256 highestBidAmount, uint256 endTime, bool isActive)
func (_AuctionMarket *AuctionMarketSession) Auctions(arg0 *big.Int) (struct {
	Seller           common.Address
	NftAddr          common.Address
	TokenId          *big.Int
	MinPriceInUsd    *big.Int
	HighestBidder    common.Address
	HighestBidAmount *big.Int
	EndTime          *big.Int
	IsActive         bool
}, error) {
	return _AuctionMarket.Contract.Auctions(&_AuctionMarket.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftAddr, uint256 tokenId, uint256 minPriceInUsd, address highestBidder, uint256 highestBidAmount, uint256 endTime, bool isActive)
func (_AuctionMarket *AuctionMarketCallerSession) Auctions(arg0 *big.Int) (struct {
	Seller           common.Address
	NftAddr          common.Address
	TokenId          *big.Int
	MinPriceInUsd    *big.Int
	HighestBidder    common.Address
	HighestBidAmount *big.Int
	EndTime          *big.Int
	IsActive         bool
}, error) {
	return _AuctionMarket.Contract.Auctions(&_AuctionMarket.CallOpts, arg0)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_AuctionMarket *AuctionMarketCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarket.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_AuctionMarket *AuctionMarketSession) GetLatestPrice() (*big.Int, error) {
	return _AuctionMarket.Contract.GetLatestPrice(&_AuctionMarket.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_AuctionMarket *AuctionMarketCallerSession) GetLatestPrice() (*big.Int, error) {
	return _AuctionMarket.Contract.GetLatestPrice(&_AuctionMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionMarket *AuctionMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionMarket *AuctionMarketSession) Owner() (common.Address, error) {
	return _AuctionMarket.Contract.Owner(&_AuctionMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionMarket *AuctionMarketCallerSession) Owner() (common.Address, error) {
	return _AuctionMarket.Contract.Owner(&_AuctionMarket.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuctionMarket *AuctionMarketCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuctionMarket.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuctionMarket *AuctionMarketSession) ProxiableUUID() ([32]byte, error) {
	return _AuctionMarket.Contract.ProxiableUUID(&_AuctionMarket.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuctionMarket *AuctionMarketCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AuctionMarket.Contract.ProxiableUUID(&_AuctionMarket.CallOpts)
}

// TotalAuctions is a free data retrieval call binding the contract method 0x16002f4a.
//
// Solidity: function totalAuctions() view returns(uint256)
func (_AuctionMarket *AuctionMarketCaller) TotalAuctions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarket.contract.Call(opts, &out, "totalAuctions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAuctions is a free data retrieval call binding the contract method 0x16002f4a.
//
// Solidity: function totalAuctions() view returns(uint256)
func (_AuctionMarket *AuctionMarketSession) TotalAuctions() (*big.Int, error) {
	return _AuctionMarket.Contract.TotalAuctions(&_AuctionMarket.CallOpts)
}

// TotalAuctions is a free data retrieval call binding the contract method 0x16002f4a.
//
// Solidity: function totalAuctions() view returns(uint256)
func (_AuctionMarket *AuctionMarketCallerSession) TotalAuctions() (*big.Int, error) {
	return _AuctionMarket.Contract.TotalAuctions(&_AuctionMarket.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns()
func (_AuctionMarket *AuctionMarketTransactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "bid", _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns()
func (_AuctionMarket *AuctionMarketSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.Contract.Bid(&_AuctionMarket.TransactOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns()
func (_AuctionMarket *AuctionMarketTransactorSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.Contract.Bid(&_AuctionMarket.TransactOpts, _auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x61beb1d7.
//
// Solidity: function createAuction(address _nftAddr, uint256 _tokenId, uint256 _minPriceUsd, uint256 _duration) returns()
func (_AuctionMarket *AuctionMarketTransactor) CreateAuction(opts *bind.TransactOpts, _nftAddr common.Address, _tokenId *big.Int, _minPriceUsd *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "createAuction", _nftAddr, _tokenId, _minPriceUsd, _duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x61beb1d7.
//
// Solidity: function createAuction(address _nftAddr, uint256 _tokenId, uint256 _minPriceUsd, uint256 _duration) returns()
func (_AuctionMarket *AuctionMarketSession) CreateAuction(_nftAddr common.Address, _tokenId *big.Int, _minPriceUsd *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.Contract.CreateAuction(&_AuctionMarket.TransactOpts, _nftAddr, _tokenId, _minPriceUsd, _duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x61beb1d7.
//
// Solidity: function createAuction(address _nftAddr, uint256 _tokenId, uint256 _minPriceUsd, uint256 _duration) returns()
func (_AuctionMarket *AuctionMarketTransactorSession) CreateAuction(_nftAddr common.Address, _tokenId *big.Int, _minPriceUsd *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.Contract.CreateAuction(&_AuctionMarket.TransactOpts, _nftAddr, _tokenId, _minPriceUsd, _duration)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_AuctionMarket *AuctionMarketTransactor) EndAuction(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "endAuction", _auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_AuctionMarket *AuctionMarketSession) EndAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.Contract.EndAuction(&_AuctionMarket.TransactOpts, _auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_AuctionMarket *AuctionMarketTransactorSession) EndAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionMarket.Contract.EndAuction(&_AuctionMarket.TransactOpts, _auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _priceFeed) returns()
func (_AuctionMarket *AuctionMarketTransactor) Initialize(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "initialize", _priceFeed)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _priceFeed) returns()
func (_AuctionMarket *AuctionMarketSession) Initialize(_priceFeed common.Address) (*types.Transaction, error) {
	return _AuctionMarket.Contract.Initialize(&_AuctionMarket.TransactOpts, _priceFeed)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _priceFeed) returns()
func (_AuctionMarket *AuctionMarketTransactorSession) Initialize(_priceFeed common.Address) (*types.Transaction, error) {
	return _AuctionMarket.Contract.Initialize(&_AuctionMarket.TransactOpts, _priceFeed)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionMarket *AuctionMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionMarket *AuctionMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionMarket.Contract.RenounceOwnership(&_AuctionMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionMarket *AuctionMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionMarket.Contract.RenounceOwnership(&_AuctionMarket.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionMarket *AuctionMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionMarket *AuctionMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionMarket.Contract.TransferOwnership(&_AuctionMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionMarket *AuctionMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionMarket.Contract.TransferOwnership(&_AuctionMarket.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuctionMarket *AuctionMarketTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuctionMarket.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuctionMarket *AuctionMarketSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuctionMarket.Contract.UpgradeToAndCall(&_AuctionMarket.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuctionMarket *AuctionMarketTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuctionMarket.Contract.UpgradeToAndCall(&_AuctionMarket.TransactOpts, newImplementation, data)
}

// AuctionMarketAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the AuctionMarket contract.
type AuctionMarketAuctionCreatedIterator struct {
	Event *AuctionMarketAuctionCreated // Event containing the contract specifics and raw log

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
func (it *AuctionMarketAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketAuctionCreated)
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
		it.Event = new(AuctionMarketAuctionCreated)
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
func (it *AuctionMarketAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketAuctionCreated represents a AuctionCreated event raised by the AuctionMarket contract.
type AuctionMarketAuctionCreated struct {
	AuctionId     *big.Int
	Seller        common.Address
	TokenId       *big.Int
	MinPriceInUsd *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xc9050d42180a61cb0d9ebb8ad118b62fe6eab12cf12ff752c4a0cc7da9ddf627.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, uint256 tokenId, uint256 minPriceInUsd)
func (_AuctionMarket *AuctionMarketFilterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address) (*AuctionMarketAuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _AuctionMarket.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketAuctionCreatedIterator{contract: _AuctionMarket.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xc9050d42180a61cb0d9ebb8ad118b62fe6eab12cf12ff752c4a0cc7da9ddf627.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, uint256 tokenId, uint256 minPriceInUsd)
func (_AuctionMarket *AuctionMarketFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *AuctionMarketAuctionCreated, auctionId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _AuctionMarket.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketAuctionCreated)
				if err := _AuctionMarket.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xc9050d42180a61cb0d9ebb8ad118b62fe6eab12cf12ff752c4a0cc7da9ddf627.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, uint256 tokenId, uint256 minPriceInUsd)
func (_AuctionMarket *AuctionMarketFilterer) ParseAuctionCreated(log types.Log) (*AuctionMarketAuctionCreated, error) {
	event := new(AuctionMarketAuctionCreated)
	if err := _AuctionMarket.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the AuctionMarket contract.
type AuctionMarketBidPlacedIterator struct {
	Event *AuctionMarketBidPlaced // Event containing the contract specifics and raw log

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
func (it *AuctionMarketBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketBidPlaced)
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
		it.Event = new(AuctionMarketBidPlaced)
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
func (it *AuctionMarketBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketBidPlaced represents a BidPlaced event raised by the AuctionMarket contract.
type AuctionMarketBidPlaced struct {
	AuctionId *big.Int
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0e54eff26401bf69b81b26f60bd85ef47f5d85275c1d268d84f68d6897431c47.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount)
func (_AuctionMarket *AuctionMarketFilterer) FilterBidPlaced(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address) (*AuctionMarketBidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _AuctionMarket.contract.FilterLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketBidPlacedIterator{contract: _AuctionMarket.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0e54eff26401bf69b81b26f60bd85ef47f5d85275c1d268d84f68d6897431c47.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount)
func (_AuctionMarket *AuctionMarketFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *AuctionMarketBidPlaced, auctionId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _AuctionMarket.contract.WatchLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketBidPlaced)
				if err := _AuctionMarket.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x0e54eff26401bf69b81b26f60bd85ef47f5d85275c1d268d84f68d6897431c47.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount)
func (_AuctionMarket *AuctionMarketFilterer) ParseBidPlaced(log types.Log) (*AuctionMarketBidPlaced, error) {
	event := new(AuctionMarketBidPlaced)
	if err := _AuctionMarket.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AuctionMarket contract.
type AuctionMarketInitializedIterator struct {
	Event *AuctionMarketInitialized // Event containing the contract specifics and raw log

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
func (it *AuctionMarketInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketInitialized)
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
		it.Event = new(AuctionMarketInitialized)
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
func (it *AuctionMarketInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketInitialized represents a Initialized event raised by the AuctionMarket contract.
type AuctionMarketInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuctionMarket *AuctionMarketFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuctionMarketInitializedIterator, error) {

	logs, sub, err := _AuctionMarket.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketInitializedIterator{contract: _AuctionMarket.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuctionMarket *AuctionMarketFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuctionMarketInitialized) (event.Subscription, error) {

	logs, sub, err := _AuctionMarket.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketInitialized)
				if err := _AuctionMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AuctionMarket *AuctionMarketFilterer) ParseInitialized(log types.Log) (*AuctionMarketInitialized, error) {
	event := new(AuctionMarketInitialized)
	if err := _AuctionMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionMarket contract.
type AuctionMarketOwnershipTransferredIterator struct {
	Event *AuctionMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketOwnershipTransferred)
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
		it.Event = new(AuctionMarketOwnershipTransferred)
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
func (it *AuctionMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionMarket contract.
type AuctionMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionMarket *AuctionMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketOwnershipTransferredIterator{contract: _AuctionMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionMarket *AuctionMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketOwnershipTransferred)
				if err := _AuctionMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuctionMarket *AuctionMarketFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionMarketOwnershipTransferred, error) {
	event := new(AuctionMarketOwnershipTransferred)
	if err := _AuctionMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AuctionMarket contract.
type AuctionMarketUpgradedIterator struct {
	Event *AuctionMarketUpgraded // Event containing the contract specifics and raw log

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
func (it *AuctionMarketUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketUpgraded)
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
		it.Event = new(AuctionMarketUpgraded)
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
func (it *AuctionMarketUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketUpgraded represents a Upgraded event raised by the AuctionMarket contract.
type AuctionMarketUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuctionMarket *AuctionMarketFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuctionMarketUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuctionMarket.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketUpgradedIterator{contract: _AuctionMarket.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuctionMarket *AuctionMarketFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuctionMarketUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuctionMarket.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketUpgraded)
				if err := _AuctionMarket.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AuctionMarket *AuctionMarketFilterer) ParseUpgraded(log types.Log) (*AuctionMarketUpgraded, error) {
	event := new(AuctionMarketUpgraded)
	if err := _AuctionMarket.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
