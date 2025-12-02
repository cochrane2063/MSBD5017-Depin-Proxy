// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clearnetregistry

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

// ClearnetregistryMetaData contains all meta data concerning the Clearnetregistry contract.
var ClearnetregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clrToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidMetadataURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnedStake\",\"type\":\"uint256\"}],\"name\":\"NodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMetadataURI\",\"type\":\"string\"}],\"name\":\"NodeMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPricePerMinute\",\"type\":\"uint256\"}],\"name\":\"NodePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"}],\"name\":\"ReputationUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_PRICE_PER_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_PRICE_PER_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPUTATION_DECAY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodeIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clrToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeID\",\"type\":\"address\"}],\"name\":\"getNodeConnectionInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reputationScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerMinute\",\"type\":\"uint256\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newMetadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_newMetadataHash\",\"type\":\"bytes32\"}],\"name\":\"updateMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPricePerMinute\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeID\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newScore\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_successfulSession\",\"type\":\"bool\"}],\"name\":\"updateReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClearnetregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ClearnetregistryMetaData.ABI instead.
var ClearnetregistryABI = ClearnetregistryMetaData.ABI

// Clearnetregistry is an auto generated Go binding around an Ethereum contract.
type Clearnetregistry struct {
	ClearnetregistryCaller     // Read-only binding to the contract
	ClearnetregistryTransactor // Write-only binding to the contract
	ClearnetregistryFilterer   // Log filterer for contract events
}

// ClearnetregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClearnetregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearnetregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClearnetregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearnetregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClearnetregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearnetregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClearnetregistrySession struct {
	Contract     *Clearnetregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClearnetregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClearnetregistryCallerSession struct {
	Contract *ClearnetregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ClearnetregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClearnetregistryTransactorSession struct {
	Contract     *ClearnetregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ClearnetregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClearnetregistryRaw struct {
	Contract *Clearnetregistry // Generic contract binding to access the raw methods on
}

// ClearnetregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClearnetregistryCallerRaw struct {
	Contract *ClearnetregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ClearnetregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClearnetregistryTransactorRaw struct {
	Contract *ClearnetregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClearnetregistry creates a new instance of Clearnetregistry, bound to a specific deployed contract.
func NewClearnetregistry(address common.Address, backend bind.ContractBackend) (*Clearnetregistry, error) {
	contract, err := bindClearnetregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clearnetregistry{ClearnetregistryCaller: ClearnetregistryCaller{contract: contract}, ClearnetregistryTransactor: ClearnetregistryTransactor{contract: contract}, ClearnetregistryFilterer: ClearnetregistryFilterer{contract: contract}}, nil
}

// NewClearnetregistryCaller creates a new read-only instance of Clearnetregistry, bound to a specific deployed contract.
func NewClearnetregistryCaller(address common.Address, caller bind.ContractCaller) (*ClearnetregistryCaller, error) {
	contract, err := bindClearnetregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryCaller{contract: contract}, nil
}

// NewClearnetregistryTransactor creates a new write-only instance of Clearnetregistry, bound to a specific deployed contract.
func NewClearnetregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ClearnetregistryTransactor, error) {
	contract, err := bindClearnetregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryTransactor{contract: contract}, nil
}

// NewClearnetregistryFilterer creates a new log filterer instance of Clearnetregistry, bound to a specific deployed contract.
func NewClearnetregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ClearnetregistryFilterer, error) {
	contract, err := bindClearnetregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryFilterer{contract: contract}, nil
}

// bindClearnetregistry binds a generic wrapper to an already deployed contract.
func bindClearnetregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClearnetregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clearnetregistry *ClearnetregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clearnetregistry.Contract.ClearnetregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clearnetregistry *ClearnetregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.ClearnetregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clearnetregistry *ClearnetregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.ClearnetregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clearnetregistry *ClearnetregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clearnetregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clearnetregistry *ClearnetregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clearnetregistry *ClearnetregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.contract.Transact(opts, method, params...)
}

