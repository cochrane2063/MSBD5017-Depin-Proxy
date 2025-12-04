// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clearnet

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

// ClearnetMetaData contains all meta data concerning the Clearnet contract.
var ClearnetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clrToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnedStake\",\"type\":\"uint256\"}],\"name\":\"NodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPricePerMinute\",\"type\":\"uint256\"}],\"name\":\"NodePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingStake\",\"type\":\"uint256\"}],\"name\":\"NodeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newIp\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPort\",\"type\":\"uint16\"}],\"name\":\"NodeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"PaymentChannelToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minutesUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayShare\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"}],\"name\":\"RelayOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"}],\"name\":\"RelayOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"}],\"name\":\"ReputationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_REPUTATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MINUTES_PER_PAYMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRICE_PER_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REPUTATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CHANNEL_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_PRICE_PER_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKE_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPUTATION_DECREMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPUTATION_INCREMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPUTATION_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREASURY_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernanceTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeChannelIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodeIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relay\",\"type\":\"address\"}],\"name\":\"addRelayOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeID\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minutesUsed\",\"type\":\"uint256\"}],\"name\":\"calculateCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closePaymentChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clrToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getActiveNodesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"nodes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActivePaymentChannels\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getActivePaymentChannelsPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"channels\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeID\",\"type\":\"address\"}],\"name\":\"getNodeInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"ipAddress\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"port\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reputationScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinutesServed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalEarnings\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"getPaymentChannelInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernance\",\"type\":\"address\"}],\"name\":\"initiateGovernanceTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isRelayOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"ipAddress\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"port\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reputationScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinutesServed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalEarnings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastActivity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"openPaymentChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentChannels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minutesUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_clientSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_nodeSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_relaySignature\",\"type\":\"bytes\"}],\"name\":\"processPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddress\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerMinute\",\"type\":\"uint256\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayOperators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relay\",\"type\":\"address\"}],\"name\":\"removeRelayOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeID\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpPaymentChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBandwidthMinutes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProtocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newIpAddress\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_newPort\",\"type\":\"uint16\"}],\"name\":\"updateNodeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPricePerMinute\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeID\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_successfulSession\",\"type\":\"bool\"}],\"name\":\"updateReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTreasuryFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClearnetABI is the input ABI used to generate the binding from.
// Deprecated: Use ClearnetMetaData.ABI instead.
var ClearnetABI = ClearnetMetaData.ABI

// Clearnet is an auto generated Go binding around an Ethereum contract.
type Clearnet struct {
	ClearnetCaller     // Read-only binding to the contract
	ClearnetTransactor // Write-only binding to the contract
	ClearnetFilterer   // Log filterer for contract events
}

// ClearnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClearnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClearnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClearnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClearnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClearnetSession struct {
	Contract     *Clearnet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClearnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClearnetCallerSession struct {
	Contract *ClearnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ClearnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClearnetTransactorSession struct {
	Contract     *ClearnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ClearnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClearnetRaw struct {
	Contract *Clearnet // Generic contract binding to access the raw methods on
}

// ClearnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClearnetCallerRaw struct {
	Contract *ClearnetCaller // Generic read-only contract binding to access the raw methods on
}

// ClearnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClearnetTransactorRaw struct {
	Contract *ClearnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClearnet creates a new instance of Clearnet, bound to a specific deployed contract.
func NewClearnet(address common.Address, backend bind.ContractBackend) (*Clearnet, error) {
	contract, err := bindClearnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clearnet{ClearnetCaller: ClearnetCaller{contract: contract}, ClearnetTransactor: ClearnetTransactor{contract: contract}, ClearnetFilterer: ClearnetFilterer{contract: contract}}, nil
}

// NewClearnetCaller creates a new read-only instance of Clearnet, bound to a specific deployed contract.
func NewClearnetCaller(address common.Address, caller bind.ContractCaller) (*ClearnetCaller, error) {
	contract, err := bindClearnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClearnetCaller{contract: contract}, nil
}

// NewClearnetTransactor creates a new write-only instance of Clearnet, bound to a specific deployed contract.
func NewClearnetTransactor(address common.Address, transactor bind.ContractTransactor) (*ClearnetTransactor, error) {
	contract, err := bindClearnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClearnetTransactor{contract: contract}, nil
}

// NewClearnetFilterer creates a new log filterer instance of Clearnet, bound to a specific deployed contract.
func NewClearnetFilterer(address common.Address, filterer bind.ContractFilterer) (*ClearnetFilterer, error) {
	contract, err := bindClearnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClearnetFilterer{contract: contract}, nil
}

// bindClearnet binds a generic wrapper to an already deployed contract.
func bindClearnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClearnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clearnet *ClearnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clearnet.Contract.ClearnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clearnet *ClearnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.Contract.ClearnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clearnet *ClearnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clearnet.Contract.ClearnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clearnet *ClearnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clearnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clearnet *ClearnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clearnet *ClearnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clearnet.Contract.contract.Transact(opts, method, params...)
}

// INITIALREPUTATION is a free data retrieval call binding the contract method 0x7b057258.
//
// Solidity: function INITIAL_REPUTATION() view returns(uint256)
func (_Clearnet *ClearnetCaller) INITIALREPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "INITIAL_REPUTATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALREPUTATION is a free data retrieval call binding the contract method 0x7b057258.
//
// Solidity: function INITIAL_REPUTATION() view returns(uint256)
func (_Clearnet *ClearnetSession) INITIALREPUTATION() (*big.Int, error) {
	return _Clearnet.Contract.INITIALREPUTATION(&_Clearnet.CallOpts)
}

// INITIALREPUTATION is a free data retrieval call binding the contract method 0x7b057258.
//
// Solidity: function INITIAL_REPUTATION() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) INITIALREPUTATION() (*big.Int, error) {
	return _Clearnet.Contract.INITIALREPUTATION(&_Clearnet.CallOpts)
}

// MAXMINUTESPERPAYMENT is a free data retrieval call binding the contract method 0x17557939.
//
// Solidity: function MAX_MINUTES_PER_PAYMENT() view returns(uint256)
func (_Clearnet *ClearnetCaller) MAXMINUTESPERPAYMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MAX_MINUTES_PER_PAYMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMINUTESPERPAYMENT is a free data retrieval call binding the contract method 0x17557939.
//
// Solidity: function MAX_MINUTES_PER_PAYMENT() view returns(uint256)
func (_Clearnet *ClearnetSession) MAXMINUTESPERPAYMENT() (*big.Int, error) {
	return _Clearnet.Contract.MAXMINUTESPERPAYMENT(&_Clearnet.CallOpts)
}

// MAXMINUTESPERPAYMENT is a free data retrieval call binding the contract method 0x17557939.
//
// Solidity: function MAX_MINUTES_PER_PAYMENT() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MAXMINUTESPERPAYMENT() (*big.Int, error) {
	return _Clearnet.Contract.MAXMINUTESPERPAYMENT(&_Clearnet.CallOpts)
}

// MAXPRICEPERMIN is a free data retrieval call binding the contract method 0xf2e8937d.
//
// Solidity: function MAX_PRICE_PER_MIN() view returns(uint256)
func (_Clearnet *ClearnetCaller) MAXPRICEPERMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MAX_PRICE_PER_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICEPERMIN is a free data retrieval call binding the contract method 0xf2e8937d.
//
// Solidity: function MAX_PRICE_PER_MIN() view returns(uint256)
func (_Clearnet *ClearnetSession) MAXPRICEPERMIN() (*big.Int, error) {
	return _Clearnet.Contract.MAXPRICEPERMIN(&_Clearnet.CallOpts)
}

// MAXPRICEPERMIN is a free data retrieval call binding the contract method 0xf2e8937d.
//
// Solidity: function MAX_PRICE_PER_MIN() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MAXPRICEPERMIN() (*big.Int, error) {
	return _Clearnet.Contract.MAXPRICEPERMIN(&_Clearnet.CallOpts)
}

// MAXREPUTATION is a free data retrieval call binding the contract method 0xd213c0f2.
//
// Solidity: function MAX_REPUTATION() view returns(uint256)
func (_Clearnet *ClearnetCaller) MAXREPUTATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MAX_REPUTATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREPUTATION is a free data retrieval call binding the contract method 0xd213c0f2.
//
// Solidity: function MAX_REPUTATION() view returns(uint256)
func (_Clearnet *ClearnetSession) MAXREPUTATION() (*big.Int, error) {
	return _Clearnet.Contract.MAXREPUTATION(&_Clearnet.CallOpts)
}

// MAXREPUTATION is a free data retrieval call binding the contract method 0xd213c0f2.
//
// Solidity: function MAX_REPUTATION() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MAXREPUTATION() (*big.Int, error) {
	return _Clearnet.Contract.MAXREPUTATION(&_Clearnet.CallOpts)
}

// MINCHANNELBALANCE is a free data retrieval call binding the contract method 0xa2e5255e.
//
// Solidity: function MIN_CHANNEL_BALANCE() view returns(uint256)
func (_Clearnet *ClearnetCaller) MINCHANNELBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MIN_CHANNEL_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCHANNELBALANCE is a free data retrieval call binding the contract method 0xa2e5255e.
//
// Solidity: function MIN_CHANNEL_BALANCE() view returns(uint256)
func (_Clearnet *ClearnetSession) MINCHANNELBALANCE() (*big.Int, error) {
	return _Clearnet.Contract.MINCHANNELBALANCE(&_Clearnet.CallOpts)
}

// MINCHANNELBALANCE is a free data retrieval call binding the contract method 0xa2e5255e.
//
// Solidity: function MIN_CHANNEL_BALANCE() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MINCHANNELBALANCE() (*big.Int, error) {
	return _Clearnet.Contract.MINCHANNELBALANCE(&_Clearnet.CallOpts)
}

// MINPRICEPERMIN is a free data retrieval call binding the contract method 0xcdf1c563.
//
// Solidity: function MIN_PRICE_PER_MIN() view returns(uint256)
func (_Clearnet *ClearnetCaller) MINPRICEPERMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MIN_PRICE_PER_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPRICEPERMIN is a free data retrieval call binding the contract method 0xcdf1c563.
//
// Solidity: function MIN_PRICE_PER_MIN() view returns(uint256)
func (_Clearnet *ClearnetSession) MINPRICEPERMIN() (*big.Int, error) {
	return _Clearnet.Contract.MINPRICEPERMIN(&_Clearnet.CallOpts)
}

