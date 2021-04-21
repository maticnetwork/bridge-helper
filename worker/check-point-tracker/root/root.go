// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package root

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

// RootABI is the input ABI used to generate the binding from.
const RootABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"headerBlockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewHeaderBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"headerBlockId\",\"type\":\"uint256\"}],\"name\":\"ResetHeaderBlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VOTE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_nextHeaderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"headerBlocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heimdallId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"submitHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numDeposits\",\"type\":\"uint256\"}],\"name\":\"updateDepositId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentHeaderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNextHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_heimdallId\",\"type\":\"string\"}],\"name\":\"setHeimdallId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Root is an auto generated Go binding around an Ethereum contract.
type Root struct {
	RootCaller     // Read-only binding to the contract
	RootTransactor // Write-only binding to the contract
	RootFilterer   // Log filterer for contract events
}

// RootCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootSession struct {
	Contract     *Root             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootCallerSession struct {
	Contract *RootCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootTransactorSession struct {
	Contract     *RootTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootRaw struct {
	Contract *Root // Generic contract binding to access the raw methods on
}

// RootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootCallerRaw struct {
	Contract *RootCaller // Generic read-only contract binding to access the raw methods on
}

// RootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootTransactorRaw struct {
	Contract *RootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoot creates a new instance of Root, bound to a specific deployed contract.
func NewRoot(address common.Address, backend bind.ContractBackend) (*Root, error) {
	contract, err := bindRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Root{RootCaller: RootCaller{contract: contract}, RootTransactor: RootTransactor{contract: contract}, RootFilterer: RootFilterer{contract: contract}}, nil
}

// NewRootCaller creates a new read-only instance of Root, bound to a specific deployed contract.
func NewRootCaller(address common.Address, caller bind.ContractCaller) (*RootCaller, error) {
	contract, err := bindRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootCaller{contract: contract}, nil
}

// NewRootTransactor creates a new write-only instance of Root, bound to a specific deployed contract.
func NewRootTransactor(address common.Address, transactor bind.ContractTransactor) (*RootTransactor, error) {
	contract, err := bindRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootTransactor{contract: contract}, nil
}

// NewRootFilterer creates a new log filterer instance of Root, bound to a specific deployed contract.
func NewRootFilterer(address common.Address, filterer bind.ContractFilterer) (*RootFilterer, error) {
	contract, err := bindRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootFilterer{contract: contract}, nil
}

// bindRoot binds a generic wrapper to an already deployed contract.
func bindRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Root *RootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Root.Contract.RootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Root *RootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.Contract.RootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Root *RootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Root.Contract.RootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Root *RootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Root.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Root *RootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Root *RootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Root.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_Root *RootCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "CHAINID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_Root *RootSession) CHAINID() (*big.Int, error) {
	return _Root.Contract.CHAINID(&_Root.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_Root *RootCallerSession) CHAINID() (*big.Int, error) {
	return _Root.Contract.CHAINID(&_Root.CallOpts)
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_Root *RootCaller) VOTETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "VOTE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_Root *RootSession) VOTETYPE() (uint8, error) {
	return _Root.Contract.VOTETYPE(&_Root.CallOpts)
}

// VOTETYPE is a free data retrieval call binding the contract method 0xd5b844eb.
//
// Solidity: function VOTE_TYPE() view returns(uint8)
func (_Root *RootCallerSession) VOTETYPE() (uint8, error) {
	return _Root.Contract.VOTETYPE(&_Root.CallOpts)
}

// NextHeaderBlock is a free data retrieval call binding the contract method 0x8d978d88.
//
// Solidity: function _nextHeaderBlock() view returns(uint256)
func (_Root *RootCaller) NextHeaderBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "_nextHeaderBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextHeaderBlock is a free data retrieval call binding the contract method 0x8d978d88.
//
// Solidity: function _nextHeaderBlock() view returns(uint256)
func (_Root *RootSession) NextHeaderBlock() (*big.Int, error) {
	return _Root.Contract.NextHeaderBlock(&_Root.CallOpts)
}

