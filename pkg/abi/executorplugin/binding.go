// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executorplugin

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

// Reference imports to suppress goerrors if they are not otherwise used.
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

// ExecutorPluginExecutionRequest is an auto generated low-level Go binding around an user-defined struct.
type ExecutorPluginExecutionRequest struct {
	Exec               TypesExecutable
	Account            common.Address
	Executor           common.Address
	ExecutorSignature  []byte
	ValidatorSignature []byte
}

// TypesExecutable is an auto generated low-level Go binding around an user-defined struct.
type TypesExecutable struct {
	CallType uint8
	Target   common.Address
	Value    *big.Int
	Data     []byte
}

// ExecutorPluginMetaData contains all meta data concerning the ExecutorPlugin contract.
var ExecutorPluginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModuleExecutionFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnableToParseOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProviderTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumTypes.CallType\",\"name\":\"callType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Executable\",\"name\":\"exec\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"executorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"}],\"internalType\":\"structExecutorPlugin.ExecutionRequest\",\"name\":\"execRequest\",\"type\":\"tuple\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"executorNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExecutorPluginABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorPluginMetaData.ABI instead.
var ExecutorPluginABI = ExecutorPluginMetaData.ABI

// ExecutorPlugin is an auto generated Go binding around an Ethereum contract.
type ExecutorPlugin struct {
	ExecutorPluginCaller     // Read-only binding to the contract
	ExecutorPluginTransactor // Write-only binding to the contract
	ExecutorPluginFilterer   // Log filterer for contract events
}

// ExecutorPluginCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorPluginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorPluginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorPluginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorPluginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorPluginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorPluginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorPluginSession struct {
	Contract     *ExecutorPlugin   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorPluginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorPluginCallerSession struct {
	Contract *ExecutorPluginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExecutorPluginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorPluginTransactorSession struct {
	Contract     *ExecutorPluginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExecutorPluginRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorPluginRaw struct {
	Contract *ExecutorPlugin // Generic contract binding to access the raw methods on
}

// ExecutorPluginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorPluginCallerRaw struct {
	Contract *ExecutorPluginCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorPluginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorPluginTransactorRaw struct {
	Contract *ExecutorPluginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutorPlugin creates a new instance of ExecutorPlugin, bound to a specific deployed contract.
func NewExecutorPlugin(address common.Address, backend bind.ContractBackend) (*ExecutorPlugin, error) {
	contract, err := bindExecutorPlugin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutorPlugin{ExecutorPluginCaller: ExecutorPluginCaller{contract: contract}, ExecutorPluginTransactor: ExecutorPluginTransactor{contract: contract}, ExecutorPluginFilterer: ExecutorPluginFilterer{contract: contract}}, nil
}

// NewExecutorPluginCaller creates a new read-only instance of ExecutorPlugin, bound to a specific deployed contract.
func NewExecutorPluginCaller(address common.Address, caller bind.ContractCaller) (*ExecutorPluginCaller, error) {
	contract, err := bindExecutorPlugin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorPluginCaller{contract: contract}, nil
}

// NewExecutorPluginTransactor creates a new write-only instance of ExecutorPlugin, bound to a specific deployed contract.
func NewExecutorPluginTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorPluginTransactor, error) {
	contract, err := bindExecutorPlugin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorPluginTransactor{contract: contract}, nil
}

// NewExecutorPluginFilterer creates a new log filterer instance of ExecutorPlugin, bound to a specific deployed contract.
func NewExecutorPluginFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorPluginFilterer, error) {
	contract, err := bindExecutorPlugin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorPluginFilterer{contract: contract}, nil
}

// bindExecutorPlugin binds a generic wrapper to an already deployed contract.
func bindExecutorPlugin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutorPluginMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutorPlugin *ExecutorPluginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorPlugin.Contract.ExecutorPluginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutorPlugin *ExecutorPluginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorPlugin.Contract.ExecutorPluginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorPlugin *ExecutorPluginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorPlugin.Contract.ExecutorPluginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutorPlugin *ExecutorPluginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorPlugin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutorPlugin *ExecutorPluginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorPlugin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorPlugin *ExecutorPluginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorPlugin.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_ExecutorPlugin *ExecutorPluginSession) AddressProvider() (common.Address, error) {
	return _ExecutorPlugin.Contract.AddressProvider(&_ExecutorPlugin.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCallerSession) AddressProvider() (common.Address, error) {
	return _ExecutorPlugin.Contract.AddressProvider(&_ExecutorPlugin.CallOpts)
}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCaller) AddressProviderTarget(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "addressProviderTarget")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_ExecutorPlugin *ExecutorPluginSession) AddressProviderTarget() (common.Address, error) {
	return _ExecutorPlugin.Contract.AddressProviderTarget(&_ExecutorPlugin.CallOpts)
}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCallerSession) AddressProviderTarget() (common.Address, error) {
	return _ExecutorPlugin.Contract.AddressProviderTarget(&_ExecutorPlugin.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ExecutorPlugin *ExecutorPluginCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ExecutorPlugin *ExecutorPluginSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ExecutorPlugin.Contract.Eip712Domain(&_ExecutorPlugin.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ExecutorPlugin *ExecutorPluginCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ExecutorPlugin.Contract.Eip712Domain(&_ExecutorPlugin.CallOpts)
}

// ExecutorNonce is a free data retrieval call binding the contract method 0x4611b4c7.
//
// Solidity: function executorNonce(address account, address executor) view returns(uint256 nonce)
func (_ExecutorPlugin *ExecutorPluginCaller) ExecutorNonce(opts *bind.CallOpts, account common.Address, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "executorNonce", account, executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutorNonce is a free data retrieval call binding the contract method 0x4611b4c7.
//
// Solidity: function executorNonce(address account, address executor) view returns(uint256 nonce)
func (_ExecutorPlugin *ExecutorPluginSession) ExecutorNonce(account common.Address, executor common.Address) (*big.Int, error) {
	return _ExecutorPlugin.Contract.ExecutorNonce(&_ExecutorPlugin.CallOpts, account, executor)
}

// ExecutorNonce is a free data retrieval call binding the contract method 0x4611b4c7.
//
// Solidity: function executorNonce(address account, address executor) view returns(uint256 nonce)
func (_ExecutorPlugin *ExecutorPluginCallerSession) ExecutorNonce(account common.Address, executor common.Address) (*big.Int, error) {
	return _ExecutorPlugin.Contract.ExecutorNonce(&_ExecutorPlugin.CallOpts, account, executor)
}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCaller) ExecutorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "executorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginSession) ExecutorRegistry() (common.Address, error) {
	return _ExecutorPlugin.Contract.ExecutorRegistry(&_ExecutorPlugin.CallOpts)
}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCallerSession) ExecutorRegistry() (common.Address, error) {
	return _ExecutorPlugin.Contract.ExecutorRegistry(&_ExecutorPlugin.CallOpts)
}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCaller) PolicyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "policyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginSession) PolicyRegistry() (common.Address, error) {
	return _ExecutorPlugin.Contract.PolicyRegistry(&_ExecutorPlugin.CallOpts)
}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCallerSession) PolicyRegistry() (common.Address, error) {
	return _ExecutorPlugin.Contract.PolicyRegistry(&_ExecutorPlugin.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCaller) WalletRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorPlugin.contract.Call(opts, &out, "walletRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginSession) WalletRegistry() (common.Address, error) {
	return _ExecutorPlugin.Contract.WalletRegistry(&_ExecutorPlugin.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_ExecutorPlugin *ExecutorPluginCallerSession) WalletRegistry() (common.Address, error) {
	return _ExecutorPlugin.Contract.WalletRegistry(&_ExecutorPlugin.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2bf4762b.
//
// Solidity: function executeTransaction(((uint8,address,uint256,bytes),address,address,bytes,bytes) execRequest) returns(bytes)
func (_ExecutorPlugin *ExecutorPluginTransactor) ExecuteTransaction(opts *bind.TransactOpts, execRequest ExecutorPluginExecutionRequest) (*types.Transaction, error) {
	return _ExecutorPlugin.contract.Transact(opts, "executeTransaction", execRequest)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2bf4762b.
//
// Solidity: function executeTransaction(((uint8,address,uint256,bytes),address,address,bytes,bytes) execRequest) returns(bytes)
func (_ExecutorPlugin *ExecutorPluginSession) ExecuteTransaction(execRequest ExecutorPluginExecutionRequest) (*types.Transaction, error) {
	return _ExecutorPlugin.Contract.ExecuteTransaction(&_ExecutorPlugin.TransactOpts, execRequest)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2bf4762b.
//
// Solidity: function executeTransaction(((uint8,address,uint256,bytes),address,address,bytes,bytes) execRequest) returns(bytes)
func (_ExecutorPlugin *ExecutorPluginTransactorSession) ExecuteTransaction(execRequest ExecutorPluginExecutionRequest) (*types.Transaction, error) {
	return _ExecutorPlugin.Contract.ExecuteTransaction(&_ExecutorPlugin.TransactOpts, execRequest)
}
