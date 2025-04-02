// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package manageroot

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

// ManageRootMetaData contains all meta data concerning the ManageRoot contract.
var ManageRootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_balancerVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__BadFlashLoanIntentHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ManagerWithMerkleVerification__FailedToVerifyManageProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__FlashLoanNotExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__FlashLoanNotInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__InvalidDecodersAndSanitizersLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__InvalidManageProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__InvalidTargetDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__InvalidValuesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__OnlyCallableByBalancerVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__OnlyCallableByBoringVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerWithMerkleVerification__TotalSupplyMustRemainConstantDuringManagement\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractAuthority\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callsMade\",\"type\":\"uint256\"}],\"name\":\"BoringVaultManaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"ManageRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balancerVault\",\"outputs\":[{\"internalType\":\"contractBalancerVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"manageRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"manageProofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"address[]\",\"name\":\"decodersAndSanitizers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"targetData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"manageVaultWithMerkleVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"receiveFlashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategist\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_manageRoot\",\"type\":\"bytes32\"}],\"name\":\"setManageRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractBoringVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ManageRootABI is the input ABI used to generate the binding from.
// Deprecated: Use ManageRootMetaData.ABI instead.
var ManageRootABI = ManageRootMetaData.ABI

// ManageRoot is an auto generated Go binding around an Ethereum contract.
type ManageRoot struct {
	ManageRootCaller     // Read-only binding to the contract
	ManageRootTransactor // Write-only binding to the contract
	ManageRootFilterer   // Log filterer for contract events
}

// ManageRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManageRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManageRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManageRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManageRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManageRootSession struct {
	Contract     *ManageRoot       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManageRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManageRootCallerSession struct {
	Contract *ManageRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ManageRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManageRootTransactorSession struct {
	Contract     *ManageRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ManageRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManageRootRaw struct {
	Contract *ManageRoot // Generic contract binding to access the raw methods on
}

// ManageRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManageRootCallerRaw struct {
	Contract *ManageRootCaller // Generic read-only contract binding to access the raw methods on
}

// ManageRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManageRootTransactorRaw struct {
	Contract *ManageRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManageRoot creates a new instance of ManageRoot, bound to a specific deployed contract.
func NewManageRoot(address common.Address, backend bind.ContractBackend) (*ManageRoot, error) {
	contract, err := bindManageRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ManageRoot{ManageRootCaller: ManageRootCaller{contract: contract}, ManageRootTransactor: ManageRootTransactor{contract: contract}, ManageRootFilterer: ManageRootFilterer{contract: contract}}, nil
}

// NewManageRootCaller creates a new read-only instance of ManageRoot, bound to a specific deployed contract.
func NewManageRootCaller(address common.Address, caller bind.ContractCaller) (*ManageRootCaller, error) {
	contract, err := bindManageRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManageRootCaller{contract: contract}, nil
}

// NewManageRootTransactor creates a new write-only instance of ManageRoot, bound to a specific deployed contract.
func NewManageRootTransactor(address common.Address, transactor bind.ContractTransactor) (*ManageRootTransactor, error) {
	contract, err := bindManageRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManageRootTransactor{contract: contract}, nil
}

// NewManageRootFilterer creates a new log filterer instance of ManageRoot, bound to a specific deployed contract.
func NewManageRootFilterer(address common.Address, filterer bind.ContractFilterer) (*ManageRootFilterer, error) {
	contract, err := bindManageRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManageRootFilterer{contract: contract}, nil
}

// bindManageRoot binds a generic wrapper to an already deployed contract.
func bindManageRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ManageRootMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManageRoot *ManageRootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManageRoot.Contract.ManageRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManageRoot *ManageRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManageRoot.Contract.ManageRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManageRoot *ManageRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManageRoot.Contract.ManageRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManageRoot *ManageRootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManageRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManageRoot *ManageRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManageRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManageRoot *ManageRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManageRoot.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_ManageRoot *ManageRootCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManageRoot.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_ManageRoot *ManageRootSession) Authority() (common.Address, error) {
	return _ManageRoot.Contract.Authority(&_ManageRoot.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_ManageRoot *ManageRootCallerSession) Authority() (common.Address, error) {
	return _ManageRoot.Contract.Authority(&_ManageRoot.CallOpts)
}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_ManageRoot *ManageRootCaller) BalancerVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManageRoot.contract.Call(opts, &out, "balancerVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_ManageRoot *ManageRootSession) BalancerVault() (common.Address, error) {
	return _ManageRoot.Contract.BalancerVault(&_ManageRoot.CallOpts)
}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_ManageRoot *ManageRootCallerSession) BalancerVault() (common.Address, error) {
	return _ManageRoot.Contract.BalancerVault(&_ManageRoot.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_ManageRoot *ManageRootCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ManageRoot.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_ManageRoot *ManageRootSession) IsPaused() (bool, error) {
	return _ManageRoot.Contract.IsPaused(&_ManageRoot.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_ManageRoot *ManageRootCallerSession) IsPaused() (bool, error) {
	return _ManageRoot.Contract.IsPaused(&_ManageRoot.CallOpts)
}

// ManageRoot is a free data retrieval call binding the contract method 0x5ca58a99.
//
// Solidity: function manageRoot(address ) view returns(bytes32)
func (_ManageRoot *ManageRootCaller) ManageRoot(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ManageRoot.contract.Call(opts, &out, "manageRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ManageRoot is a free data retrieval call binding the contract method 0x5ca58a99.
//
// Solidity: function manageRoot(address ) view returns(bytes32)
func (_ManageRoot *ManageRootSession) ManageRoot(arg0 common.Address) ([32]byte, error) {
	return _ManageRoot.Contract.ManageRoot(&_ManageRoot.CallOpts, arg0)
}

// ManageRoot is a free data retrieval call binding the contract method 0x5ca58a99.
//
// Solidity: function manageRoot(address ) view returns(bytes32)
func (_ManageRoot *ManageRootCallerSession) ManageRoot(arg0 common.Address) ([32]byte, error) {
	return _ManageRoot.Contract.ManageRoot(&_ManageRoot.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ManageRoot *ManageRootCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManageRoot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ManageRoot *ManageRootSession) Owner() (common.Address, error) {
	return _ManageRoot.Contract.Owner(&_ManageRoot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ManageRoot *ManageRootCallerSession) Owner() (common.Address, error) {
	return _ManageRoot.Contract.Owner(&_ManageRoot.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ManageRoot *ManageRootCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManageRoot.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ManageRoot *ManageRootSession) Vault() (common.Address, error) {
	return _ManageRoot.Contract.Vault(&_ManageRoot.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ManageRoot *ManageRootCallerSession) Vault() (common.Address, error) {
	return _ManageRoot.Contract.Vault(&_ManageRoot.CallOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_ManageRoot *ManageRootTransactor) FlashLoan(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "flashLoan", recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_ManageRoot *ManageRootSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _ManageRoot.Contract.FlashLoan(&_ManageRoot.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_ManageRoot *ManageRootTransactorSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _ManageRoot.Contract.FlashLoan(&_ManageRoot.TransactOpts, recipient, tokens, amounts, userData)
}

// ManageVaultWithMerkleVerification is a paid mutator transaction binding the contract method 0x244b0f6a.
//
// Solidity: function manageVaultWithMerkleVerification(bytes32[][] manageProofs, address[] decodersAndSanitizers, address[] targets, bytes[] targetData, uint256[] values) returns()
func (_ManageRoot *ManageRootTransactor) ManageVaultWithMerkleVerification(opts *bind.TransactOpts, manageProofs [][][32]byte, decodersAndSanitizers []common.Address, targets []common.Address, targetData [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "manageVaultWithMerkleVerification", manageProofs, decodersAndSanitizers, targets, targetData, values)
}

// ManageVaultWithMerkleVerification is a paid mutator transaction binding the contract method 0x244b0f6a.
//
// Solidity: function manageVaultWithMerkleVerification(bytes32[][] manageProofs, address[] decodersAndSanitizers, address[] targets, bytes[] targetData, uint256[] values) returns()
func (_ManageRoot *ManageRootSession) ManageVaultWithMerkleVerification(manageProofs [][][32]byte, decodersAndSanitizers []common.Address, targets []common.Address, targetData [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _ManageRoot.Contract.ManageVaultWithMerkleVerification(&_ManageRoot.TransactOpts, manageProofs, decodersAndSanitizers, targets, targetData, values)
}

// ManageVaultWithMerkleVerification is a paid mutator transaction binding the contract method 0x244b0f6a.
//
// Solidity: function manageVaultWithMerkleVerification(bytes32[][] manageProofs, address[] decodersAndSanitizers, address[] targets, bytes[] targetData, uint256[] values) returns()
func (_ManageRoot *ManageRootTransactorSession) ManageVaultWithMerkleVerification(manageProofs [][][32]byte, decodersAndSanitizers []common.Address, targets []common.Address, targetData [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _ManageRoot.Contract.ManageVaultWithMerkleVerification(&_ManageRoot.TransactOpts, manageProofs, decodersAndSanitizers, targets, targetData, values)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ManageRoot *ManageRootTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ManageRoot *ManageRootSession) Pause() (*types.Transaction, error) {
	return _ManageRoot.Contract.Pause(&_ManageRoot.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ManageRoot *ManageRootTransactorSession) Pause() (*types.Transaction, error) {
	return _ManageRoot.Contract.Pause(&_ManageRoot.TransactOpts)
}

// ReceiveFlashLoan is a paid mutator transaction binding the contract method 0xf04f2707.
//
// Solidity: function receiveFlashLoan(address[] tokens, uint256[] amounts, uint256[] feeAmounts, bytes userData) returns()
func (_ManageRoot *ManageRootTransactor) ReceiveFlashLoan(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, feeAmounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "receiveFlashLoan", tokens, amounts, feeAmounts, userData)
}

// ReceiveFlashLoan is a paid mutator transaction binding the contract method 0xf04f2707.
//
// Solidity: function receiveFlashLoan(address[] tokens, uint256[] amounts, uint256[] feeAmounts, bytes userData) returns()
func (_ManageRoot *ManageRootSession) ReceiveFlashLoan(tokens []common.Address, amounts []*big.Int, feeAmounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _ManageRoot.Contract.ReceiveFlashLoan(&_ManageRoot.TransactOpts, tokens, amounts, feeAmounts, userData)
}

// ReceiveFlashLoan is a paid mutator transaction binding the contract method 0xf04f2707.
//
// Solidity: function receiveFlashLoan(address[] tokens, uint256[] amounts, uint256[] feeAmounts, bytes userData) returns()
func (_ManageRoot *ManageRootTransactorSession) ReceiveFlashLoan(tokens []common.Address, amounts []*big.Int, feeAmounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _ManageRoot.Contract.ReceiveFlashLoan(&_ManageRoot.TransactOpts, tokens, amounts, feeAmounts, userData)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_ManageRoot *ManageRootTransactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_ManageRoot *ManageRootSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _ManageRoot.Contract.SetAuthority(&_ManageRoot.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_ManageRoot *ManageRootTransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _ManageRoot.Contract.SetAuthority(&_ManageRoot.TransactOpts, newAuthority)
}

// SetManageRoot is a paid mutator transaction binding the contract method 0x21801a99.
//
// Solidity: function setManageRoot(address strategist, bytes32 _manageRoot) returns()
func (_ManageRoot *ManageRootTransactor) SetManageRoot(opts *bind.TransactOpts, strategist common.Address, _manageRoot [32]byte) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "setManageRoot", strategist, _manageRoot)
}

// SetManageRoot is a paid mutator transaction binding the contract method 0x21801a99.
//
// Solidity: function setManageRoot(address strategist, bytes32 _manageRoot) returns()
func (_ManageRoot *ManageRootSession) SetManageRoot(strategist common.Address, _manageRoot [32]byte) (*types.Transaction, error) {
	return _ManageRoot.Contract.SetManageRoot(&_ManageRoot.TransactOpts, strategist, _manageRoot)
}

// SetManageRoot is a paid mutator transaction binding the contract method 0x21801a99.
//
// Solidity: function setManageRoot(address strategist, bytes32 _manageRoot) returns()
func (_ManageRoot *ManageRootTransactorSession) SetManageRoot(strategist common.Address, _manageRoot [32]byte) (*types.Transaction, error) {
	return _ManageRoot.Contract.SetManageRoot(&_ManageRoot.TransactOpts, strategist, _manageRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ManageRoot *ManageRootTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ManageRoot *ManageRootSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ManageRoot.Contract.TransferOwnership(&_ManageRoot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ManageRoot *ManageRootTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ManageRoot.Contract.TransferOwnership(&_ManageRoot.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ManageRoot *ManageRootTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManageRoot.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ManageRoot *ManageRootSession) Unpause() (*types.Transaction, error) {
	return _ManageRoot.Contract.Unpause(&_ManageRoot.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ManageRoot *ManageRootTransactorSession) Unpause() (*types.Transaction, error) {
	return _ManageRoot.Contract.Unpause(&_ManageRoot.TransactOpts)
}

// ManageRootAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the ManageRoot contract.
type ManageRootAuthorityUpdatedIterator struct {
	Event *ManageRootAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *ManageRootAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManageRootAuthorityUpdated)
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
		it.Event = new(ManageRootAuthorityUpdated)
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
func (it *ManageRootAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManageRootAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManageRootAuthorityUpdated represents a AuthorityUpdated event raised by the ManageRoot contract.
type ManageRootAuthorityUpdated struct {
	User         common.Address
	NewAuthority common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_ManageRoot *ManageRootFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, user []common.Address, newAuthority []common.Address) (*ManageRootAuthorityUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _ManageRoot.contract.FilterLogs(opts, "AuthorityUpdated", userRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &ManageRootAuthorityUpdatedIterator{contract: _ManageRoot.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_ManageRoot *ManageRootFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *ManageRootAuthorityUpdated, user []common.Address, newAuthority []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _ManageRoot.contract.WatchLogs(opts, "AuthorityUpdated", userRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManageRootAuthorityUpdated)
				if err := _ManageRoot.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_ManageRoot *ManageRootFilterer) ParseAuthorityUpdated(log types.Log) (*ManageRootAuthorityUpdated, error) {
	event := new(ManageRootAuthorityUpdated)
	if err := _ManageRoot.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManageRootBoringVaultManagedIterator is returned from FilterBoringVaultManaged and is used to iterate over the raw logs and unpacked data for BoringVaultManaged events raised by the ManageRoot contract.
type ManageRootBoringVaultManagedIterator struct {
	Event *ManageRootBoringVaultManaged // Event containing the contract specifics and raw log

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
func (it *ManageRootBoringVaultManagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManageRootBoringVaultManaged)
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
		it.Event = new(ManageRootBoringVaultManaged)
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
func (it *ManageRootBoringVaultManagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManageRootBoringVaultManagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManageRootBoringVaultManaged represents a BoringVaultManaged event raised by the ManageRoot contract.
type ManageRootBoringVaultManaged struct {
	CallsMade *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBoringVaultManaged is a free log retrieval operation binding the contract event 0x53d426e7d80bb2c8674d3b45577e2d464d423faad6531b21f95ac11ac18b1cb6.
//
// Solidity: event BoringVaultManaged(uint256 callsMade)
func (_ManageRoot *ManageRootFilterer) FilterBoringVaultManaged(opts *bind.FilterOpts) (*ManageRootBoringVaultManagedIterator, error) {

	logs, sub, err := _ManageRoot.contract.FilterLogs(opts, "BoringVaultManaged")
	if err != nil {
		return nil, err
	}
	return &ManageRootBoringVaultManagedIterator{contract: _ManageRoot.contract, event: "BoringVaultManaged", logs: logs, sub: sub}, nil
}

// WatchBoringVaultManaged is a free log subscription operation binding the contract event 0x53d426e7d80bb2c8674d3b45577e2d464d423faad6531b21f95ac11ac18b1cb6.
//
// Solidity: event BoringVaultManaged(uint256 callsMade)
func (_ManageRoot *ManageRootFilterer) WatchBoringVaultManaged(opts *bind.WatchOpts, sink chan<- *ManageRootBoringVaultManaged) (event.Subscription, error) {

	logs, sub, err := _ManageRoot.contract.WatchLogs(opts, "BoringVaultManaged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManageRootBoringVaultManaged)
				if err := _ManageRoot.contract.UnpackLog(event, "BoringVaultManaged", log); err != nil {
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

// ParseBoringVaultManaged is a log parse operation binding the contract event 0x53d426e7d80bb2c8674d3b45577e2d464d423faad6531b21f95ac11ac18b1cb6.
//
// Solidity: event BoringVaultManaged(uint256 callsMade)
func (_ManageRoot *ManageRootFilterer) ParseBoringVaultManaged(log types.Log) (*ManageRootBoringVaultManaged, error) {
	event := new(ManageRootBoringVaultManaged)
	if err := _ManageRoot.contract.UnpackLog(event, "BoringVaultManaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManageRootManageRootUpdatedIterator is returned from FilterManageRootUpdated and is used to iterate over the raw logs and unpacked data for ManageRootUpdated events raised by the ManageRoot contract.
type ManageRootManageRootUpdatedIterator struct {
	Event *ManageRootManageRootUpdated // Event containing the contract specifics and raw log

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
func (it *ManageRootManageRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManageRootManageRootUpdated)
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
		it.Event = new(ManageRootManageRootUpdated)
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
func (it *ManageRootManageRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManageRootManageRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManageRootManageRootUpdated represents a ManageRootUpdated event raised by the ManageRoot contract.
type ManageRootManageRootUpdated struct {
	Strategist common.Address
	OldRoot    [32]byte
	NewRoot    [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManageRootUpdated is a free log retrieval operation binding the contract event 0x0b958dec85f1470000479dfb22c365829411f52bcde602d24ea0abf5ac7e8860.
//
// Solidity: event ManageRootUpdated(address indexed strategist, bytes32 oldRoot, bytes32 newRoot)
func (_ManageRoot *ManageRootFilterer) FilterManageRootUpdated(opts *bind.FilterOpts, strategist []common.Address) (*ManageRootManageRootUpdatedIterator, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _ManageRoot.contract.FilterLogs(opts, "ManageRootUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return &ManageRootManageRootUpdatedIterator{contract: _ManageRoot.contract, event: "ManageRootUpdated", logs: logs, sub: sub}, nil
}

// WatchManageRootUpdated is a free log subscription operation binding the contract event 0x0b958dec85f1470000479dfb22c365829411f52bcde602d24ea0abf5ac7e8860.
//
// Solidity: event ManageRootUpdated(address indexed strategist, bytes32 oldRoot, bytes32 newRoot)
func (_ManageRoot *ManageRootFilterer) WatchManageRootUpdated(opts *bind.WatchOpts, sink chan<- *ManageRootManageRootUpdated, strategist []common.Address) (event.Subscription, error) {

	var strategistRule []interface{}
	for _, strategistItem := range strategist {
		strategistRule = append(strategistRule, strategistItem)
	}

	logs, sub, err := _ManageRoot.contract.WatchLogs(opts, "ManageRootUpdated", strategistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManageRootManageRootUpdated)
				if err := _ManageRoot.contract.UnpackLog(event, "ManageRootUpdated", log); err != nil {
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

// ParseManageRootUpdated is a log parse operation binding the contract event 0x0b958dec85f1470000479dfb22c365829411f52bcde602d24ea0abf5ac7e8860.
//
// Solidity: event ManageRootUpdated(address indexed strategist, bytes32 oldRoot, bytes32 newRoot)
func (_ManageRoot *ManageRootFilterer) ParseManageRootUpdated(log types.Log) (*ManageRootManageRootUpdated, error) {
	event := new(ManageRootManageRootUpdated)
	if err := _ManageRoot.contract.UnpackLog(event, "ManageRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManageRootOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ManageRoot contract.
type ManageRootOwnershipTransferredIterator struct {
	Event *ManageRootOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ManageRootOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManageRootOwnershipTransferred)
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
		it.Event = new(ManageRootOwnershipTransferred)
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
func (it *ManageRootOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManageRootOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManageRootOwnershipTransferred represents a OwnershipTransferred event raised by the ManageRoot contract.
type ManageRootOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_ManageRoot *ManageRootFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*ManageRootOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ManageRoot.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManageRootOwnershipTransferredIterator{contract: _ManageRoot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_ManageRoot *ManageRootFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ManageRootOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ManageRoot.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManageRootOwnershipTransferred)
				if err := _ManageRoot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_ManageRoot *ManageRootFilterer) ParseOwnershipTransferred(log types.Log) (*ManageRootOwnershipTransferred, error) {
	event := new(ManageRootOwnershipTransferred)
	if err := _ManageRoot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManageRootPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ManageRoot contract.
type ManageRootPausedIterator struct {
	Event *ManageRootPaused // Event containing the contract specifics and raw log

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
func (it *ManageRootPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManageRootPaused)
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
		it.Event = new(ManageRootPaused)
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
func (it *ManageRootPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManageRootPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManageRootPaused represents a Paused event raised by the ManageRoot contract.
type ManageRootPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_ManageRoot *ManageRootFilterer) FilterPaused(opts *bind.FilterOpts) (*ManageRootPausedIterator, error) {

	logs, sub, err := _ManageRoot.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ManageRootPausedIterator{contract: _ManageRoot.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_ManageRoot *ManageRootFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ManageRootPaused) (event.Subscription, error) {

	logs, sub, err := _ManageRoot.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManageRootPaused)
				if err := _ManageRoot.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_ManageRoot *ManageRootFilterer) ParsePaused(log types.Log) (*ManageRootPaused, error) {
	event := new(ManageRootPaused)
	if err := _ManageRoot.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManageRootUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ManageRoot contract.
type ManageRootUnpausedIterator struct {
	Event *ManageRootUnpaused // Event containing the contract specifics and raw log

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
func (it *ManageRootUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManageRootUnpaused)
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
		it.Event = new(ManageRootUnpaused)
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
func (it *ManageRootUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManageRootUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManageRootUnpaused represents a Unpaused event raised by the ManageRoot contract.
type ManageRootUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_ManageRoot *ManageRootFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ManageRootUnpausedIterator, error) {

	logs, sub, err := _ManageRoot.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ManageRootUnpausedIterator{contract: _ManageRoot.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_ManageRoot *ManageRootFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ManageRootUnpaused) (event.Subscription, error) {

	logs, sub, err := _ManageRoot.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManageRootUnpaused)
				if err := _ManageRoot.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_ManageRoot *ManageRootFilterer) ParseUnpaused(log types.Log) (*ManageRootUnpaused, error) {
	event := new(ManageRootUnpaused)
	if err := _ManageRoot.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