// NextHeaderBlock is a free data retrieval call binding the contract method 0x8d978d88.
//
// Solidity: function _nextHeaderBlock() view returns(uint256)
func (_Root *RootCallerSession) NextHeaderBlock() (*big.Int, error) {
	return _Root.Contract.NextHeaderBlock(&_Root.CallOpts)
}

// CurrentHeaderBlock is a free data retrieval call binding the contract method 0xec7e4855.
//
// Solidity: function currentHeaderBlock() view returns(uint256)
func (_Root *RootCaller) CurrentHeaderBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "currentHeaderBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentHeaderBlock is a free data retrieval call binding the contract method 0xec7e4855.
//
// Solidity: function currentHeaderBlock() view returns(uint256)
func (_Root *RootSession) CurrentHeaderBlock() (*big.Int, error) {
	return _Root.Contract.CurrentHeaderBlock(&_Root.CallOpts)
}

// CurrentHeaderBlock is a free data retrieval call binding the contract method 0xec7e4855.
//
// Solidity: function currentHeaderBlock() view returns(uint256)
func (_Root *RootCallerSession) CurrentHeaderBlock() (*big.Int, error) {
	return _Root.Contract.CurrentHeaderBlock(&_Root.CallOpts)
}

// GetLastChildBlock is a free data retrieval call binding the contract method 0xb87e1b66.
//
// Solidity: function getLastChildBlock() view returns(uint256)
func (_Root *RootCaller) GetLastChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "getLastChildBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastChildBlock is a free data retrieval call binding the contract method 0xb87e1b66.
//
// Solidity: function getLastChildBlock() view returns(uint256)
func (_Root *RootSession) GetLastChildBlock() (*big.Int, error) {
	return _Root.Contract.GetLastChildBlock(&_Root.CallOpts)
}

// GetLastChildBlock is a free data retrieval call binding the contract method 0xb87e1b66.
//
// Solidity: function getLastChildBlock() view returns(uint256)
func (_Root *RootCallerSession) GetLastChildBlock() (*big.Int, error) {
	return _Root.Contract.GetLastChildBlock(&_Root.CallOpts)
}

// HeaderBlocks is a free data retrieval call binding the contract method 0x41539d4a.
//
// Solidity: function headerBlocks(uint256 ) view returns(bytes32 root, uint256 start, uint256 end, uint256 createdAt, address proposer)
func (_Root *RootCaller) HeaderBlocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Start     *big.Int
	End       *big.Int
	CreatedAt *big.Int
	Proposer  common.Address
}, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "headerBlocks", arg0)

	outstruct := new(struct {
		Root      [32]byte
		Start     *big.Int
		End       *big.Int
		CreatedAt *big.Int
		Proposer  common.Address
	})

	outstruct.Root = out[0].([32]byte)
	outstruct.Start = out[1].(*big.Int)
	outstruct.End = out[2].(*big.Int)
	outstruct.CreatedAt = out[3].(*big.Int)
	outstruct.Proposer = out[4].(common.Address)

	return *outstruct, err

}

// HeaderBlocks is a free data retrieval call binding the contract method 0x41539d4a.
//
// Solidity: function headerBlocks(uint256 ) view returns(bytes32 root, uint256 start, uint256 end, uint256 createdAt, address proposer)
func (_Root *RootSession) HeaderBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Start     *big.Int
	End       *big.Int
	CreatedAt *big.Int
	Proposer  common.Address
}, error) {
	return _Root.Contract.HeaderBlocks(&_Root.CallOpts, arg0)
}

// HeaderBlocks is a free data retrieval call binding the contract method 0x41539d4a.
//
// Solidity: function headerBlocks(uint256 ) view returns(bytes32 root, uint256 start, uint256 end, uint256 createdAt, address proposer)
func (_Root *RootCallerSession) HeaderBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Start     *big.Int
	End       *big.Int
	CreatedAt *big.Int
	Proposer  common.Address
}, error) {
	return _Root.Contract.HeaderBlocks(&_Root.CallOpts, arg0)
}