// MAXPRICEPERMIN is a free data retrieval call binding the contract method 0xf2e8937d.
//
// Solidity: function MAX_PRICE_PER_MIN() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCaller) MAXPRICEPERMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "MAX_PRICE_PER_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICEPERMIN is a free data retrieval call binding the contract method 0xf2e8937d.
//
// Solidity: function MAX_PRICE_PER_MIN() view returns(uint256)
func (_Clearnetregistry *ClearnetregistrySession) MAXPRICEPERMIN() (*big.Int, error) {
	return _Clearnetregistry.Contract.MAXPRICEPERMIN(&_Clearnetregistry.CallOpts)
}

// MAXPRICEPERMIN is a free data retrieval call binding the contract method 0xf2e8937d.
//
// Solidity: function MAX_PRICE_PER_MIN() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCallerSession) MAXPRICEPERMIN() (*big.Int, error) {
	return _Clearnetregistry.Contract.MAXPRICEPERMIN(&_Clearnetregistry.CallOpts)
}

// MINPRICEPERMIN is a free data retrieval call binding the contract method 0xcdf1c563.
//
// Solidity: function MIN_PRICE_PER_MIN() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCaller) MINPRICEPERMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "MIN_PRICE_PER_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPRICEPERMIN is a free data retrieval call binding the contract method 0xcdf1c563.
//
// Solidity: function MIN_PRICE_PER_MIN() view returns(uint256)
func (_Clearnetregistry *ClearnetregistrySession) MINPRICEPERMIN() (*big.Int, error) {
	return _Clearnetregistry.Contract.MINPRICEPERMIN(&_Clearnetregistry.CallOpts)
}