// MINPRICEPERMIN is a free data retrieval call binding the contract method 0xcdf1c563.
//
// Solidity: function MIN_PRICE_PER_MIN() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MINPRICEPERMIN() (*big.Int, error) {
	return _Clearnet.Contract.MINPRICEPERMIN(&_Clearnet.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_Clearnet *ClearnetCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_Clearnet *ClearnetSession) MINSTAKE() (*big.Int, error) {
	return _Clearnet.Contract.MINSTAKE(&_Clearnet.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MINSTAKE() (*big.Int, error) {
	return _Clearnet.Contract.MINSTAKE(&_Clearnet.CallOpts)
}

// MINSTAKETHRESHOLD is a free data retrieval call binding the contract method 0xed19dcc0.
//
// Solidity: function MIN_STAKE_THRESHOLD() view returns(uint256)
func (_Clearnet *ClearnetCaller) MINSTAKETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "MIN_STAKE_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKETHRESHOLD is a free data retrieval call binding the contract method 0xed19dcc0.
//
// Solidity: function MIN_STAKE_THRESHOLD() view returns(uint256)
func (_Clearnet *ClearnetSession) MINSTAKETHRESHOLD() (*big.Int, error) {
	return _Clearnet.Contract.MINSTAKETHRESHOLD(&_Clearnet.CallOpts)
}

// MINSTAKETHRESHOLD is a free data retrieval call binding the contract method 0xed19dcc0.
//
// Solidity: function MIN_STAKE_THRESHOLD() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) MINSTAKETHRESHOLD() (*big.Int, error) {
	return _Clearnet.Contract.MINSTAKETHRESHOLD(&_Clearnet.CallOpts)
}

// NODESHARE is a free data retrieval call binding the contract method 0x2ffb61fb.
//
// Solidity: function NODE_SHARE() view returns(uint256)
func (_Clearnet *ClearnetCaller) NODESHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "NODE_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NODESHARE is a free data retrieval call binding the contract method 0x2ffb61fb.
//
// Solidity: function NODE_SHARE() view returns(uint256)
func (_Clearnet *ClearnetSession) NODESHARE() (*big.Int, error) {
	return _Clearnet.Contract.NODESHARE(&_Clearnet.CallOpts)
}

// NODESHARE is a free data retrieval call binding the contract method 0x2ffb61fb.
//
// Solidity: function NODE_SHARE() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) NODESHARE() (*big.Int, error) {
	return _Clearnet.Contract.NODESHARE(&_Clearnet.CallOpts)
}

// RELAYSHARE is a free data retrieval call binding the contract method 0xe0cb9935.
//
// Solidity: function RELAY_SHARE() view returns(uint256)
func (_Clearnet *ClearnetCaller) RELAYSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "RELAY_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELAYSHARE is a free data retrieval call binding the contract method 0xe0cb9935.
//
// Solidity: function RELAY_SHARE() view returns(uint256)
func (_Clearnet *ClearnetSession) RELAYSHARE() (*big.Int, error) {
	return _Clearnet.Contract.RELAYSHARE(&_Clearnet.CallOpts)
}

// RELAYSHARE is a free data retrieval call binding the contract method 0xe0cb9935.
//
// Solidity: function RELAY_SHARE() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) RELAYSHARE() (*big.Int, error) {
	return _Clearnet.Contract.RELAYSHARE(&_Clearnet.CallOpts)
}

// REPUTATIONDECREMENT is a free data retrieval call binding the contract method 0x2b16d75a.
//
// Solidity: function REPUTATION_DECREMENT() view returns(uint256)
func (_Clearnet *ClearnetCaller) REPUTATIONDECREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "REPUTATION_DECREMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REPUTATIONDECREMENT is a free data retrieval call binding the contract method 0x2b16d75a.
//
// Solidity: function REPUTATION_DECREMENT() view returns(uint256)
func (_Clearnet *ClearnetSession) REPUTATIONDECREMENT() (*big.Int, error) {
	return _Clearnet.Contract.REPUTATIONDECREMENT(&_Clearnet.CallOpts)
}

// REPUTATIONDECREMENT is a free data retrieval call binding the contract method 0x2b16d75a.
//
// Solidity: function REPUTATION_DECREMENT() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) REPUTATIONDECREMENT() (*big.Int, error) {
	return _Clearnet.Contract.REPUTATIONDECREMENT(&_Clearnet.CallOpts)
}

// REPUTATIONINCREMENT is a free data retrieval call binding the contract method 0x3acc361a.
//
// Solidity: function REPUTATION_INCREMENT() view returns(uint256)
func (_Clearnet *ClearnetCaller) REPUTATIONINCREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "REPUTATION_INCREMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REPUTATIONINCREMENT is a free data retrieval call binding the contract method 0x3acc361a.
//
// Solidity: function REPUTATION_INCREMENT() view returns(uint256)
func (_Clearnet *ClearnetSession) REPUTATIONINCREMENT() (*big.Int, error) {
	return _Clearnet.Contract.REPUTATIONINCREMENT(&_Clearnet.CallOpts)
}

// REPUTATIONINCREMENT is a free data retrieval call binding the contract method 0x3acc361a.
//
// Solidity: function REPUTATION_INCREMENT() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) REPUTATIONINCREMENT() (*big.Int, error) {
	return _Clearnet.Contract.REPUTATIONINCREMENT(&_Clearnet.CallOpts)
}

// REPUTATIONPRECISION is a free data retrieval call binding the contract method 0xbae2e378.
//
// Solidity: function REPUTATION_PRECISION() view returns(uint256)
func (_Clearnet *ClearnetCaller) REPUTATIONPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "REPUTATION_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REPUTATIONPRECISION is a free data retrieval call binding the contract method 0xbae2e378.
//
// Solidity: function REPUTATION_PRECISION() view returns(uint256)
func (_Clearnet *ClearnetSession) REPUTATIONPRECISION() (*big.Int, error) {
	return _Clearnet.Contract.REPUTATIONPRECISION(&_Clearnet.CallOpts)
}

// REPUTATIONPRECISION is a free data retrieval call binding the contract method 0xbae2e378.
//
// Solidity: function REPUTATION_PRECISION() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) REPUTATIONPRECISION() (*big.Int, error) {
	return _Clearnet.Contract.REPUTATIONPRECISION(&_Clearnet.CallOpts)
}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_Clearnet *ClearnetCaller) SHAREDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "SHARE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_Clearnet *ClearnetSession) SHAREDENOMINATOR() (*big.Int, error) {
	return _Clearnet.Contract.SHAREDENOMINATOR(&_Clearnet.CallOpts)
}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) SHAREDENOMINATOR() (*big.Int, error) {
	return _Clearnet.Contract.SHAREDENOMINATOR(&_Clearnet.CallOpts)
}

// TREASURYSHARE is a free data retrieval call binding the contract method 0x32696174.
//
// Solidity: function TREASURY_SHARE() view returns(uint256)
func (_Clearnet *ClearnetCaller) TREASURYSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "TREASURY_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TREASURYSHARE is a free data retrieval call binding the contract method 0x32696174.
//
// Solidity: function TREASURY_SHARE() view returns(uint256)
func (_Clearnet *ClearnetSession) TREASURYSHARE() (*big.Int, error) {
	return _Clearnet.Contract.TREASURYSHARE(&_Clearnet.CallOpts)
}

// TREASURYSHARE is a free data retrieval call binding the contract method 0x32696174.
//
// Solidity: function TREASURY_SHARE() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) TREASURYSHARE() (*big.Int, error) {
	return _Clearnet.Contract.TREASURYSHARE(&_Clearnet.CallOpts)
}