// HeimdallId is a free data retrieval call binding the contract method 0xfbc3dd36.
//
// Solidity: function heimdallId() view returns(bytes32)
func (_Root *RootCaller) HeimdallId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "heimdallId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HeimdallId is a free data retrieval call binding the contract method 0xfbc3dd36.
//
// Solidity: function heimdallId() view returns(bytes32)
func (_Root *RootSession) HeimdallId() ([32]byte, error) {
	return _Root.Contract.HeimdallId(&_Root.CallOpts)
}

// HeimdallId is a free data retrieval call binding the contract method 0xfbc3dd36.
//
// Solidity: function heimdallId() view returns(bytes32)
func (_Root *RootCallerSession) HeimdallId() ([32]byte, error) {
	return _Root.Contract.HeimdallId(&_Root.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Root *RootCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Root *RootSession) IsOwner() (bool, error) {
	return _Root.Contract.IsOwner(&_Root.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Root *RootCallerSession) IsOwner() (bool, error) {
	return _Root.Contract.IsOwner(&_Root.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_Root *RootCaller) NetworkId(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "networkId")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_Root *RootSession) NetworkId() ([]byte, error) {
	return _Root.Contract.NetworkId(&_Root.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_Root *RootCallerSession) NetworkId() ([]byte, error) {
	return _Root.Contract.NetworkId(&_Root.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Root *RootCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Root.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Root *RootSession) Owner() (common.Address, error) {
	return _Root.Contract.Owner(&_Root.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Root *RootCallerSession) Owner() (common.Address, error) {
	return _Root.Contract.Owner(&_Root.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Root *RootTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Root *RootSession) RenounceOwnership() (*types.Transaction, error) {
	return _Root.Contract.RenounceOwnership(&_Root.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Root *RootTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Root.Contract.RenounceOwnership(&_Root.TransactOpts)
}

// SetHeimdallId is a paid mutator transaction binding the contract method 0xea0688b3.
//
// Solidity: function setHeimdallId(string _heimdallId) returns()
func (_Root *RootTransactor) SetHeimdallId(opts *bind.TransactOpts, _heimdallId string) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setHeimdallId", _heimdallId)
}

// SetHeimdallId is a paid mutator transaction binding the contract method 0xea0688b3.
//
// Solidity: function setHeimdallId(string _heimdallId) returns()
func (_Root *RootSession) SetHeimdallId(_heimdallId string) (*types.Transaction, error) {
	return _Root.Contract.SetHeimdallId(&_Root.TransactOpts, _heimdallId)
}

// SetHeimdallId is a paid mutator transaction binding the contract method 0xea0688b3.
//
// Solidity: function setHeimdallId(string _heimdallId) returns()
func (_Root *RootTransactorSession) SetHeimdallId(_heimdallId string) (*types.Transaction, error) {
	return _Root.Contract.SetHeimdallId(&_Root.TransactOpts, _heimdallId)
}

// SetNextHeaderBlock is a paid mutator transaction binding the contract method 0xcf24a0ea.
//
// Solidity: function setNextHeaderBlock(uint256 _value) returns()
func (_Root *RootTransactor) SetNextHeaderBlock(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setNextHeaderBlock", _value)
}

// SetNextHeaderBlock is a paid mutator transaction binding the contract method 0xcf24a0ea.
//
// Solidity: function setNextHeaderBlock(uint256 _value) returns()
func (_Root *RootSession) SetNextHeaderBlock(_value *big.Int) (*types.Transaction, error) {
	return _Root.Contract.SetNextHeaderBlock(&_Root.TransactOpts, _value)
}

// SetNextHeaderBlock is a paid mutator transaction binding the contract method 0xcf24a0ea.
//
// Solidity: function setNextHeaderBlock(uint256 _value) returns()
func (_Root *RootTransactorSession) SetNextHeaderBlock(_value *big.Int) (*types.Transaction, error) {
	return _Root.Contract.SetNextHeaderBlock(&_Root.TransactOpts, _value)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_Root *RootTransactor) Slash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "slash")
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_Root *RootSession) Slash() (*types.Transaction, error) {
	return _Root.Contract.Slash(&_Root.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_Root *RootTransactorSession) Slash() (*types.Transaction, error) {
	return _Root.Contract.Slash(&_Root.TransactOpts)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0x6a791f11.
//
// Solidity: function submitHeaderBlock(bytes data, bytes sigs) returns()
func (_Root *RootTransactor) SubmitHeaderBlock(opts *bind.TransactOpts, data []byte, sigs []byte) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "submitHeaderBlock", data, sigs)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0x6a791f11.
//
// Solidity: function submitHeaderBlock(bytes data, bytes sigs) returns()
func (_Root *RootSession) SubmitHeaderBlock(data []byte, sigs []byte) (*types.Transaction, error) {
	return _Root.Contract.SubmitHeaderBlock(&_Root.TransactOpts, data, sigs)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0x6a791f11.
//
// Solidity: function submitHeaderBlock(bytes data, bytes sigs) returns()
func (_Root *RootTransactorSession) SubmitHeaderBlock(data []byte, sigs []byte) (*types.Transaction, error) {
	return _Root.Contract.SubmitHeaderBlock(&_Root.TransactOpts, data, sigs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Root *RootTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Root *RootSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Root.Contract.TransferOwnership(&_Root.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Root *RootTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Root.Contract.TransferOwnership(&_Root.TransactOpts, newOwner)
}

// UpdateDepositId is a paid mutator transaction binding the contract method 0x5391f483.
//
// Solidity: function updateDepositId(uint256 numDeposits) returns(uint256 depositId)
func (_Root *RootTransactor) UpdateDepositId(opts *bind.TransactOpts, numDeposits *big.Int) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "updateDepositId", numDeposits)
}

// UpdateDepositId is a paid mutator transaction binding the contract method 0x5391f483.
//
// Solidity: function updateDepositId(uint256 numDeposits) returns(uint256 depositId)
func (_Root *RootSession) UpdateDepositId(numDeposits *big.Int) (*types.Transaction, error) {
	return _Root.Contract.UpdateDepositId(&_Root.TransactOpts, numDeposits)
}

// UpdateDepositId is a paid mutator transaction binding the contract method 0x5391f483.
//
// Solidity: function updateDepositId(uint256 numDeposits) returns(uint256 depositId)
func (_Root *RootTransactorSession) UpdateDepositId(numDeposits *big.Int) (*types.Transaction, error) {
	return _Root.Contract.UpdateDepositId(&_Root.TransactOpts, numDeposits)
}

// RootNewHeaderBlockIterator is returned from FilterNewHeaderBlock and is used to iterate over the raw logs and unpacked data for NewHeaderBlock events raised by the Root contract.
type RootNewHeaderBlockIterator struct {
	Event *RootNewHeaderBlock // Event containing the contract specifics and raw log

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
func (it *RootNewHeaderBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootNewHeaderBlock)
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
		it.Event = new(RootNewHeaderBlock)
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
func (it *RootNewHeaderBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootNewHeaderBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootNewHeaderBlock represents a NewHeaderBlock event raised by the Root contract.
type RootNewHeaderBlock struct {
	Proposer      common.Address
	HeaderBlockId *big.Int
	Reward        *big.Int
	Start         *big.Int
	End           *big.Int
	Root          [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewHeaderBlock is a free log retrieval operation binding the contract event 0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527.
//
// Solidity: event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)
func (_Root *RootFilterer) FilterNewHeaderBlock(opts *bind.FilterOpts, proposer []common.Address, headerBlockId []*big.Int, reward []*big.Int) (*RootNewHeaderBlockIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Root.contract.FilterLogs(opts, "NewHeaderBlock", proposerRule, headerBlockIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &RootNewHeaderBlockIterator{contract: _Root.contract, event: "NewHeaderBlock", logs: logs, sub: sub}, nil
}

// WatchNewHeaderBlock is a free log subscription operation binding the contract event 0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527.
//
// Solidity: event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)
func (_Root *RootFilterer) WatchNewHeaderBlock(opts *bind.WatchOpts, sink chan<- *RootNewHeaderBlock, proposer []common.Address, headerBlockId []*big.Int, reward []*big.Int) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _Root.contract.WatchLogs(opts, "NewHeaderBlock", proposerRule, headerBlockIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootNewHeaderBlock)
				if err := _Root.contract.UnpackLog(event, "NewHeaderBlock", log); err != nil {
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

// ParseNewHeaderBlock is a log parse operation binding the contract event 0xba5de06d22af2685c6c7765f60067f7d2b08c2d29f53cdf14d67f6d1c9bfb527.
//
// Solidity: event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)
func (_Root *RootFilterer) ParseNewHeaderBlock(log types.Log) (*RootNewHeaderBlock, error) {
	event := new(RootNewHeaderBlock)
	if err := _Root.contract.UnpackLog(event, "NewHeaderBlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Root contract.
type RootOwnershipTransferredIterator struct {
	Event *RootOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RootOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootOwnershipTransferred)
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
		it.Event = new(RootOwnershipTransferred)
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
func (it *RootOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootOwnershipTransferred represents a OwnershipTransferred event raised by the Root contract.
type RootOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Root *RootFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RootOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Root.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RootOwnershipTransferredIterator{contract: _Root.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Root *RootFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RootOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Root.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootOwnershipTransferred)
				if err := _Root.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Root *RootFilterer) ParseOwnershipTransferred(log types.Log) (*RootOwnershipTransferred, error) {
	event := new(RootOwnershipTransferred)
	if err := _Root.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RootResetHeaderBlockIterator is returned from FilterResetHeaderBlock and is used to iterate over the raw logs and unpacked data for ResetHeaderBlock events raised by the Root contract.
type RootResetHeaderBlockIterator struct {
	Event *RootResetHeaderBlock // Event containing the contract specifics and raw log

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
func (it *RootResetHeaderBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootResetHeaderBlock)
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
		it.Event = new(RootResetHeaderBlock)
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
func (it *RootResetHeaderBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootResetHeaderBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootResetHeaderBlock represents a ResetHeaderBlock event raised by the Root contract.
type RootResetHeaderBlock struct {
	Proposer      common.Address
	HeaderBlockId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterResetHeaderBlock is a free log retrieval operation binding the contract event 0xca1d8316287f938830e225956a7bb10fd5a1a1506dd2eb3a476751a488117205.
//
// Solidity: event ResetHeaderBlock(address indexed proposer, uint256 indexed headerBlockId)
func (_Root *RootFilterer) FilterResetHeaderBlock(opts *bind.FilterOpts, proposer []common.Address, headerBlockId []*big.Int) (*RootResetHeaderBlockIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}

	logs, sub, err := _Root.contract.FilterLogs(opts, "ResetHeaderBlock", proposerRule, headerBlockIdRule)
	if err != nil {
		return nil, err
	}
	return &RootResetHeaderBlockIterator{contract: _Root.contract, event: "ResetHeaderBlock", logs: logs, sub: sub}, nil
}

// WatchResetHeaderBlock is a free log subscription operation binding the contract event 0xca1d8316287f938830e225956a7bb10fd5a1a1506dd2eb3a476751a488117205.
//
// Solidity: event ResetHeaderBlock(address indexed proposer, uint256 indexed headerBlockId)
func (_Root *RootFilterer) WatchResetHeaderBlock(opts *bind.WatchOpts, sink chan<- *RootResetHeaderBlock, proposer []common.Address, headerBlockId []*big.Int) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var headerBlockIdRule []interface{}
	for _, headerBlockIdItem := range headerBlockId {
		headerBlockIdRule = append(headerBlockIdRule, headerBlockIdItem)
	}

	logs, sub, err := _Root.contract.WatchLogs(opts, "ResetHeaderBlock", proposerRule, headerBlockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootResetHeaderBlock)
				if err := _Root.contract.UnpackLog(event, "ResetHeaderBlock", log); err != nil {
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

// ParseResetHeaderBlock is a log parse operation binding the contract event 0xca1d8316287f938830e225956a7bb10fd5a1a1506dd2eb3a476751a488117205.
//
// Solidity: event ResetHeaderBlock(address indexed proposer, uint256 indexed headerBlockId)
func (_Root *RootFilterer) ParseResetHeaderBlock(log types.Log) (*RootResetHeaderBlock, error) {
	event := new(RootResetHeaderBlock)
	if err := _Root.contract.UnpackLog(event, "ResetHeaderBlock", log); err != nil {
		return nil, err
	}
	return event, nil
}