// MINPRICEPERMIN is a free data retrieval call binding the contract method 0xcdf1c563.
//
// Solidity: function MIN_PRICE_PER_MIN() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCallerSession) MINPRICEPERMIN() (*big.Int, error) {
	return _Clearnetregistry.Contract.MINPRICEPERMIN(&_Clearnetregistry.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_Clearnetregistry *ClearnetregistrySession) MINSTAKE() (*big.Int, error) {
	return _Clearnetregistry.Contract.MINSTAKE(&_Clearnetregistry.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCallerSession) MINSTAKE() (*big.Int, error) {
	return _Clearnetregistry.Contract.MINSTAKE(&_Clearnetregistry.CallOpts)
}

// REPUTATIONDECAYPERIOD is a free data retrieval call binding the contract method 0x8257886d.
//
// Solidity: function REPUTATION_DECAY_PERIOD() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCaller) REPUTATIONDECAYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "REPUTATION_DECAY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REPUTATIONDECAYPERIOD is a free data retrieval call binding the contract method 0x8257886d.
//
// Solidity: function REPUTATION_DECAY_PERIOD() view returns(uint256)
func (_Clearnetregistry *ClearnetregistrySession) REPUTATIONDECAYPERIOD() (*big.Int, error) {
	return _Clearnetregistry.Contract.REPUTATIONDECAYPERIOD(&_Clearnetregistry.CallOpts)
}

// REPUTATIONDECAYPERIOD is a free data retrieval call binding the contract method 0x8257886d.
//
// Solidity: function REPUTATION_DECAY_PERIOD() view returns(uint256)
func (_Clearnetregistry *ClearnetregistryCallerSession) REPUTATIONDECAYPERIOD() (*big.Int, error) {
	return _Clearnetregistry.Contract.REPUTATIONDECAYPERIOD(&_Clearnetregistry.CallOpts)
}

// ActiveNodeIds is a free data retrieval call binding the contract method 0xea17efdd.
//
// Solidity: function activeNodeIds(uint256 ) view returns(address)
func (_Clearnetregistry *ClearnetregistryCaller) ActiveNodeIds(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "activeNodeIds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveNodeIds is a free data retrieval call binding the contract method 0xea17efdd.
//
// Solidity: function activeNodeIds(uint256 ) view returns(address)
func (_Clearnetregistry *ClearnetregistrySession) ActiveNodeIds(arg0 *big.Int) (common.Address, error) {
	return _Clearnetregistry.Contract.ActiveNodeIds(&_Clearnetregistry.CallOpts, arg0)
}

// ActiveNodeIds is a free data retrieval call binding the contract method 0xea17efdd.
//
// Solidity: function activeNodeIds(uint256 ) view returns(address)
func (_Clearnetregistry *ClearnetregistryCallerSession) ActiveNodeIds(arg0 *big.Int) (common.Address, error) {
	return _Clearnetregistry.Contract.ActiveNodeIds(&_Clearnetregistry.CallOpts, arg0)
}

// ClrToken is a free data retrieval call binding the contract method 0x0329742e.
//
// Solidity: function clrToken() view returns(address)
func (_Clearnetregistry *ClearnetregistryCaller) ClrToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "clrToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClrToken is a free data retrieval call binding the contract method 0x0329742e.
//
// Solidity: function clrToken() view returns(address)
func (_Clearnetregistry *ClearnetregistrySession) ClrToken() (common.Address, error) {
	return _Clearnetregistry.Contract.ClrToken(&_Clearnetregistry.CallOpts)
}

// ClrToken is a free data retrieval call binding the contract method 0x0329742e.
//
// Solidity: function clrToken() view returns(address)
func (_Clearnetregistry *ClearnetregistryCallerSession) ClrToken() (common.Address, error) {
	return _Clearnetregistry.Contract.ClrToken(&_Clearnetregistry.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_Clearnetregistry *ClearnetregistryCaller) GetActiveNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "getActiveNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_Clearnetregistry *ClearnetregistrySession) GetActiveNodes() ([]common.Address, error) {
	return _Clearnetregistry.Contract.GetActiveNodes(&_Clearnetregistry.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_Clearnetregistry *ClearnetregistryCallerSession) GetActiveNodes() ([]common.Address, error) {
	return _Clearnetregistry.Contract.GetActiveNodes(&_Clearnetregistry.CallOpts)
}

// GetNodeConnectionInfo is a free data retrieval call binding the contract method 0xd3c49218.
//
// Solidity: function getNodeConnectionInfo(address _nodeID) view returns(string, uint256, uint256)
func (_Clearnetregistry *ClearnetregistryCaller) GetNodeConnectionInfo(opts *bind.CallOpts, _nodeID common.Address) (string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "getNodeConnectionInfo", _nodeID)

	if err != nil {
		return *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetNodeConnectionInfo is a free data retrieval call binding the contract method 0xd3c49218.
//
// Solidity: function getNodeConnectionInfo(address _nodeID) view returns(string, uint256, uint256)
func (_Clearnetregistry *ClearnetregistrySession) GetNodeConnectionInfo(_nodeID common.Address) (string, *big.Int, *big.Int, error) {
	return _Clearnetregistry.Contract.GetNodeConnectionInfo(&_Clearnetregistry.CallOpts, _nodeID)
}

// GetNodeConnectionInfo is a free data retrieval call binding the contract method 0xd3c49218.
//
// Solidity: function getNodeConnectionInfo(address _nodeID) view returns(string, uint256, uint256)
func (_Clearnetregistry *ClearnetregistryCallerSession) GetNodeConnectionInfo(_nodeID common.Address) (string, *big.Int, *big.Int, error) {
	return _Clearnetregistry.Contract.GetNodeConnectionInfo(&_Clearnetregistry.CallOpts, _nodeID)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(address nodeID, string metadataURI, bytes32 metadataHash, uint256 stakeAmount, uint256 reputationScore, uint256 pricePerMinute, bool isActive)
func (_Clearnetregistry *ClearnetregistryCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeID          common.Address
	MetadataURI     string
	MetadataHash    [32]byte
	StakeAmount     *big.Int
	ReputationScore *big.Int
	PricePerMinute  *big.Int
	IsActive        bool
}, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		NodeID          common.Address
		MetadataURI     string
		MetadataHash    [32]byte
		StakeAmount     *big.Int
		ReputationScore *big.Int
		PricePerMinute  *big.Int
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeID = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MetadataURI = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MetadataHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StakeAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReputationScore = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PricePerMinute = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(address nodeID, string metadataURI, bytes32 metadataHash, uint256 stakeAmount, uint256 reputationScore, uint256 pricePerMinute, bool isActive)
func (_Clearnetregistry *ClearnetregistrySession) Nodes(arg0 common.Address) (struct {
	NodeID          common.Address
	MetadataURI     string
	MetadataHash    [32]byte
	StakeAmount     *big.Int
	ReputationScore *big.Int
	PricePerMinute  *big.Int
	IsActive        bool
}, error) {
	return _Clearnetregistry.Contract.Nodes(&_Clearnetregistry.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(address nodeID, string metadataURI, bytes32 metadataHash, uint256 stakeAmount, uint256 reputationScore, uint256 pricePerMinute, bool isActive)
func (_Clearnetregistry *ClearnetregistryCallerSession) Nodes(arg0 common.Address) (struct {
	NodeID          common.Address
	MetadataURI     string
	MetadataHash    [32]byte
	StakeAmount     *big.Int
	ReputationScore *big.Int
	PricePerMinute  *big.Int
	IsActive        bool
}, error) {
	return _Clearnetregistry.Contract.Nodes(&_Clearnetregistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clearnetregistry *ClearnetregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clearnetregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clearnetregistry *ClearnetregistrySession) Owner() (common.Address, error) {
	return _Clearnetregistry.Contract.Owner(&_Clearnetregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clearnetregistry *ClearnetregistryCallerSession) Owner() (common.Address, error) {
	return _Clearnetregistry.Contract.Owner(&_Clearnetregistry.CallOpts)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_Clearnetregistry *ClearnetregistryTransactor) DeregisterNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "deregisterNode")
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_Clearnetregistry *ClearnetregistrySession) DeregisterNode() (*types.Transaction, error) {
	return _Clearnetregistry.Contract.DeregisterNode(&_Clearnetregistry.TransactOpts)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) DeregisterNode() (*types.Transaction, error) {
	return _Clearnetregistry.Contract.DeregisterNode(&_Clearnetregistry.TransactOpts)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xfa1bfe3d.
//
// Solidity: function registerNode(string _metadataURI, bytes32 _metadataHash, uint256 _pricePerMinute) returns()
func (_Clearnetregistry *ClearnetregistryTransactor) RegisterNode(opts *bind.TransactOpts, _metadataURI string, _metadataHash [32]byte, _pricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "registerNode", _metadataURI, _metadataHash, _pricePerMinute)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xfa1bfe3d.
//
// Solidity: function registerNode(string _metadataURI, bytes32 _metadataHash, uint256 _pricePerMinute) returns()
func (_Clearnetregistry *ClearnetregistrySession) RegisterNode(_metadataURI string, _metadataHash [32]byte, _pricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.RegisterNode(&_Clearnetregistry.TransactOpts, _metadataURI, _metadataHash, _pricePerMinute)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xfa1bfe3d.
//
// Solidity: function registerNode(string _metadataURI, bytes32 _metadataHash, uint256 _pricePerMinute) returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) RegisterNode(_metadataURI string, _metadataHash [32]byte, _pricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.RegisterNode(&_Clearnetregistry.TransactOpts, _metadataURI, _metadataHash, _pricePerMinute)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clearnetregistry *ClearnetregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clearnetregistry *ClearnetregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Clearnetregistry.Contract.RenounceOwnership(&_Clearnetregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clearnetregistry.Contract.RenounceOwnership(&_Clearnetregistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clearnetregistry *ClearnetregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clearnetregistry *ClearnetregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.TransferOwnership(&_Clearnetregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.TransferOwnership(&_Clearnetregistry.TransactOpts, newOwner)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x278bab3e.
//
// Solidity: function updateMetadata(string _newMetadataURI, bytes32 _newMetadataHash) returns()
func (_Clearnetregistry *ClearnetregistryTransactor) UpdateMetadata(opts *bind.TransactOpts, _newMetadataURI string, _newMetadataHash [32]byte) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "updateMetadata", _newMetadataURI, _newMetadataHash)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x278bab3e.
//
// Solidity: function updateMetadata(string _newMetadataURI, bytes32 _newMetadataHash) returns()
func (_Clearnetregistry *ClearnetregistrySession) UpdateMetadata(_newMetadataURI string, _newMetadataHash [32]byte) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.UpdateMetadata(&_Clearnetregistry.TransactOpts, _newMetadataURI, _newMetadataHash)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x278bab3e.
//
// Solidity: function updateMetadata(string _newMetadataURI, bytes32 _newMetadataHash) returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) UpdateMetadata(_newMetadataURI string, _newMetadataHash [32]byte) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.UpdateMetadata(&_Clearnetregistry.TransactOpts, _newMetadataURI, _newMetadataHash)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPricePerMinute) returns()
func (_Clearnetregistry *ClearnetregistryTransactor) UpdatePrice(opts *bind.TransactOpts, _newPricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "updatePrice", _newPricePerMinute)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPricePerMinute) returns()
func (_Clearnetregistry *ClearnetregistrySession) UpdatePrice(_newPricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.UpdatePrice(&_Clearnetregistry.TransactOpts, _newPricePerMinute)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPricePerMinute) returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) UpdatePrice(_newPricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.UpdatePrice(&_Clearnetregistry.TransactOpts, _newPricePerMinute)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xec604e46.
//
// Solidity: function updateReputation(address _nodeID, uint256 _newScore, bool _successfulSession) returns()
func (_Clearnetregistry *ClearnetregistryTransactor) UpdateReputation(opts *bind.TransactOpts, _nodeID common.Address, _newScore *big.Int, _successfulSession bool) (*types.Transaction, error) {
	return _Clearnetregistry.contract.Transact(opts, "updateReputation", _nodeID, _newScore, _successfulSession)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xec604e46.
//
// Solidity: function updateReputation(address _nodeID, uint256 _newScore, bool _successfulSession) returns()
func (_Clearnetregistry *ClearnetregistrySession) UpdateReputation(_nodeID common.Address, _newScore *big.Int, _successfulSession bool) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.UpdateReputation(&_Clearnetregistry.TransactOpts, _nodeID, _newScore, _successfulSession)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xec604e46.
//
// Solidity: function updateReputation(address _nodeID, uint256 _newScore, bool _successfulSession) returns()
func (_Clearnetregistry *ClearnetregistryTransactorSession) UpdateReputation(_nodeID common.Address, _newScore *big.Int, _successfulSession bool) (*types.Transaction, error) {
	return _Clearnetregistry.Contract.UpdateReputation(&_Clearnetregistry.TransactOpts, _nodeID, _newScore, _successfulSession)
}

// ClearnetregistryNodeDeregisteredIterator is returned from FilterNodeDeregistered and is used to iterate over the raw logs and unpacked data for NodeDeregistered events raised by the Clearnetregistry contract.
type ClearnetregistryNodeDeregisteredIterator struct {
	Event *ClearnetregistryNodeDeregistered // Event containing the contract specifics and raw log

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
func (it *ClearnetregistryNodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetregistryNodeDeregistered)
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
		it.Event = new(ClearnetregistryNodeDeregistered)
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
func (it *ClearnetregistryNodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetregistryNodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetregistryNodeDeregistered represents a NodeDeregistered event raised by the Clearnetregistry contract.
type ClearnetregistryNodeDeregistered struct {
	NodeID        common.Address
	ReturnedStake *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeDeregistered is a free log retrieval operation binding the contract event 0xaa433836fcb47675047002afc96e88cdbb2ba277f2508b8719e646fcd9b116f1.
//
// Solidity: event NodeDeregistered(address indexed nodeID, uint256 returnedStake)
func (_Clearnetregistry *ClearnetregistryFilterer) FilterNodeDeregistered(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetregistryNodeDeregisteredIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.FilterLogs(opts, "NodeDeregistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryNodeDeregisteredIterator{contract: _Clearnetregistry.contract, event: "NodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchNodeDeregistered is a free log subscription operation binding the contract event 0xaa433836fcb47675047002afc96e88cdbb2ba277f2508b8719e646fcd9b116f1.
//
// Solidity: event NodeDeregistered(address indexed nodeID, uint256 returnedStake)
func (_Clearnetregistry *ClearnetregistryFilterer) WatchNodeDeregistered(opts *bind.WatchOpts, sink chan<- *ClearnetregistryNodeDeregistered, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.WatchLogs(opts, "NodeDeregistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetregistryNodeDeregistered)
				if err := _Clearnetregistry.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
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

// ParseNodeDeregistered is a log parse operation binding the contract event 0xaa433836fcb47675047002afc96e88cdbb2ba277f2508b8719e646fcd9b116f1.
//
// Solidity: event NodeDeregistered(address indexed nodeID, uint256 returnedStake)
func (_Clearnetregistry *ClearnetregistryFilterer) ParseNodeDeregistered(log types.Log) (*ClearnetregistryNodeDeregistered, error) {
	event := new(ClearnetregistryNodeDeregistered)
	if err := _Clearnetregistry.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetregistryNodeMetadataUpdatedIterator is returned from FilterNodeMetadataUpdated and is used to iterate over the raw logs and unpacked data for NodeMetadataUpdated events raised by the Clearnetregistry contract.
type ClearnetregistryNodeMetadataUpdatedIterator struct {
	Event *ClearnetregistryNodeMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *ClearnetregistryNodeMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetregistryNodeMetadataUpdated)
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
		it.Event = new(ClearnetregistryNodeMetadataUpdated)
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
func (it *ClearnetregistryNodeMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetregistryNodeMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetregistryNodeMetadataUpdated represents a NodeMetadataUpdated event raised by the Clearnetregistry contract.
type ClearnetregistryNodeMetadataUpdated struct {
	NodeID         common.Address
	NewMetadataURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeMetadataUpdated is a free log retrieval operation binding the contract event 0xb60bc9c721160d82f8d49c7831e3690e535c3baee7a94314f36b93bb2d9dbfd1.
//
// Solidity: event NodeMetadataUpdated(address indexed nodeID, string newMetadataURI)
func (_Clearnetregistry *ClearnetregistryFilterer) FilterNodeMetadataUpdated(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetregistryNodeMetadataUpdatedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.FilterLogs(opts, "NodeMetadataUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryNodeMetadataUpdatedIterator{contract: _Clearnetregistry.contract, event: "NodeMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeMetadataUpdated is a free log subscription operation binding the contract event 0xb60bc9c721160d82f8d49c7831e3690e535c3baee7a94314f36b93bb2d9dbfd1.
//
// Solidity: event NodeMetadataUpdated(address indexed nodeID, string newMetadataURI)
func (_Clearnetregistry *ClearnetregistryFilterer) WatchNodeMetadataUpdated(opts *bind.WatchOpts, sink chan<- *ClearnetregistryNodeMetadataUpdated, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.WatchLogs(opts, "NodeMetadataUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetregistryNodeMetadataUpdated)
				if err := _Clearnetregistry.contract.UnpackLog(event, "NodeMetadataUpdated", log); err != nil {
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

// ParseNodeMetadataUpdated is a log parse operation binding the contract event 0xb60bc9c721160d82f8d49c7831e3690e535c3baee7a94314f36b93bb2d9dbfd1.
//
// Solidity: event NodeMetadataUpdated(address indexed nodeID, string newMetadataURI)
func (_Clearnetregistry *ClearnetregistryFilterer) ParseNodeMetadataUpdated(log types.Log) (*ClearnetregistryNodeMetadataUpdated, error) {
	event := new(ClearnetregistryNodeMetadataUpdated)
	if err := _Clearnetregistry.contract.UnpackLog(event, "NodeMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetregistryNodePriceUpdatedIterator is returned from FilterNodePriceUpdated and is used to iterate over the raw logs and unpacked data for NodePriceUpdated events raised by the Clearnetregistry contract.
type ClearnetregistryNodePriceUpdatedIterator struct {
	Event *ClearnetregistryNodePriceUpdated // Event containing the contract specifics and raw log

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
func (it *ClearnetregistryNodePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetregistryNodePriceUpdated)
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
		it.Event = new(ClearnetregistryNodePriceUpdated)
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
func (it *ClearnetregistryNodePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetregistryNodePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetregistryNodePriceUpdated represents a NodePriceUpdated event raised by the Clearnetregistry contract.
type ClearnetregistryNodePriceUpdated struct {
	NodeID            common.Address
	NewPricePerMinute *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNodePriceUpdated is a free log retrieval operation binding the contract event 0xaa357beaa928ff528af3cf3853f5075732f925103274cf6652df31fe9b224b0a.
//
// Solidity: event NodePriceUpdated(address indexed nodeID, uint256 newPricePerMinute)
func (_Clearnetregistry *ClearnetregistryFilterer) FilterNodePriceUpdated(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetregistryNodePriceUpdatedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.FilterLogs(opts, "NodePriceUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryNodePriceUpdatedIterator{contract: _Clearnetregistry.contract, event: "NodePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchNodePriceUpdated is a free log subscription operation binding the contract event 0xaa357beaa928ff528af3cf3853f5075732f925103274cf6652df31fe9b224b0a.
//
// Solidity: event NodePriceUpdated(address indexed nodeID, uint256 newPricePerMinute)
func (_Clearnetregistry *ClearnetregistryFilterer) WatchNodePriceUpdated(opts *bind.WatchOpts, sink chan<- *ClearnetregistryNodePriceUpdated, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.WatchLogs(opts, "NodePriceUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetregistryNodePriceUpdated)
				if err := _Clearnetregistry.contract.UnpackLog(event, "NodePriceUpdated", log); err != nil {
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

// ParseNodePriceUpdated is a log parse operation binding the contract event 0xaa357beaa928ff528af3cf3853f5075732f925103274cf6652df31fe9b224b0a.
//
// Solidity: event NodePriceUpdated(address indexed nodeID, uint256 newPricePerMinute)
func (_Clearnetregistry *ClearnetregistryFilterer) ParseNodePriceUpdated(log types.Log) (*ClearnetregistryNodePriceUpdated, error) {
	event := new(ClearnetregistryNodePriceUpdated)
	if err := _Clearnetregistry.contract.UnpackLog(event, "NodePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetregistryNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the Clearnetregistry contract.
type ClearnetregistryNodeRegisteredIterator struct {
	Event *ClearnetregistryNodeRegistered // Event containing the contract specifics and raw log

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
func (it *ClearnetregistryNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetregistryNodeRegistered)
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
		it.Event = new(ClearnetregistryNodeRegistered)
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
func (it *ClearnetregistryNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetregistryNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetregistryNodeRegistered represents a NodeRegistered event raised by the Clearnetregistry contract.
type ClearnetregistryNodeRegistered struct {
	NodeID         common.Address
	MetadataURI    string
	PricePerMinute *big.Int
	StakeAmount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0xf848331a8f5de76ec7a3fa83294f97f5113b398a6898b2d62529a8d513c81673.
//
// Solidity: event NodeRegistered(address indexed nodeID, string metadataURI, uint256 pricePerMinute, uint256 stakeAmount)
func (_Clearnetregistry *ClearnetregistryFilterer) FilterNodeRegistered(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetregistryNodeRegisteredIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.FilterLogs(opts, "NodeRegistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryNodeRegisteredIterator{contract: _Clearnetregistry.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0xf848331a8f5de76ec7a3fa83294f97f5113b398a6898b2d62529a8d513c81673.
//
// Solidity: event NodeRegistered(address indexed nodeID, string metadataURI, uint256 pricePerMinute, uint256 stakeAmount)
func (_Clearnetregistry *ClearnetregistryFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *ClearnetregistryNodeRegistered, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.WatchLogs(opts, "NodeRegistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetregistryNodeRegistered)
				if err := _Clearnetregistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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

// ParseNodeRegistered is a log parse operation binding the contract event 0xf848331a8f5de76ec7a3fa83294f97f5113b398a6898b2d62529a8d513c81673.
//
// Solidity: event NodeRegistered(address indexed nodeID, string metadataURI, uint256 pricePerMinute, uint256 stakeAmount)
func (_Clearnetregistry *ClearnetregistryFilterer) ParseNodeRegistered(log types.Log) (*ClearnetregistryNodeRegistered, error) {
	event := new(ClearnetregistryNodeRegistered)
	if err := _Clearnetregistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Clearnetregistry contract.
type ClearnetregistryOwnershipTransferredIterator struct {
	Event *ClearnetregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClearnetregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetregistryOwnershipTransferred)
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
		it.Event = new(ClearnetregistryOwnershipTransferred)
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
func (it *ClearnetregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Clearnetregistry contract.
type ClearnetregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clearnetregistry *ClearnetregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClearnetregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clearnetregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryOwnershipTransferredIterator{contract: _Clearnetregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clearnetregistry *ClearnetregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClearnetregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clearnetregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetregistryOwnershipTransferred)
				if err := _Clearnetregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Clearnetregistry *ClearnetregistryFilterer) ParseOwnershipTransferred(log types.Log) (*ClearnetregistryOwnershipTransferred, error) {
	event := new(ClearnetregistryOwnershipTransferred)
	if err := _Clearnetregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetregistryReputationUpdatedIterator is returned from FilterReputationUpdated and is used to iterate over the raw logs and unpacked data for ReputationUpdated events raised by the Clearnetregistry contract.
type ClearnetregistryReputationUpdatedIterator struct {
	Event *ClearnetregistryReputationUpdated // Event containing the contract specifics and raw log

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
func (it *ClearnetregistryReputationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetregistryReputationUpdated)
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
		it.Event = new(ClearnetregistryReputationUpdated)
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
func (it *ClearnetregistryReputationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetregistryReputationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetregistryReputationUpdated represents a ReputationUpdated event raised by the Clearnetregistry contract.
type ClearnetregistryReputationUpdated struct {
	NodeID   common.Address
	NewScore *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReputationUpdated is a free log retrieval operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed nodeID, uint256 newScore)
func (_Clearnetregistry *ClearnetregistryFilterer) FilterReputationUpdated(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetregistryReputationUpdatedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.FilterLogs(opts, "ReputationUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetregistryReputationUpdatedIterator{contract: _Clearnetregistry.contract, event: "ReputationUpdated", logs: logs, sub: sub}, nil
}

// WatchReputationUpdated is a free log subscription operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed nodeID, uint256 newScore)
func (_Clearnetregistry *ClearnetregistryFilterer) WatchReputationUpdated(opts *bind.WatchOpts, sink chan<- *ClearnetregistryReputationUpdated, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnetregistry.contract.WatchLogs(opts, "ReputationUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetregistryReputationUpdated)
				if err := _Clearnetregistry.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
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

// ParseReputationUpdated is a log parse operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed nodeID, uint256 newScore)
func (_Clearnetregistry *ClearnetregistryFilterer) ParseReputationUpdated(log types.Log) (*ClearnetregistryReputationUpdated, error) {
	event := new(ClearnetregistryReputationUpdated)
	if err := _Clearnetregistry.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