// ActiveChannelIds is a free data retrieval call binding the contract method 0x29ca68ee.
//
// Solidity: function activeChannelIds(uint256 ) view returns(address)
func (_Clearnet *ClearnetCaller) ActiveChannelIds(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "activeChannelIds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveChannelIds is a free data retrieval call binding the contract method 0x29ca68ee.
//
// Solidity: function activeChannelIds(uint256 ) view returns(address)
func (_Clearnet *ClearnetSession) ActiveChannelIds(arg0 *big.Int) (common.Address, error) {
	return _Clearnet.Contract.ActiveChannelIds(&_Clearnet.CallOpts, arg0)
}

// ActiveChannelIds is a free data retrieval call binding the contract method 0x29ca68ee.
//
// Solidity: function activeChannelIds(uint256 ) view returns(address)
func (_Clearnet *ClearnetCallerSession) ActiveChannelIds(arg0 *big.Int) (common.Address, error) {
	return _Clearnet.Contract.ActiveChannelIds(&_Clearnet.CallOpts, arg0)
}

// ActiveNodeIds is a free data retrieval call binding the contract method 0xea17efdd.
//
// Solidity: function activeNodeIds(uint256 ) view returns(address)
func (_Clearnet *ClearnetCaller) ActiveNodeIds(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "activeNodeIds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveNodeIds is a free data retrieval call binding the contract method 0xea17efdd.
//
// Solidity: function activeNodeIds(uint256 ) view returns(address)
func (_Clearnet *ClearnetSession) ActiveNodeIds(arg0 *big.Int) (common.Address, error) {
	return _Clearnet.Contract.ActiveNodeIds(&_Clearnet.CallOpts, arg0)
}

// ActiveNodeIds is a free data retrieval call binding the contract method 0xea17efdd.
//
// Solidity: function activeNodeIds(uint256 ) view returns(address)
func (_Clearnet *ClearnetCallerSession) ActiveNodeIds(arg0 *big.Int) (common.Address, error) {
	return _Clearnet.Contract.ActiveNodeIds(&_Clearnet.CallOpts, arg0)
}

// CalculateCost is a free data retrieval call binding the contract method 0x22d34454.
//
// Solidity: function calculateCost(address _nodeID, uint256 _minutesUsed) view returns(uint256)
func (_Clearnet *ClearnetCaller) CalculateCost(opts *bind.CallOpts, _nodeID common.Address, _minutesUsed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "calculateCost", _nodeID, _minutesUsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCost is a free data retrieval call binding the contract method 0x22d34454.
//
// Solidity: function calculateCost(address _nodeID, uint256 _minutesUsed) view returns(uint256)
func (_Clearnet *ClearnetSession) CalculateCost(_nodeID common.Address, _minutesUsed *big.Int) (*big.Int, error) {
	return _Clearnet.Contract.CalculateCost(&_Clearnet.CallOpts, _nodeID, _minutesUsed)
}

// CalculateCost is a free data retrieval call binding the contract method 0x22d34454.
//
// Solidity: function calculateCost(address _nodeID, uint256 _minutesUsed) view returns(uint256)
func (_Clearnet *ClearnetCallerSession) CalculateCost(_nodeID common.Address, _minutesUsed *big.Int) (*big.Int, error) {
	return _Clearnet.Contract.CalculateCost(&_Clearnet.CallOpts, _nodeID, _minutesUsed)
}

// ClrToken is a free data retrieval call binding the contract method 0x0329742e.
//
// Solidity: function clrToken() view returns(address)
func (_Clearnet *ClearnetCaller) ClrToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "clrToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClrToken is a free data retrieval call binding the contract method 0x0329742e.
//
// Solidity: function clrToken() view returns(address)
func (_Clearnet *ClearnetSession) ClrToken() (common.Address, error) {
	return _Clearnet.Contract.ClrToken(&_Clearnet.CallOpts)
}

// ClrToken is a free data retrieval call binding the contract method 0x0329742e.
//
// Solidity: function clrToken() view returns(address)
func (_Clearnet *ClearnetCallerSession) ClrToken() (common.Address, error) {
	return _Clearnet.Contract.ClrToken(&_Clearnet.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_Clearnet *ClearnetCaller) GetActiveNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "getActiveNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_Clearnet *ClearnetSession) GetActiveNodes() ([]common.Address, error) {
	return _Clearnet.Contract.GetActiveNodes(&_Clearnet.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_Clearnet *ClearnetCallerSession) GetActiveNodes() ([]common.Address, error) {
	return _Clearnet.Contract.GetActiveNodes(&_Clearnet.CallOpts)
}

// GetActiveNodesPaginated is a free data retrieval call binding the contract method 0x04f4d1b6.
//
// Solidity: function getActiveNodesPaginated(uint256 _offset, uint256 _limit) view returns(address[] nodes, uint256 total)
func (_Clearnet *ClearnetCaller) GetActiveNodesPaginated(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) (struct {
	Nodes []common.Address
	Total *big.Int
}, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "getActiveNodesPaginated", _offset, _limit)

	outstruct := new(struct {
		Nodes []common.Address
		Total *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nodes = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetActiveNodesPaginated is a free data retrieval call binding the contract method 0x04f4d1b6.
//
// Solidity: function getActiveNodesPaginated(uint256 _offset, uint256 _limit) view returns(address[] nodes, uint256 total)
func (_Clearnet *ClearnetSession) GetActiveNodesPaginated(_offset *big.Int, _limit *big.Int) (struct {
	Nodes []common.Address
	Total *big.Int
}, error) {
	return _Clearnet.Contract.GetActiveNodesPaginated(&_Clearnet.CallOpts, _offset, _limit)
}

// GetActiveNodesPaginated is a free data retrieval call binding the contract method 0x04f4d1b6.
//
// Solidity: function getActiveNodesPaginated(uint256 _offset, uint256 _limit) view returns(address[] nodes, uint256 total)
func (_Clearnet *ClearnetCallerSession) GetActiveNodesPaginated(_offset *big.Int, _limit *big.Int) (struct {
	Nodes []common.Address
	Total *big.Int
}, error) {
	return _Clearnet.Contract.GetActiveNodesPaginated(&_Clearnet.CallOpts, _offset, _limit)
}

// GetActivePaymentChannels is a free data retrieval call binding the contract method 0x79a1bad5.
//
// Solidity: function getActivePaymentChannels() view returns(address[])
func (_Clearnet *ClearnetCaller) GetActivePaymentChannels(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "getActivePaymentChannels")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActivePaymentChannels is a free data retrieval call binding the contract method 0x79a1bad5.
//
// Solidity: function getActivePaymentChannels() view returns(address[])
func (_Clearnet *ClearnetSession) GetActivePaymentChannels() ([]common.Address, error) {
	return _Clearnet.Contract.GetActivePaymentChannels(&_Clearnet.CallOpts)
}

// GetActivePaymentChannels is a free data retrieval call binding the contract method 0x79a1bad5.
//
// Solidity: function getActivePaymentChannels() view returns(address[])
func (_Clearnet *ClearnetCallerSession) GetActivePaymentChannels() ([]common.Address, error) {
	return _Clearnet.Contract.GetActivePaymentChannels(&_Clearnet.CallOpts)
}

// GetActivePaymentChannelsPaginated is a free data retrieval call binding the contract method 0x0f4ac374.
//
// Solidity: function getActivePaymentChannelsPaginated(uint256 _offset, uint256 _limit) view returns(address[] channels, uint256 total)
func (_Clearnet *ClearnetCaller) GetActivePaymentChannelsPaginated(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) (struct {
	Channels []common.Address
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "getActivePaymentChannelsPaginated", _offset, _limit)

	outstruct := new(struct {
		Channels []common.Address
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Channels = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetActivePaymentChannelsPaginated is a free data retrieval call binding the contract method 0x0f4ac374.
//
// Solidity: function getActivePaymentChannelsPaginated(uint256 _offset, uint256 _limit) view returns(address[] channels, uint256 total)
func (_Clearnet *ClearnetSession) GetActivePaymentChannelsPaginated(_offset *big.Int, _limit *big.Int) (struct {
	Channels []common.Address
	Total    *big.Int
}, error) {
	return _Clearnet.Contract.GetActivePaymentChannelsPaginated(&_Clearnet.CallOpts, _offset, _limit)
}

// GetActivePaymentChannelsPaginated is a free data retrieval call binding the contract method 0x0f4ac374.
//
// Solidity: function getActivePaymentChannelsPaginated(uint256 _offset, uint256 _limit) view returns(address[] channels, uint256 total)
func (_Clearnet *ClearnetCallerSession) GetActivePaymentChannelsPaginated(_offset *big.Int, _limit *big.Int) (struct {
	Channels []common.Address
	Total    *big.Int
}, error) {
	return _Clearnet.Contract.GetActivePaymentChannelsPaginated(&_Clearnet.CallOpts, _offset, _limit)
}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address _nodeID) view returns(string ipAddress, uint16 port, uint256 pricePerMinute, uint256 reputationScore, uint256 totalMinutesServed, uint256 totalEarnings)
func (_Clearnet *ClearnetCaller) GetNodeInfo(opts *bind.CallOpts, _nodeID common.Address) (struct {
	IpAddress          string
	Port               uint16
	PricePerMinute     *big.Int
	ReputationScore    *big.Int
	TotalMinutesServed *big.Int
	TotalEarnings      *big.Int
}, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "getNodeInfo", _nodeID)

	outstruct := new(struct {
		IpAddress          string
		Port               uint16
		PricePerMinute     *big.Int
		ReputationScore    *big.Int
		TotalMinutesServed *big.Int
		TotalEarnings      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IpAddress = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Port = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.PricePerMinute = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReputationScore = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalMinutesServed = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalEarnings = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address _nodeID) view returns(string ipAddress, uint16 port, uint256 pricePerMinute, uint256 reputationScore, uint256 totalMinutesServed, uint256 totalEarnings)
func (_Clearnet *ClearnetSession) GetNodeInfo(_nodeID common.Address) (struct {
	IpAddress          string
	Port               uint16
	PricePerMinute     *big.Int
	ReputationScore    *big.Int
	TotalMinutesServed *big.Int
	TotalEarnings      *big.Int
}, error) {
	return _Clearnet.Contract.GetNodeInfo(&_Clearnet.CallOpts, _nodeID)
}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address _nodeID) view returns(string ipAddress, uint16 port, uint256 pricePerMinute, uint256 reputationScore, uint256 totalMinutesServed, uint256 totalEarnings)
func (_Clearnet *ClearnetCallerSession) GetNodeInfo(_nodeID common.Address) (struct {
	IpAddress          string
	Port               uint16
	PricePerMinute     *big.Int
	ReputationScore    *big.Int
	TotalMinutesServed *big.Int
	TotalEarnings      *big.Int
}, error) {
	return _Clearnet.Contract.GetNodeInfo(&_Clearnet.CallOpts, _nodeID)
}

// GetPaymentChannelInfo is a free data retrieval call binding the contract method 0x47e9e803.
//
// Solidity: function getPaymentChannelInfo(address _client) view returns(uint256 balance, uint256 nonce, bool isActive)
func (_Clearnet *ClearnetCaller) GetPaymentChannelInfo(opts *bind.CallOpts, _client common.Address) (struct {
	Balance  *big.Int
	Nonce    *big.Int
	IsActive bool
}, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "getPaymentChannelInfo", _client)

	outstruct := new(struct {
		Balance  *big.Int
		Nonce    *big.Int
		IsActive bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetPaymentChannelInfo is a free data retrieval call binding the contract method 0x47e9e803.
//
// Solidity: function getPaymentChannelInfo(address _client) view returns(uint256 balance, uint256 nonce, bool isActive)
func (_Clearnet *ClearnetSession) GetPaymentChannelInfo(_client common.Address) (struct {
	Balance  *big.Int
	Nonce    *big.Int
	IsActive bool
}, error) {
	return _Clearnet.Contract.GetPaymentChannelInfo(&_Clearnet.CallOpts, _client)
}

// GetPaymentChannelInfo is a free data retrieval call binding the contract method 0x47e9e803.
//
// Solidity: function getPaymentChannelInfo(address _client) view returns(uint256 balance, uint256 nonce, bool isActive)
func (_Clearnet *ClearnetCallerSession) GetPaymentChannelInfo(_client common.Address) (struct {
	Balance  *big.Int
	Nonce    *big.Int
	IsActive bool
}, error) {
	return _Clearnet.Contract.GetPaymentChannelInfo(&_Clearnet.CallOpts, _client)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(address)
func (_Clearnet *ClearnetCaller) GovernanceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "governanceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(address)
func (_Clearnet *ClearnetSession) GovernanceContract() (common.Address, error) {
	return _Clearnet.Contract.GovernanceContract(&_Clearnet.CallOpts)
}

// GovernanceContract is a free data retrieval call binding the contract method 0xb172b222.
//
// Solidity: function governanceContract() view returns(address)
func (_Clearnet *ClearnetCallerSession) GovernanceContract() (common.Address, error) {
	return _Clearnet.Contract.GovernanceContract(&_Clearnet.CallOpts)
}

// IsRelayOperator is a free data retrieval call binding the contract method 0xc6da94c5.
//
// Solidity: function isRelayOperator(address _address) view returns(bool)
func (_Clearnet *ClearnetCaller) IsRelayOperator(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "isRelayOperator", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayOperator is a free data retrieval call binding the contract method 0xc6da94c5.
//
// Solidity: function isRelayOperator(address _address) view returns(bool)
func (_Clearnet *ClearnetSession) IsRelayOperator(_address common.Address) (bool, error) {
	return _Clearnet.Contract.IsRelayOperator(&_Clearnet.CallOpts, _address)
}

// IsRelayOperator is a free data retrieval call binding the contract method 0xc6da94c5.
//
// Solidity: function isRelayOperator(address _address) view returns(bool)
func (_Clearnet *ClearnetCallerSession) IsRelayOperator(_address common.Address) (bool, error) {
	return _Clearnet.Contract.IsRelayOperator(&_Clearnet.CallOpts, _address)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(string ipAddress, uint16 port, bool isActive, uint256 stakeAmount, uint256 reputationScore, uint256 pricePerMinute, uint256 totalMinutesServed, uint256 totalEarnings, uint256 lastActivity)
func (_Clearnet *ClearnetCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	IpAddress          string
	Port               uint16
	IsActive           bool
	StakeAmount        *big.Int
	ReputationScore    *big.Int
	PricePerMinute     *big.Int
	TotalMinutesServed *big.Int
	TotalEarnings      *big.Int
	LastActivity       *big.Int
}, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		IpAddress          string
		Port               uint16
		IsActive           bool
		StakeAmount        *big.Int
		ReputationScore    *big.Int
		PricePerMinute     *big.Int
		TotalMinutesServed *big.Int
		TotalEarnings      *big.Int
		LastActivity       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IpAddress = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Port = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.StakeAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReputationScore = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PricePerMinute = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalMinutesServed = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalEarnings = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LastActivity = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(string ipAddress, uint16 port, bool isActive, uint256 stakeAmount, uint256 reputationScore, uint256 pricePerMinute, uint256 totalMinutesServed, uint256 totalEarnings, uint256 lastActivity)
func (_Clearnet *ClearnetSession) Nodes(arg0 common.Address) (struct {
	IpAddress          string
	Port               uint16
	IsActive           bool
	StakeAmount        *big.Int
	ReputationScore    *big.Int
	PricePerMinute     *big.Int
	TotalMinutesServed *big.Int
	TotalEarnings      *big.Int
	LastActivity       *big.Int
}, error) {
	return _Clearnet.Contract.Nodes(&_Clearnet.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(string ipAddress, uint16 port, bool isActive, uint256 stakeAmount, uint256 reputationScore, uint256 pricePerMinute, uint256 totalMinutesServed, uint256 totalEarnings, uint256 lastActivity)
func (_Clearnet *ClearnetCallerSession) Nodes(arg0 common.Address) (struct {
	IpAddress          string
	Port               uint16
	IsActive           bool
	StakeAmount        *big.Int
	ReputationScore    *big.Int
	PricePerMinute     *big.Int
	TotalMinutesServed *big.Int
	TotalEarnings      *big.Int
	LastActivity       *big.Int
}, error) {
	return _Clearnet.Contract.Nodes(&_Clearnet.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clearnet *ClearnetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clearnet *ClearnetSession) Owner() (common.Address, error) {
	return _Clearnet.Contract.Owner(&_Clearnet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clearnet *ClearnetCallerSession) Owner() (common.Address, error) {
	return _Clearnet.Contract.Owner(&_Clearnet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Clearnet *ClearnetCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Clearnet *ClearnetSession) Paused() (bool, error) {
	return _Clearnet.Contract.Paused(&_Clearnet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Clearnet *ClearnetCallerSession) Paused() (bool, error) {
	return _Clearnet.Contract.Paused(&_Clearnet.CallOpts)
}

// PaymentChannels is a free data retrieval call binding the contract method 0x8639497c.
//
// Solidity: function paymentChannels(address ) view returns(uint256 balance, uint256 nonce, bool isActive)
func (_Clearnet *ClearnetCaller) PaymentChannels(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance  *big.Int
	Nonce    *big.Int
	IsActive bool
}, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "paymentChannels", arg0)

	outstruct := new(struct {
		Balance  *big.Int
		Nonce    *big.Int
		IsActive bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// PaymentChannels is a free data retrieval call binding the contract method 0x8639497c.
//
// Solidity: function paymentChannels(address ) view returns(uint256 balance, uint256 nonce, bool isActive)
func (_Clearnet *ClearnetSession) PaymentChannels(arg0 common.Address) (struct {
	Balance  *big.Int
	Nonce    *big.Int
	IsActive bool
}, error) {
	return _Clearnet.Contract.PaymentChannels(&_Clearnet.CallOpts, arg0)
}

// PaymentChannels is a free data retrieval call binding the contract method 0x8639497c.
//
// Solidity: function paymentChannels(address ) view returns(uint256 balance, uint256 nonce, bool isActive)
func (_Clearnet *ClearnetCallerSession) PaymentChannels(arg0 common.Address) (struct {
	Balance  *big.Int
	Nonce    *big.Int
	IsActive bool
}, error) {
	return _Clearnet.Contract.PaymentChannels(&_Clearnet.CallOpts, arg0)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Clearnet *ClearnetCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Clearnet *ClearnetSession) PendingGovernance() (common.Address, error) {
	return _Clearnet.Contract.PendingGovernance(&_Clearnet.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Clearnet *ClearnetCallerSession) PendingGovernance() (common.Address, error) {
	return _Clearnet.Contract.PendingGovernance(&_Clearnet.CallOpts)
}

// RelayOperators is a free data retrieval call binding the contract method 0x059b2725.
//
// Solidity: function relayOperators(address ) view returns(bool)
func (_Clearnet *ClearnetCaller) RelayOperators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "relayOperators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayOperators is a free data retrieval call binding the contract method 0x059b2725.
//
// Solidity: function relayOperators(address ) view returns(bool)
func (_Clearnet *ClearnetSession) RelayOperators(arg0 common.Address) (bool, error) {
	return _Clearnet.Contract.RelayOperators(&_Clearnet.CallOpts, arg0)
}

// RelayOperators is a free data retrieval call binding the contract method 0x059b2725.
//
// Solidity: function relayOperators(address ) view returns(bool)
func (_Clearnet *ClearnetCallerSession) RelayOperators(arg0 common.Address) (bool, error) {
	return _Clearnet.Contract.RelayOperators(&_Clearnet.CallOpts, arg0)
}

// TotalBandwidthMinutes is a free data retrieval call binding the contract method 0xed555a05.
//
// Solidity: function totalBandwidthMinutes() view returns(uint256)
func (_Clearnet *ClearnetCaller) TotalBandwidthMinutes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "totalBandwidthMinutes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBandwidthMinutes is a free data retrieval call binding the contract method 0xed555a05.
//
// Solidity: function totalBandwidthMinutes() view returns(uint256)
func (_Clearnet *ClearnetSession) TotalBandwidthMinutes() (*big.Int, error) {
	return _Clearnet.Contract.TotalBandwidthMinutes(&_Clearnet.CallOpts)
}

// TotalBandwidthMinutes is a free data retrieval call binding the contract method 0xed555a05.
//
// Solidity: function totalBandwidthMinutes() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) TotalBandwidthMinutes() (*big.Int, error) {
	return _Clearnet.Contract.TotalBandwidthMinutes(&_Clearnet.CallOpts)
}

// TotalProtocolFees is a free data retrieval call binding the contract method 0xf526e39a.
//
// Solidity: function totalProtocolFees() view returns(uint256)
func (_Clearnet *ClearnetCaller) TotalProtocolFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "totalProtocolFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProtocolFees is a free data retrieval call binding the contract method 0xf526e39a.
//
// Solidity: function totalProtocolFees() view returns(uint256)
func (_Clearnet *ClearnetSession) TotalProtocolFees() (*big.Int, error) {
	return _Clearnet.Contract.TotalProtocolFees(&_Clearnet.CallOpts)
}

// TotalProtocolFees is a free data retrieval call binding the contract method 0xf526e39a.
//
// Solidity: function totalProtocolFees() view returns(uint256)
func (_Clearnet *ClearnetCallerSession) TotalProtocolFees() (*big.Int, error) {
	return _Clearnet.Contract.TotalProtocolFees(&_Clearnet.CallOpts)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 ) view returns(bool)
func (_Clearnet *ClearnetCaller) UsedSignatures(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Clearnet.contract.Call(opts, &out, "usedSignatures", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 ) view returns(bool)
func (_Clearnet *ClearnetSession) UsedSignatures(arg0 [32]byte) (bool, error) {
	return _Clearnet.Contract.UsedSignatures(&_Clearnet.CallOpts, arg0)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 ) view returns(bool)
func (_Clearnet *ClearnetCallerSession) UsedSignatures(arg0 [32]byte) (bool, error) {
	return _Clearnet.Contract.UsedSignatures(&_Clearnet.CallOpts, arg0)
}

// AcceptGovernanceTransfer is a paid mutator transaction binding the contract method 0x2a90b8da.
//
// Solidity: function acceptGovernanceTransfer() returns()
func (_Clearnet *ClearnetTransactor) AcceptGovernanceTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "acceptGovernanceTransfer")
}

// AcceptGovernanceTransfer is a paid mutator transaction binding the contract method 0x2a90b8da.
//
// Solidity: function acceptGovernanceTransfer() returns()
func (_Clearnet *ClearnetSession) AcceptGovernanceTransfer() (*types.Transaction, error) {
	return _Clearnet.Contract.AcceptGovernanceTransfer(&_Clearnet.TransactOpts)
}

// AcceptGovernanceTransfer is a paid mutator transaction binding the contract method 0x2a90b8da.
//
// Solidity: function acceptGovernanceTransfer() returns()
func (_Clearnet *ClearnetTransactorSession) AcceptGovernanceTransfer() (*types.Transaction, error) {
	return _Clearnet.Contract.AcceptGovernanceTransfer(&_Clearnet.TransactOpts)
}

// AddRelayOperator is a paid mutator transaction binding the contract method 0x1ee8b2c0.
//
// Solidity: function addRelayOperator(address _relay) returns()
func (_Clearnet *ClearnetTransactor) AddRelayOperator(opts *bind.TransactOpts, _relay common.Address) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "addRelayOperator", _relay)
}

// AddRelayOperator is a paid mutator transaction binding the contract method 0x1ee8b2c0.
//
// Solidity: function addRelayOperator(address _relay) returns()
func (_Clearnet *ClearnetSession) AddRelayOperator(_relay common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.AddRelayOperator(&_Clearnet.TransactOpts, _relay)
}

// AddRelayOperator is a paid mutator transaction binding the contract method 0x1ee8b2c0.
//
// Solidity: function addRelayOperator(address _relay) returns()
func (_Clearnet *ClearnetTransactorSession) AddRelayOperator(_relay common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.AddRelayOperator(&_Clearnet.TransactOpts, _relay)
}

// ClosePaymentChannel is a paid mutator transaction binding the contract method 0x5959d8be.
//
// Solidity: function closePaymentChannel() returns()
func (_Clearnet *ClearnetTransactor) ClosePaymentChannel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "closePaymentChannel")
}

// ClosePaymentChannel is a paid mutator transaction binding the contract method 0x5959d8be.
//
// Solidity: function closePaymentChannel() returns()
func (_Clearnet *ClearnetSession) ClosePaymentChannel() (*types.Transaction, error) {
	return _Clearnet.Contract.ClosePaymentChannel(&_Clearnet.TransactOpts)
}

// ClosePaymentChannel is a paid mutator transaction binding the contract method 0x5959d8be.
//
// Solidity: function closePaymentChannel() returns()
func (_Clearnet *ClearnetTransactorSession) ClosePaymentChannel() (*types.Transaction, error) {
	return _Clearnet.Contract.ClosePaymentChannel(&_Clearnet.TransactOpts)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_Clearnet *ClearnetTransactor) DeregisterNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "deregisterNode")
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_Clearnet *ClearnetSession) DeregisterNode() (*types.Transaction, error) {
	return _Clearnet.Contract.DeregisterNode(&_Clearnet.TransactOpts)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_Clearnet *ClearnetTransactorSession) DeregisterNode() (*types.Transaction, error) {
	return _Clearnet.Contract.DeregisterNode(&_Clearnet.TransactOpts)
}

// InitiateGovernanceTransfer is a paid mutator transaction binding the contract method 0x8f8022bc.
//
// Solidity: function initiateGovernanceTransfer(address _newGovernance) returns()
func (_Clearnet *ClearnetTransactor) InitiateGovernanceTransfer(opts *bind.TransactOpts, _newGovernance common.Address) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "initiateGovernanceTransfer", _newGovernance)
}

// InitiateGovernanceTransfer is a paid mutator transaction binding the contract method 0x8f8022bc.
//
// Solidity: function initiateGovernanceTransfer(address _newGovernance) returns()
func (_Clearnet *ClearnetSession) InitiateGovernanceTransfer(_newGovernance common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.InitiateGovernanceTransfer(&_Clearnet.TransactOpts, _newGovernance)
}

// InitiateGovernanceTransfer is a paid mutator transaction binding the contract method 0x8f8022bc.
//
// Solidity: function initiateGovernanceTransfer(address _newGovernance) returns()
func (_Clearnet *ClearnetTransactorSession) InitiateGovernanceTransfer(_newGovernance common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.InitiateGovernanceTransfer(&_Clearnet.TransactOpts, _newGovernance)
}

// OpenPaymentChannel is a paid mutator transaction binding the contract method 0xfe841536.
//
// Solidity: function openPaymentChannel(uint256 _amount) returns()
func (_Clearnet *ClearnetTransactor) OpenPaymentChannel(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "openPaymentChannel", _amount)
}

// OpenPaymentChannel is a paid mutator transaction binding the contract method 0xfe841536.
//
// Solidity: function openPaymentChannel(uint256 _amount) returns()
func (_Clearnet *ClearnetSession) OpenPaymentChannel(_amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.OpenPaymentChannel(&_Clearnet.TransactOpts, _amount)
}

// OpenPaymentChannel is a paid mutator transaction binding the contract method 0xfe841536.
//
// Solidity: function openPaymentChannel(uint256 _amount) returns()
func (_Clearnet *ClearnetTransactorSession) OpenPaymentChannel(_amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.OpenPaymentChannel(&_Clearnet.TransactOpts, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Clearnet *ClearnetTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Clearnet *ClearnetSession) Pause() (*types.Transaction, error) {
	return _Clearnet.Contract.Pause(&_Clearnet.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Clearnet *ClearnetTransactorSession) Pause() (*types.Transaction, error) {
	return _Clearnet.Contract.Pause(&_Clearnet.TransactOpts)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x908466b9.
//
// Solidity: function processPayment(address _client, address _node, uint256 _minutesUsed, uint256 _nonce, bytes _clientSignature, bytes _nodeSignature, bytes _relaySignature) returns()
func (_Clearnet *ClearnetTransactor) ProcessPayment(opts *bind.TransactOpts, _client common.Address, _node common.Address, _minutesUsed *big.Int, _nonce *big.Int, _clientSignature []byte, _nodeSignature []byte, _relaySignature []byte) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "processPayment", _client, _node, _minutesUsed, _nonce, _clientSignature, _nodeSignature, _relaySignature)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x908466b9.
//
// Solidity: function processPayment(address _client, address _node, uint256 _minutesUsed, uint256 _nonce, bytes _clientSignature, bytes _nodeSignature, bytes _relaySignature) returns()
func (_Clearnet *ClearnetSession) ProcessPayment(_client common.Address, _node common.Address, _minutesUsed *big.Int, _nonce *big.Int, _clientSignature []byte, _nodeSignature []byte, _relaySignature []byte) (*types.Transaction, error) {
	return _Clearnet.Contract.ProcessPayment(&_Clearnet.TransactOpts, _client, _node, _minutesUsed, _nonce, _clientSignature, _nodeSignature, _relaySignature)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x908466b9.
//
// Solidity: function processPayment(address _client, address _node, uint256 _minutesUsed, uint256 _nonce, bytes _clientSignature, bytes _nodeSignature, bytes _relaySignature) returns()
func (_Clearnet *ClearnetTransactorSession) ProcessPayment(_client common.Address, _node common.Address, _minutesUsed *big.Int, _nonce *big.Int, _clientSignature []byte, _nodeSignature []byte, _relaySignature []byte) (*types.Transaction, error) {
	return _Clearnet.Contract.ProcessPayment(&_Clearnet.TransactOpts, _client, _node, _minutesUsed, _nonce, _clientSignature, _nodeSignature, _relaySignature)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xdab71a06.
//
// Solidity: function registerNode(string _ipAddress, uint16 _port, uint256 _pricePerMinute) returns()
func (_Clearnet *ClearnetTransactor) RegisterNode(opts *bind.TransactOpts, _ipAddress string, _port uint16, _pricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "registerNode", _ipAddress, _port, _pricePerMinute)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xdab71a06.
//
// Solidity: function registerNode(string _ipAddress, uint16 _port, uint256 _pricePerMinute) returns()
func (_Clearnet *ClearnetSession) RegisterNode(_ipAddress string, _port uint16, _pricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.RegisterNode(&_Clearnet.TransactOpts, _ipAddress, _port, _pricePerMinute)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xdab71a06.
//
// Solidity: function registerNode(string _ipAddress, uint16 _port, uint256 _pricePerMinute) returns()
func (_Clearnet *ClearnetTransactorSession) RegisterNode(_ipAddress string, _port uint16, _pricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.RegisterNode(&_Clearnet.TransactOpts, _ipAddress, _port, _pricePerMinute)
}

// RemoveRelayOperator is a paid mutator transaction binding the contract method 0x323c9b09.
//
// Solidity: function removeRelayOperator(address _relay) returns()
func (_Clearnet *ClearnetTransactor) RemoveRelayOperator(opts *bind.TransactOpts, _relay common.Address) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "removeRelayOperator", _relay)
}

// RemoveRelayOperator is a paid mutator transaction binding the contract method 0x323c9b09.
//
// Solidity: function removeRelayOperator(address _relay) returns()
func (_Clearnet *ClearnetSession) RemoveRelayOperator(_relay common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.RemoveRelayOperator(&_Clearnet.TransactOpts, _relay)
}

// RemoveRelayOperator is a paid mutator transaction binding the contract method 0x323c9b09.
//
// Solidity: function removeRelayOperator(address _relay) returns()
func (_Clearnet *ClearnetTransactorSession) RemoveRelayOperator(_relay common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.RemoveRelayOperator(&_Clearnet.TransactOpts, _relay)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clearnet *ClearnetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clearnet *ClearnetSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clearnet.Contract.RenounceOwnership(&_Clearnet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clearnet *ClearnetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clearnet.Contract.RenounceOwnership(&_Clearnet.TransactOpts)
}

// SlashNode is a paid mutator transaction binding the contract method 0xd2f9f285.
//
// Solidity: function slashNode(address _nodeID, uint256 _slashAmount) returns()
func (_Clearnet *ClearnetTransactor) SlashNode(opts *bind.TransactOpts, _nodeID common.Address, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "slashNode", _nodeID, _slashAmount)
}

// SlashNode is a paid mutator transaction binding the contract method 0xd2f9f285.
//
// Solidity: function slashNode(address _nodeID, uint256 _slashAmount) returns()
func (_Clearnet *ClearnetSession) SlashNode(_nodeID common.Address, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.SlashNode(&_Clearnet.TransactOpts, _nodeID, _slashAmount)
}

// SlashNode is a paid mutator transaction binding the contract method 0xd2f9f285.
//
// Solidity: function slashNode(address _nodeID, uint256 _slashAmount) returns()
func (_Clearnet *ClearnetTransactorSession) SlashNode(_nodeID common.Address, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.SlashNode(&_Clearnet.TransactOpts, _nodeID, _slashAmount)
}

// TopUpPaymentChannel is a paid mutator transaction binding the contract method 0xd60aad00.
//
// Solidity: function topUpPaymentChannel(uint256 _amount) returns()
func (_Clearnet *ClearnetTransactor) TopUpPaymentChannel(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "topUpPaymentChannel", _amount)
}

// TopUpPaymentChannel is a paid mutator transaction binding the contract method 0xd60aad00.
//
// Solidity: function topUpPaymentChannel(uint256 _amount) returns()
func (_Clearnet *ClearnetSession) TopUpPaymentChannel(_amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.TopUpPaymentChannel(&_Clearnet.TransactOpts, _amount)
}

// TopUpPaymentChannel is a paid mutator transaction binding the contract method 0xd60aad00.
//
// Solidity: function topUpPaymentChannel(uint256 _amount) returns()
func (_Clearnet *ClearnetTransactorSession) TopUpPaymentChannel(_amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.TopUpPaymentChannel(&_Clearnet.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clearnet *ClearnetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clearnet *ClearnetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.TransferOwnership(&_Clearnet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clearnet *ClearnetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clearnet.Contract.TransferOwnership(&_Clearnet.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Clearnet *ClearnetTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Clearnet *ClearnetSession) Unpause() (*types.Transaction, error) {
	return _Clearnet.Contract.Unpause(&_Clearnet.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Clearnet *ClearnetTransactorSession) Unpause() (*types.Transaction, error) {
	return _Clearnet.Contract.Unpause(&_Clearnet.TransactOpts)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0x9b2218ff.
//
// Solidity: function updateNodeInfo(string _newIpAddress, uint16 _newPort) returns()
func (_Clearnet *ClearnetTransactor) UpdateNodeInfo(opts *bind.TransactOpts, _newIpAddress string, _newPort uint16) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "updateNodeInfo", _newIpAddress, _newPort)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0x9b2218ff.
//
// Solidity: function updateNodeInfo(string _newIpAddress, uint16 _newPort) returns()
func (_Clearnet *ClearnetSession) UpdateNodeInfo(_newIpAddress string, _newPort uint16) (*types.Transaction, error) {
	return _Clearnet.Contract.UpdateNodeInfo(&_Clearnet.TransactOpts, _newIpAddress, _newPort)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0x9b2218ff.
//
// Solidity: function updateNodeInfo(string _newIpAddress, uint16 _newPort) returns()
func (_Clearnet *ClearnetTransactorSession) UpdateNodeInfo(_newIpAddress string, _newPort uint16) (*types.Transaction, error) {
	return _Clearnet.Contract.UpdateNodeInfo(&_Clearnet.TransactOpts, _newIpAddress, _newPort)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPricePerMinute) returns()
func (_Clearnet *ClearnetTransactor) UpdatePrice(opts *bind.TransactOpts, _newPricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "updatePrice", _newPricePerMinute)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPricePerMinute) returns()
func (_Clearnet *ClearnetSession) UpdatePrice(_newPricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.UpdatePrice(&_Clearnet.TransactOpts, _newPricePerMinute)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPricePerMinute) returns()
func (_Clearnet *ClearnetTransactorSession) UpdatePrice(_newPricePerMinute *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.UpdatePrice(&_Clearnet.TransactOpts, _newPricePerMinute)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xe2c13e32.
//
// Solidity: function updateReputation(address _nodeID, bool _successfulSession) returns()
func (_Clearnet *ClearnetTransactor) UpdateReputation(opts *bind.TransactOpts, _nodeID common.Address, _successfulSession bool) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "updateReputation", _nodeID, _successfulSession)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xe2c13e32.
//
// Solidity: function updateReputation(address _nodeID, bool _successfulSession) returns()
func (_Clearnet *ClearnetSession) UpdateReputation(_nodeID common.Address, _successfulSession bool) (*types.Transaction, error) {
	return _Clearnet.Contract.UpdateReputation(&_Clearnet.TransactOpts, _nodeID, _successfulSession)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xe2c13e32.
//
// Solidity: function updateReputation(address _nodeID, bool _successfulSession) returns()
func (_Clearnet *ClearnetTransactorSession) UpdateReputation(_nodeID common.Address, _successfulSession bool) (*types.Transaction, error) {
	return _Clearnet.Contract.UpdateReputation(&_Clearnet.TransactOpts, _nodeID, _successfulSession)
}

// WithdrawTreasuryFunds is a paid mutator transaction binding the contract method 0x661b8702.
//
// Solidity: function withdrawTreasuryFunds(uint256 _amount) returns()
func (_Clearnet *ClearnetTransactor) WithdrawTreasuryFunds(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.contract.Transact(opts, "withdrawTreasuryFunds", _amount)
}

// WithdrawTreasuryFunds is a paid mutator transaction binding the contract method 0x661b8702.
//
// Solidity: function withdrawTreasuryFunds(uint256 _amount) returns()
func (_Clearnet *ClearnetSession) WithdrawTreasuryFunds(_amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.WithdrawTreasuryFunds(&_Clearnet.TransactOpts, _amount)
}

// WithdrawTreasuryFunds is a paid mutator transaction binding the contract method 0x661b8702.
//
// Solidity: function withdrawTreasuryFunds(uint256 _amount) returns()
func (_Clearnet *ClearnetTransactorSession) WithdrawTreasuryFunds(_amount *big.Int) (*types.Transaction, error) {
	return _Clearnet.Contract.WithdrawTreasuryFunds(&_Clearnet.TransactOpts, _amount)
}

// ClearnetGovernanceTransferInitiatedIterator is returned from FilterGovernanceTransferInitiated and is used to iterate over the raw logs and unpacked data for GovernanceTransferInitiated events raised by the Clearnet contract.
type ClearnetGovernanceTransferInitiatedIterator struct {
	Event *ClearnetGovernanceTransferInitiated // Event containing the contract specifics and raw log

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
func (it *ClearnetGovernanceTransferInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetGovernanceTransferInitiated)
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
		it.Event = new(ClearnetGovernanceTransferInitiated)
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
func (it *ClearnetGovernanceTransferInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetGovernanceTransferInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetGovernanceTransferInitiated represents a GovernanceTransferInitiated event raised by the Clearnet contract.
type ClearnetGovernanceTransferInitiated struct {
	CurrentGovernance common.Address
	PendingGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferInitiated is a free log retrieval operation binding the contract event 0x10705730e98da50fb257fd92cd2817af22aba5a69113af740907ab35133237a1.
//
// Solidity: event GovernanceTransferInitiated(address indexed currentGovernance, address indexed pendingGovernance)
func (_Clearnet *ClearnetFilterer) FilterGovernanceTransferInitiated(opts *bind.FilterOpts, currentGovernance []common.Address, pendingGovernance []common.Address) (*ClearnetGovernanceTransferInitiatedIterator, error) {

	var currentGovernanceRule []interface{}
	for _, currentGovernanceItem := range currentGovernance {
		currentGovernanceRule = append(currentGovernanceRule, currentGovernanceItem)
	}
	var pendingGovernanceRule []interface{}
	for _, pendingGovernanceItem := range pendingGovernance {
		pendingGovernanceRule = append(pendingGovernanceRule, pendingGovernanceItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "GovernanceTransferInitiated", currentGovernanceRule, pendingGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetGovernanceTransferInitiatedIterator{contract: _Clearnet.contract, event: "GovernanceTransferInitiated", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferInitiated is a free log subscription operation binding the contract event 0x10705730e98da50fb257fd92cd2817af22aba5a69113af740907ab35133237a1.
//
// Solidity: event GovernanceTransferInitiated(address indexed currentGovernance, address indexed pendingGovernance)
func (_Clearnet *ClearnetFilterer) WatchGovernanceTransferInitiated(opts *bind.WatchOpts, sink chan<- *ClearnetGovernanceTransferInitiated, currentGovernance []common.Address, pendingGovernance []common.Address) (event.Subscription, error) {

	var currentGovernanceRule []interface{}
	for _, currentGovernanceItem := range currentGovernance {
		currentGovernanceRule = append(currentGovernanceRule, currentGovernanceItem)
	}
	var pendingGovernanceRule []interface{}
	for _, pendingGovernanceItem := range pendingGovernance {
		pendingGovernanceRule = append(pendingGovernanceRule, pendingGovernanceItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "GovernanceTransferInitiated", currentGovernanceRule, pendingGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetGovernanceTransferInitiated)
				if err := _Clearnet.contract.UnpackLog(event, "GovernanceTransferInitiated", log); err != nil {
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

// ParseGovernanceTransferInitiated is a log parse operation binding the contract event 0x10705730e98da50fb257fd92cd2817af22aba5a69113af740907ab35133237a1.
//
// Solidity: event GovernanceTransferInitiated(address indexed currentGovernance, address indexed pendingGovernance)
func (_Clearnet *ClearnetFilterer) ParseGovernanceTransferInitiated(log types.Log) (*ClearnetGovernanceTransferInitiated, error) {
	event := new(ClearnetGovernanceTransferInitiated)
	if err := _Clearnet.contract.UnpackLog(event, "GovernanceTransferInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the Clearnet contract.
type ClearnetGovernanceTransferredIterator struct {
	Event *ClearnetGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *ClearnetGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetGovernanceTransferred)
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
		it.Event = new(ClearnetGovernanceTransferred)
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
func (it *ClearnetGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetGovernanceTransferred represents a GovernanceTransferred event raised by the Clearnet contract.
type ClearnetGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Clearnet *ClearnetFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*ClearnetGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetGovernanceTransferredIterator{contract: _Clearnet.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Clearnet *ClearnetFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *ClearnetGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetGovernanceTransferred)
				if err := _Clearnet.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Clearnet *ClearnetFilterer) ParseGovernanceTransferred(log types.Log) (*ClearnetGovernanceTransferred, error) {
	event := new(ClearnetGovernanceTransferred)
	if err := _Clearnet.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetNodeDeregisteredIterator is returned from FilterNodeDeregistered and is used to iterate over the raw logs and unpacked data for NodeDeregistered events raised by the Clearnet contract.
type ClearnetNodeDeregisteredIterator struct {
	Event *ClearnetNodeDeregistered // Event containing the contract specifics and raw log

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
func (it *ClearnetNodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetNodeDeregistered)
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
		it.Event = new(ClearnetNodeDeregistered)
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
func (it *ClearnetNodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetNodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetNodeDeregistered represents a NodeDeregistered event raised by the Clearnet contract.
type ClearnetNodeDeregistered struct {
	NodeID        common.Address
	ReturnedStake *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeDeregistered is a free log retrieval operation binding the contract event 0xaa433836fcb47675047002afc96e88cdbb2ba277f2508b8719e646fcd9b116f1.
//
// Solidity: event NodeDeregistered(address indexed nodeID, uint256 returnedStake)
func (_Clearnet *ClearnetFilterer) FilterNodeDeregistered(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetNodeDeregisteredIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "NodeDeregistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetNodeDeregisteredIterator{contract: _Clearnet.contract, event: "NodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchNodeDeregistered is a free log subscription operation binding the contract event 0xaa433836fcb47675047002afc96e88cdbb2ba277f2508b8719e646fcd9b116f1.
//
// Solidity: event NodeDeregistered(address indexed nodeID, uint256 returnedStake)
func (_Clearnet *ClearnetFilterer) WatchNodeDeregistered(opts *bind.WatchOpts, sink chan<- *ClearnetNodeDeregistered, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "NodeDeregistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetNodeDeregistered)
				if err := _Clearnet.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
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
func (_Clearnet *ClearnetFilterer) ParseNodeDeregistered(log types.Log) (*ClearnetNodeDeregistered, error) {
	event := new(ClearnetNodeDeregistered)
	if err := _Clearnet.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetNodePriceUpdatedIterator is returned from FilterNodePriceUpdated and is used to iterate over the raw logs and unpacked data for NodePriceUpdated events raised by the Clearnet contract.
type ClearnetNodePriceUpdatedIterator struct {
	Event *ClearnetNodePriceUpdated // Event containing the contract specifics and raw log

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
func (it *ClearnetNodePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetNodePriceUpdated)
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
		it.Event = new(ClearnetNodePriceUpdated)
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
func (it *ClearnetNodePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetNodePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetNodePriceUpdated represents a NodePriceUpdated event raised by the Clearnet contract.
type ClearnetNodePriceUpdated struct {
	NodeID            common.Address
	NewPricePerMinute *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNodePriceUpdated is a free log retrieval operation binding the contract event 0xaa357beaa928ff528af3cf3853f5075732f925103274cf6652df31fe9b224b0a.
//
// Solidity: event NodePriceUpdated(address indexed nodeID, uint256 newPricePerMinute)
func (_Clearnet *ClearnetFilterer) FilterNodePriceUpdated(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetNodePriceUpdatedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "NodePriceUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetNodePriceUpdatedIterator{contract: _Clearnet.contract, event: "NodePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchNodePriceUpdated is a free log subscription operation binding the contract event 0xaa357beaa928ff528af3cf3853f5075732f925103274cf6652df31fe9b224b0a.
//
// Solidity: event NodePriceUpdated(address indexed nodeID, uint256 newPricePerMinute)
func (_Clearnet *ClearnetFilterer) WatchNodePriceUpdated(opts *bind.WatchOpts, sink chan<- *ClearnetNodePriceUpdated, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "NodePriceUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetNodePriceUpdated)
				if err := _Clearnet.contract.UnpackLog(event, "NodePriceUpdated", log); err != nil {
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
func (_Clearnet *ClearnetFilterer) ParseNodePriceUpdated(log types.Log) (*ClearnetNodePriceUpdated, error) {
	event := new(ClearnetNodePriceUpdated)
	if err := _Clearnet.contract.UnpackLog(event, "NodePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the Clearnet contract.
type ClearnetNodeRegisteredIterator struct {
	Event *ClearnetNodeRegistered // Event containing the contract specifics and raw log

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
func (it *ClearnetNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetNodeRegistered)
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
		it.Event = new(ClearnetNodeRegistered)
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
func (it *ClearnetNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetNodeRegistered represents a NodeRegistered event raised by the Clearnet contract.
type ClearnetNodeRegistered struct {
	NodeID         common.Address
	IpAddress      string
	Port           uint16
	PricePerMinute *big.Int
	StakeAmount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0xe480bbb8644627415bc0bac0b7bba21f29c00ca1557b2c0098a6354e3799c241.
//
// Solidity: event NodeRegistered(address indexed nodeID, string ipAddress, uint16 port, uint256 pricePerMinute, uint256 stakeAmount)
func (_Clearnet *ClearnetFilterer) FilterNodeRegistered(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetNodeRegisteredIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "NodeRegistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetNodeRegisteredIterator{contract: _Clearnet.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0xe480bbb8644627415bc0bac0b7bba21f29c00ca1557b2c0098a6354e3799c241.
//
// Solidity: event NodeRegistered(address indexed nodeID, string ipAddress, uint16 port, uint256 pricePerMinute, uint256 stakeAmount)
func (_Clearnet *ClearnetFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *ClearnetNodeRegistered, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "NodeRegistered", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetNodeRegistered)
				if err := _Clearnet.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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

// ParseNodeRegistered is a log parse operation binding the contract event 0xe480bbb8644627415bc0bac0b7bba21f29c00ca1557b2c0098a6354e3799c241.
//
// Solidity: event NodeRegistered(address indexed nodeID, string ipAddress, uint16 port, uint256 pricePerMinute, uint256 stakeAmount)
func (_Clearnet *ClearnetFilterer) ParseNodeRegistered(log types.Log) (*ClearnetNodeRegistered, error) {
	event := new(ClearnetNodeRegistered)
	if err := _Clearnet.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetNodeSlashedIterator is returned from FilterNodeSlashed and is used to iterate over the raw logs and unpacked data for NodeSlashed events raised by the Clearnet contract.
type ClearnetNodeSlashedIterator struct {
	Event *ClearnetNodeSlashed // Event containing the contract specifics and raw log

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
func (it *ClearnetNodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetNodeSlashed)
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
		it.Event = new(ClearnetNodeSlashed)
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
func (it *ClearnetNodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetNodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetNodeSlashed represents a NodeSlashed event raised by the Clearnet contract.
type ClearnetNodeSlashed struct {
	NodeID         common.Address
	SlashAmount    *big.Int
	RemainingStake *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeSlashed is a free log retrieval operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeID, uint256 slashAmount, uint256 remainingStake)
func (_Clearnet *ClearnetFilterer) FilterNodeSlashed(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetNodeSlashedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "NodeSlashed", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetNodeSlashedIterator{contract: _Clearnet.contract, event: "NodeSlashed", logs: logs, sub: sub}, nil
}

// WatchNodeSlashed is a free log subscription operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeID, uint256 slashAmount, uint256 remainingStake)
func (_Clearnet *ClearnetFilterer) WatchNodeSlashed(opts *bind.WatchOpts, sink chan<- *ClearnetNodeSlashed, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "NodeSlashed", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetNodeSlashed)
				if err := _Clearnet.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
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

// ParseNodeSlashed is a log parse operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeID, uint256 slashAmount, uint256 remainingStake)
func (_Clearnet *ClearnetFilterer) ParseNodeSlashed(log types.Log) (*ClearnetNodeSlashed, error) {
	event := new(ClearnetNodeSlashed)
	if err := _Clearnet.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetNodeUpdatedIterator is returned from FilterNodeUpdated and is used to iterate over the raw logs and unpacked data for NodeUpdated events raised by the Clearnet contract.
type ClearnetNodeUpdatedIterator struct {
	Event *ClearnetNodeUpdated // Event containing the contract specifics and raw log

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
func (it *ClearnetNodeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetNodeUpdated)
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
		it.Event = new(ClearnetNodeUpdated)
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
func (it *ClearnetNodeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetNodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetNodeUpdated represents a NodeUpdated event raised by the Clearnet contract.
type ClearnetNodeUpdated struct {
	NodeID  common.Address
	NewIp   string
	NewPort uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeUpdated is a free log retrieval operation binding the contract event 0xaea8346ad1588b9f6e64c53d2298294f2b9f1310ca57bade59c1cd5a2e3213b1.
//
// Solidity: event NodeUpdated(address indexed nodeID, string newIp, uint16 newPort)
func (_Clearnet *ClearnetFilterer) FilterNodeUpdated(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetNodeUpdatedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "NodeUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetNodeUpdatedIterator{contract: _Clearnet.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeUpdated is a free log subscription operation binding the contract event 0xaea8346ad1588b9f6e64c53d2298294f2b9f1310ca57bade59c1cd5a2e3213b1.
//
// Solidity: event NodeUpdated(address indexed nodeID, string newIp, uint16 newPort)
func (_Clearnet *ClearnetFilterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *ClearnetNodeUpdated, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "NodeUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetNodeUpdated)
				if err := _Clearnet.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
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

// ParseNodeUpdated is a log parse operation binding the contract event 0xaea8346ad1588b9f6e64c53d2298294f2b9f1310ca57bade59c1cd5a2e3213b1.
//
// Solidity: event NodeUpdated(address indexed nodeID, string newIp, uint16 newPort)
func (_Clearnet *ClearnetFilterer) ParseNodeUpdated(log types.Log) (*ClearnetNodeUpdated, error) {
	event := new(ClearnetNodeUpdated)
	if err := _Clearnet.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Clearnet contract.
type ClearnetOwnershipTransferredIterator struct {
	Event *ClearnetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClearnetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetOwnershipTransferred)
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
		it.Event = new(ClearnetOwnershipTransferred)
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
func (it *ClearnetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetOwnershipTransferred represents a OwnershipTransferred event raised by the Clearnet contract.
type ClearnetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clearnet *ClearnetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClearnetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetOwnershipTransferredIterator{contract: _Clearnet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clearnet *ClearnetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClearnetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetOwnershipTransferred)
				if err := _Clearnet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Clearnet *ClearnetFilterer) ParseOwnershipTransferred(log types.Log) (*ClearnetOwnershipTransferred, error) {
	event := new(ClearnetOwnershipTransferred)
	if err := _Clearnet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Clearnet contract.
type ClearnetPausedIterator struct {
	Event *ClearnetPaused // Event containing the contract specifics and raw log

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
func (it *ClearnetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetPaused)
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
		it.Event = new(ClearnetPaused)
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
func (it *ClearnetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetPaused represents a Paused event raised by the Clearnet contract.
type ClearnetPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_Clearnet *ClearnetFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ClearnetPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetPausedIterator{contract: _Clearnet.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_Clearnet *ClearnetFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ClearnetPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetPaused)
				if err := _Clearnet.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address indexed account)
func (_Clearnet *ClearnetFilterer) ParsePaused(log types.Log) (*ClearnetPaused, error) {
	event := new(ClearnetPaused)
	if err := _Clearnet.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetPaymentChannelClosedIterator is returned from FilterPaymentChannelClosed and is used to iterate over the raw logs and unpacked data for PaymentChannelClosed events raised by the Clearnet contract.
type ClearnetPaymentChannelClosedIterator struct {
	Event *ClearnetPaymentChannelClosed // Event containing the contract specifics and raw log

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
func (it *ClearnetPaymentChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetPaymentChannelClosed)
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
		it.Event = new(ClearnetPaymentChannelClosed)
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
func (it *ClearnetPaymentChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetPaymentChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetPaymentChannelClosed represents a PaymentChannelClosed event raised by the Clearnet contract.
type ClearnetPaymentChannelClosed struct {
	Client       common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaymentChannelClosed is a free log retrieval operation binding the contract event 0x46a3246526c7e80ff10cd40c7426786202a6b08a50e81dae3d7db9f80349c7ab.
//
// Solidity: event PaymentChannelClosed(address indexed client, uint256 refundAmount)
func (_Clearnet *ClearnetFilterer) FilterPaymentChannelClosed(opts *bind.FilterOpts, client []common.Address) (*ClearnetPaymentChannelClosedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "PaymentChannelClosed", clientRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetPaymentChannelClosedIterator{contract: _Clearnet.contract, event: "PaymentChannelClosed", logs: logs, sub: sub}, nil
}

// WatchPaymentChannelClosed is a free log subscription operation binding the contract event 0x46a3246526c7e80ff10cd40c7426786202a6b08a50e81dae3d7db9f80349c7ab.
//
// Solidity: event PaymentChannelClosed(address indexed client, uint256 refundAmount)
func (_Clearnet *ClearnetFilterer) WatchPaymentChannelClosed(opts *bind.WatchOpts, sink chan<- *ClearnetPaymentChannelClosed, client []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "PaymentChannelClosed", clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetPaymentChannelClosed)
				if err := _Clearnet.contract.UnpackLog(event, "PaymentChannelClosed", log); err != nil {
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

// ParsePaymentChannelClosed is a log parse operation binding the contract event 0x46a3246526c7e80ff10cd40c7426786202a6b08a50e81dae3d7db9f80349c7ab.
//
// Solidity: event PaymentChannelClosed(address indexed client, uint256 refundAmount)
func (_Clearnet *ClearnetFilterer) ParsePaymentChannelClosed(log types.Log) (*ClearnetPaymentChannelClosed, error) {
	event := new(ClearnetPaymentChannelClosed)
	if err := _Clearnet.contract.UnpackLog(event, "PaymentChannelClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetPaymentChannelOpenedIterator is returned from FilterPaymentChannelOpened and is used to iterate over the raw logs and unpacked data for PaymentChannelOpened events raised by the Clearnet contract.
type ClearnetPaymentChannelOpenedIterator struct {
	Event *ClearnetPaymentChannelOpened // Event containing the contract specifics and raw log

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
func (it *ClearnetPaymentChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetPaymentChannelOpened)
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
		it.Event = new(ClearnetPaymentChannelOpened)
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
func (it *ClearnetPaymentChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetPaymentChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetPaymentChannelOpened represents a PaymentChannelOpened event raised by the Clearnet contract.
type ClearnetPaymentChannelOpened struct {
	Client common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentChannelOpened is a free log retrieval operation binding the contract event 0x6ccbd6b45459eec478864788528dfc5c6b33cb7b86fcdad9b409d3b036f8e77c.
//
// Solidity: event PaymentChannelOpened(address indexed client, uint256 amount)
func (_Clearnet *ClearnetFilterer) FilterPaymentChannelOpened(opts *bind.FilterOpts, client []common.Address) (*ClearnetPaymentChannelOpenedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "PaymentChannelOpened", clientRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetPaymentChannelOpenedIterator{contract: _Clearnet.contract, event: "PaymentChannelOpened", logs: logs, sub: sub}, nil
}

// WatchPaymentChannelOpened is a free log subscription operation binding the contract event 0x6ccbd6b45459eec478864788528dfc5c6b33cb7b86fcdad9b409d3b036f8e77c.
//
// Solidity: event PaymentChannelOpened(address indexed client, uint256 amount)
func (_Clearnet *ClearnetFilterer) WatchPaymentChannelOpened(opts *bind.WatchOpts, sink chan<- *ClearnetPaymentChannelOpened, client []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "PaymentChannelOpened", clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetPaymentChannelOpened)
				if err := _Clearnet.contract.UnpackLog(event, "PaymentChannelOpened", log); err != nil {
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

// ParsePaymentChannelOpened is a log parse operation binding the contract event 0x6ccbd6b45459eec478864788528dfc5c6b33cb7b86fcdad9b409d3b036f8e77c.
//
// Solidity: event PaymentChannelOpened(address indexed client, uint256 amount)
func (_Clearnet *ClearnetFilterer) ParsePaymentChannelOpened(log types.Log) (*ClearnetPaymentChannelOpened, error) {
	event := new(ClearnetPaymentChannelOpened)
	if err := _Clearnet.contract.UnpackLog(event, "PaymentChannelOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetPaymentChannelToppedUpIterator is returned from FilterPaymentChannelToppedUp and is used to iterate over the raw logs and unpacked data for PaymentChannelToppedUp events raised by the Clearnet contract.
type ClearnetPaymentChannelToppedUpIterator struct {
	Event *ClearnetPaymentChannelToppedUp // Event containing the contract specifics and raw log

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
func (it *ClearnetPaymentChannelToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetPaymentChannelToppedUp)
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
		it.Event = new(ClearnetPaymentChannelToppedUp)
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
func (it *ClearnetPaymentChannelToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetPaymentChannelToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetPaymentChannelToppedUp represents a PaymentChannelToppedUp event raised by the Clearnet contract.
type ClearnetPaymentChannelToppedUp struct {
	Client     common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaymentChannelToppedUp is a free log retrieval operation binding the contract event 0x55546d832e5d12d1e843a67fb0198c156e51766c5781e7f49476e79f284fb1f4.
//
// Solidity: event PaymentChannelToppedUp(address indexed client, uint256 amount, uint256 newBalance)
func (_Clearnet *ClearnetFilterer) FilterPaymentChannelToppedUp(opts *bind.FilterOpts, client []common.Address) (*ClearnetPaymentChannelToppedUpIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "PaymentChannelToppedUp", clientRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetPaymentChannelToppedUpIterator{contract: _Clearnet.contract, event: "PaymentChannelToppedUp", logs: logs, sub: sub}, nil
}

// WatchPaymentChannelToppedUp is a free log subscription operation binding the contract event 0x55546d832e5d12d1e843a67fb0198c156e51766c5781e7f49476e79f284fb1f4.
//
// Solidity: event PaymentChannelToppedUp(address indexed client, uint256 amount, uint256 newBalance)
func (_Clearnet *ClearnetFilterer) WatchPaymentChannelToppedUp(opts *bind.WatchOpts, sink chan<- *ClearnetPaymentChannelToppedUp, client []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "PaymentChannelToppedUp", clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetPaymentChannelToppedUp)
				if err := _Clearnet.contract.UnpackLog(event, "PaymentChannelToppedUp", log); err != nil {
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

// ParsePaymentChannelToppedUp is a log parse operation binding the contract event 0x55546d832e5d12d1e843a67fb0198c156e51766c5781e7f49476e79f284fb1f4.
//
// Solidity: event PaymentChannelToppedUp(address indexed client, uint256 amount, uint256 newBalance)
func (_Clearnet *ClearnetFilterer) ParsePaymentChannelToppedUp(log types.Log) (*ClearnetPaymentChannelToppedUp, error) {
	event := new(ClearnetPaymentChannelToppedUp)
	if err := _Clearnet.contract.UnpackLog(event, "PaymentChannelToppedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetPaymentProcessedIterator is returned from FilterPaymentProcessed and is used to iterate over the raw logs and unpacked data for PaymentProcessed events raised by the Clearnet contract.
type ClearnetPaymentProcessedIterator struct {
	Event *ClearnetPaymentProcessed // Event containing the contract specifics and raw log

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
func (it *ClearnetPaymentProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetPaymentProcessed)
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
		it.Event = new(ClearnetPaymentProcessed)
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
func (it *ClearnetPaymentProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetPaymentProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetPaymentProcessed represents a PaymentProcessed event raised by the Clearnet contract.
type ClearnetPaymentProcessed struct {
	Client        common.Address
	NodeID        common.Address
	Amount        *big.Int
	MinutesUsed   *big.Int
	NodeShare     *big.Int
	TreasuryShare *big.Int
	RelayShare    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentProcessed is a free log retrieval operation binding the contract event 0x9383ea3334618ec4baeccb12cd192b2f264193c226bcf129ea3b7334f4663187.
//
// Solidity: event PaymentProcessed(address indexed client, address indexed nodeID, uint256 amount, uint256 minutesUsed, uint256 nodeShare, uint256 treasuryShare, uint256 relayShare)
func (_Clearnet *ClearnetFilterer) FilterPaymentProcessed(opts *bind.FilterOpts, client []common.Address, nodeID []common.Address) (*ClearnetPaymentProcessedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "PaymentProcessed", clientRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetPaymentProcessedIterator{contract: _Clearnet.contract, event: "PaymentProcessed", logs: logs, sub: sub}, nil
}

// WatchPaymentProcessed is a free log subscription operation binding the contract event 0x9383ea3334618ec4baeccb12cd192b2f264193c226bcf129ea3b7334f4663187.
//
// Solidity: event PaymentProcessed(address indexed client, address indexed nodeID, uint256 amount, uint256 minutesUsed, uint256 nodeShare, uint256 treasuryShare, uint256 relayShare)
func (_Clearnet *ClearnetFilterer) WatchPaymentProcessed(opts *bind.WatchOpts, sink chan<- *ClearnetPaymentProcessed, client []common.Address, nodeID []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "PaymentProcessed", clientRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetPaymentProcessed)
				if err := _Clearnet.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
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

// ParsePaymentProcessed is a log parse operation binding the contract event 0x9383ea3334618ec4baeccb12cd192b2f264193c226bcf129ea3b7334f4663187.
//
// Solidity: event PaymentProcessed(address indexed client, address indexed nodeID, uint256 amount, uint256 minutesUsed, uint256 nodeShare, uint256 treasuryShare, uint256 relayShare)
func (_Clearnet *ClearnetFilterer) ParsePaymentProcessed(log types.Log) (*ClearnetPaymentProcessed, error) {
	event := new(ClearnetPaymentProcessed)
	if err := _Clearnet.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetRelayOperatorAddedIterator is returned from FilterRelayOperatorAdded and is used to iterate over the raw logs and unpacked data for RelayOperatorAdded events raised by the Clearnet contract.
type ClearnetRelayOperatorAddedIterator struct {
	Event *ClearnetRelayOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ClearnetRelayOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetRelayOperatorAdded)
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
		it.Event = new(ClearnetRelayOperatorAdded)
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
func (it *ClearnetRelayOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetRelayOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetRelayOperatorAdded represents a RelayOperatorAdded event raised by the Clearnet contract.
type ClearnetRelayOperatorAdded struct {
	Relay common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRelayOperatorAdded is a free log retrieval operation binding the contract event 0x30e3a6db2d3b6d9fd79c76cdbdd1e2373e15863cb257e2aa761f148546004126.
//
// Solidity: event RelayOperatorAdded(address indexed relay)
func (_Clearnet *ClearnetFilterer) FilterRelayOperatorAdded(opts *bind.FilterOpts, relay []common.Address) (*ClearnetRelayOperatorAddedIterator, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "RelayOperatorAdded", relayRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetRelayOperatorAddedIterator{contract: _Clearnet.contract, event: "RelayOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchRelayOperatorAdded is a free log subscription operation binding the contract event 0x30e3a6db2d3b6d9fd79c76cdbdd1e2373e15863cb257e2aa761f148546004126.
//
// Solidity: event RelayOperatorAdded(address indexed relay)
func (_Clearnet *ClearnetFilterer) WatchRelayOperatorAdded(opts *bind.WatchOpts, sink chan<- *ClearnetRelayOperatorAdded, relay []common.Address) (event.Subscription, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "RelayOperatorAdded", relayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetRelayOperatorAdded)
				if err := _Clearnet.contract.UnpackLog(event, "RelayOperatorAdded", log); err != nil {
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

// ParseRelayOperatorAdded is a log parse operation binding the contract event 0x30e3a6db2d3b6d9fd79c76cdbdd1e2373e15863cb257e2aa761f148546004126.
//
// Solidity: event RelayOperatorAdded(address indexed relay)
func (_Clearnet *ClearnetFilterer) ParseRelayOperatorAdded(log types.Log) (*ClearnetRelayOperatorAdded, error) {
	event := new(ClearnetRelayOperatorAdded)
	if err := _Clearnet.contract.UnpackLog(event, "RelayOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetRelayOperatorRemovedIterator is returned from FilterRelayOperatorRemoved and is used to iterate over the raw logs and unpacked data for RelayOperatorRemoved events raised by the Clearnet contract.
type ClearnetRelayOperatorRemovedIterator struct {
	Event *ClearnetRelayOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ClearnetRelayOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetRelayOperatorRemoved)
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
		it.Event = new(ClearnetRelayOperatorRemoved)
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
func (it *ClearnetRelayOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetRelayOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetRelayOperatorRemoved represents a RelayOperatorRemoved event raised by the Clearnet contract.
type ClearnetRelayOperatorRemoved struct {
	Relay common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRelayOperatorRemoved is a free log retrieval operation binding the contract event 0xa9160e48ced70f0a7bbb59246f46cfd9fb43d7750f9f412e3d434dd9fd5cbc94.
//
// Solidity: event RelayOperatorRemoved(address indexed relay)
func (_Clearnet *ClearnetFilterer) FilterRelayOperatorRemoved(opts *bind.FilterOpts, relay []common.Address) (*ClearnetRelayOperatorRemovedIterator, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "RelayOperatorRemoved", relayRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetRelayOperatorRemovedIterator{contract: _Clearnet.contract, event: "RelayOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayOperatorRemoved is a free log subscription operation binding the contract event 0xa9160e48ced70f0a7bbb59246f46cfd9fb43d7750f9f412e3d434dd9fd5cbc94.
//
// Solidity: event RelayOperatorRemoved(address indexed relay)
func (_Clearnet *ClearnetFilterer) WatchRelayOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ClearnetRelayOperatorRemoved, relay []common.Address) (event.Subscription, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "RelayOperatorRemoved", relayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetRelayOperatorRemoved)
				if err := _Clearnet.contract.UnpackLog(event, "RelayOperatorRemoved", log); err != nil {
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

// ParseRelayOperatorRemoved is a log parse operation binding the contract event 0xa9160e48ced70f0a7bbb59246f46cfd9fb43d7750f9f412e3d434dd9fd5cbc94.
//
// Solidity: event RelayOperatorRemoved(address indexed relay)
func (_Clearnet *ClearnetFilterer) ParseRelayOperatorRemoved(log types.Log) (*ClearnetRelayOperatorRemoved, error) {
	event := new(ClearnetRelayOperatorRemoved)
	if err := _Clearnet.contract.UnpackLog(event, "RelayOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetReputationUpdatedIterator is returned from FilterReputationUpdated and is used to iterate over the raw logs and unpacked data for ReputationUpdated events raised by the Clearnet contract.
type ClearnetReputationUpdatedIterator struct {
	Event *ClearnetReputationUpdated // Event containing the contract specifics and raw log

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
func (it *ClearnetReputationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetReputationUpdated)
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
		it.Event = new(ClearnetReputationUpdated)
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
func (it *ClearnetReputationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetReputationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetReputationUpdated represents a ReputationUpdated event raised by the Clearnet contract.
type ClearnetReputationUpdated struct {
	NodeID   common.Address
	NewScore *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReputationUpdated is a free log retrieval operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed nodeID, uint256 newScore)
func (_Clearnet *ClearnetFilterer) FilterReputationUpdated(opts *bind.FilterOpts, nodeID []common.Address) (*ClearnetReputationUpdatedIterator, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "ReputationUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetReputationUpdatedIterator{contract: _Clearnet.contract, event: "ReputationUpdated", logs: logs, sub: sub}, nil
}

// WatchReputationUpdated is a free log subscription operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed nodeID, uint256 newScore)
func (_Clearnet *ClearnetFilterer) WatchReputationUpdated(opts *bind.WatchOpts, sink chan<- *ClearnetReputationUpdated, nodeID []common.Address) (event.Subscription, error) {

	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "ReputationUpdated", nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetReputationUpdated)
				if err := _Clearnet.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
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
func (_Clearnet *ClearnetFilterer) ParseReputationUpdated(log types.Log) (*ClearnetReputationUpdated, error) {
	event := new(ClearnetReputationUpdated)
	if err := _Clearnet.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClearnetUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Clearnet contract.
type ClearnetUnpausedIterator struct {
	Event *ClearnetUnpaused // Event containing the contract specifics and raw log

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
func (it *ClearnetUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClearnetUnpaused)
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
		it.Event = new(ClearnetUnpaused)
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
func (it *ClearnetUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClearnetUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClearnetUnpaused represents a Unpaused event raised by the Clearnet contract.
type ClearnetUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed account)
func (_Clearnet *ClearnetFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ClearnetUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Clearnet.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ClearnetUnpausedIterator{contract: _Clearnet.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed account)
func (_Clearnet *ClearnetFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ClearnetUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Clearnet.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClearnetUnpaused)
				if err := _Clearnet.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address indexed account)
func (_Clearnet *ClearnetFilterer) ParseUnpaused(log types.Log) (*ClearnetUnpaused, error) {
	event := new(ClearnetUnpaused)
	if err := _Clearnet.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
